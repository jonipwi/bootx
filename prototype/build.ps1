[CmdletBinding()]
param(
    [ValidateSet('build', 'test', 'verify', 'run', 'clean')]
    [string]$Action = 'build',

    [string]$OutputDirectory = '',

    [switch]$SkipTests
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$moduleRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'personal-companion'))
$repositoryRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot '..'))
if (-not (Test-Path -LiteralPath (Join-Path $moduleRoot 'go.mod'))) {
    throw "BootX companion module not found at: $moduleRoot"
}

if ([string]::IsNullOrWhiteSpace($OutputDirectory)) {
    $OutputDirectory = Join-Path $moduleRoot 'dist'
}
$outputRoot = [IO.Path]::GetFullPath($OutputDirectory)
$defaultOutputRoot = [IO.Path]::GetFullPath((Join-Path $moduleRoot 'dist'))

function Invoke-Go {
    param([Parameter(Mandatory = $true)][string[]]$GoArguments)

    Write-Host ('> go ' + ($GoArguments -join ' ')) -ForegroundColor DarkGray
    $commandOutput = @(& go @GoArguments 2>&1)
    $exitCode = $LASTEXITCODE
    foreach ($line in $commandOutput) {
        Write-Host $line
    }
    if ($exitCode -ne 0) {
        throw "go command failed with exit code ${exitCode}: go $($GoArguments -join ' ')"
    }
}

function Test-GoFormatting {
    Write-Host 'Checking Go formatting...' -ForegroundColor Cyan
    $unformatted = @(& gofmt -l .\cmd .\internal)
    if ($LASTEXITCODE -ne 0) {
        throw "gofmt failed with exit code $LASTEXITCODE"
    }
    if ($unformatted.Count -gt 0) {
        throw "Go files require formatting:`n$($unformatted -join [Environment]::NewLine)`nRun: gofmt -w .\cmd .\internal"
    }
}

function Test-JSONAssets {
    Write-Host 'Validating JSON schemas, policy, and fixtures...' -ForegroundColor Cyan
    $jsonFiles = Get-ChildItem -Path @(
        (Join-Path $moduleRoot 'schemas\*.json'),
        (Join-Path $moduleRoot 'internal\policy\config\*.json'),
        (Join-Path $moduleRoot 'testdata\*.json')
    ) -File

    foreach ($file in $jsonFiles) {
        try {
            $null = Get-Content -LiteralPath $file.FullName -Raw -Encoding UTF8 | ConvertFrom-Json
        }
        catch {
            throw "Invalid JSON asset $($file.FullName): $($_.Exception.Message)"
        }
    }
    Write-Host "Validated $($jsonFiles.Count) JSON assets." -ForegroundColor Green
}

function Invoke-GoTests {
    Test-GoFormatting
    Invoke-Go -GoArguments @('test', './...')
    Invoke-Go -GoArguments @('vet', './...')
}

function New-BootXBuild {
    if (-not $SkipTests) {
        Invoke-GoTests
    }
    else {
        Write-Warning 'Tests and vet were skipped by explicit request.'
    }

    Test-JSONAssets
    $null = New-Item -ItemType Directory -Force -Path $outputRoot

    $goos = (& go env GOOS).Trim()
    if ($LASTEXITCODE -ne 0) { throw 'Unable to determine GOOS.' }
    $goarch = (& go env GOARCH).Trim()
    if ($LASTEXITCODE -ne 0) { throw 'Unable to determine GOARCH.' }

    $extension = if ($goos -eq 'windows') { '.exe' } else { '' }
    $binaryName = "bootx-companion-$goos-$goarch$extension"
    $binaryPath = Join-Path $outputRoot $binaryName

    Invoke-Go -GoArguments @(
        'build',
        '-trimpath',
        '-buildvcs=false',
        '-o', $binaryPath,
        './cmd/bootx-companion'
    )

    $appVersion = (& $binaryPath -version).Trim()
    if ($LASTEXITCODE -ne 0) {
        throw 'Built executable failed its version smoke test.'
    }
    $hash = (Get-FileHash -LiteralPath $binaryPath -Algorithm SHA256).Hash.ToLowerInvariant()
    $goVersion = (& go version).Trim()
    $moduleList = @(& go list -m all)
    if ($LASTEXITCODE -ne 0) {
        throw 'Unable to enumerate Go module dependencies.'
    }
    $externalModuleCount = [Math]::Max(0, $moduleList.Count - 1)

    $manifest = [ordered]@{
        project = 'BootX Personal Companion MVP'
        capability = 'assist.personal-decision.v1'
        capabilities = @(
            'assist.personal-decision.v1'
            'assist.law-clarity.v1'
            'assist.ethical-review.v1'
        )
        application_version = $appVersion
        build_utc = [DateTime]::UtcNow.ToString('o')
        go_version = $goVersion
        goos = $goos
        goarch = $goarch
        binary = $binaryName
        sha256 = $hash
        tests_skipped = [bool]$SkipTests
        external_modules = $externalModuleCount
        safety_status = 'DEV-1 hybrid prototype; deterministic core plus explicit-consent OpenAI advisory; no deployment, certification, legal authority, automatic publication, or enforcement capability'
    }

    $manifestPath = Join-Path $outputRoot 'build-manifest.json'
    $manifest | ConvertTo-Json -Depth 4 | Set-Content -LiteralPath $manifestPath -Encoding UTF8
    "$hash  $binaryName" | Set-Content -LiteralPath (Join-Path $outputRoot "$binaryName.sha256") -Encoding ASCII

    Write-Host ''
    Write-Host 'Build completed.' -ForegroundColor Green
    Write-Host "  Binary:   $binaryPath"
    Write-Host "  Version:  $appVersion"
    Write-Host "  SHA-256:  $hash"
    Write-Host "  Manifest: $manifestPath"

    return $binaryPath
}

function Test-Fixtures {
    param([Parameter(Mandatory = $true)][string]$BinaryPath)

    Write-Host 'Running backend fixture smoke tests...' -ForegroundColor Cyan
    $expected = [ordered]@{
        'safe-study.json' = @{ Class = 'D0'; Mode = 'INFORM'; Warning = $null }
        'suspicious-message.json' = @{ Class = 'D2'; Mode = 'VERIFY'; Warning = $null }
        'warning-prepare.json' = @{ Class = 'D2'; Mode = 'PREPARE'; Warning = 'W2' }
        'warning-urgent.json' = @{ Class = 'D4'; Mode = 'URGENT_GUIDANCE'; Warning = 'W4' }
        'typhoon-bavi-exercise.json' = @{ Class = 'D2'; Mode = 'PREPARE'; Warning = 'W2' }
        'space-governance-readiness-exercise.json' = @{ Class = 'D3'; Mode = 'ABSTAIN'; Warning = $null }
    }

    foreach ($name in $expected.Keys) {
        $fixturePath = Join-Path (Join-Path $moduleRoot 'testdata') $name
        $packet = (& $BinaryPath -input $fixturePath) | ConvertFrom-Json
        if ($LASTEXITCODE -ne 0) {
            throw "Fixture execution failed: $name"
        }
        $want = $expected[$name]
        if ($packet.decision_class -ne $want.Class -or $packet.response_mode -ne $want.Mode) {
            throw "Fixture $name returned $($packet.decision_class)/$($packet.response_mode); expected $($want.Class)/$($want.Mode)."
        }
        $warningProperty = $packet.PSObject.Properties['warning']
        $warningLevel = if ($null -ne $warningProperty -and $null -ne $warningProperty.Value) { $warningProperty.Value.level } else { '' }
        if ($null -ne $want.Warning -and $warningLevel -ne $want.Warning) {
            throw "Fixture $name returned warning $warningLevel; expected $($want.Warning)."
        }
        Write-Host "  PASS $name -> $($packet.decision_class)/$($packet.response_mode) $warningLevel" -ForegroundColor Green
    }
}

function Test-RealDocumentMode {
    param([Parameter(Mandatory = $true)][string]$BinaryPath)

    Write-Host 'Running contained real-document smoke test...' -ForegroundColor Cyan
    $packet = (& $BinaryPath `
        -workspace $repositoryRoot `
        -document 'docs\research\civilization\religion-ideology-and-decision-integrity.md' `
        -document-public `
        -goal 'Choose the next evidence-improvement task' `
        -question 'Which missing fact should be verified first?' `
        -priorities 'truth,reversibility') | ConvertFrom-Json
    if ($LASTEXITCODE -ne 0) {
        throw 'Contained real-document execution failed.'
    }
    if ($packet.decision_class -ne 'D0' -or $packet.response_mode -ne 'INFORM') {
        throw "Real-document mode returned $($packet.decision_class)/$($packet.response_mode); expected D0/INFORM."
    }
    if (-not [bool]$packet.evidence_receipt.integrity_verified -or $packet.evidence_receipt.reference -ne 'docs/research/civilization/religion-ideology-and-decision-integrity.md') {
        throw 'Real-document evidence receipt is missing or incorrect.'
    }
    if ($packet.evidence_receipt.origin_status -ne 'not_authenticated' -or [bool]$packet.data_receipt.synthetic) {
        throw 'Real-document origin or synthetic status is incorrect.'
    }
    $documentScan = @($packet.observations | Where-Object { $_.source -eq 'deterministic_local_document_scan' })
    if ($documentScan.Count -lt 1) {
        throw 'Real-document structural scan observation is missing.'
    }
    Write-Host "  PASS real research document -> D0/INFORM; SHA-256 $($packet.evidence_receipt.sha256); origin not authenticated" -ForegroundColor Green
}

function Test-LawClarityMode {
    param([Parameter(Mandatory = $true)][string]$BinaryPath)

    Write-Host 'Running Law Clarity Logic smoke test...' -ForegroundColor Cyan
    $fixturePath = Join-Path (Join-Path $moduleRoot 'testdata') 'law-clarity-gray-zone.json'
    $report = (& $BinaryPath -law-input $fixturePath) | ConvertFrom-Json
    if ($LASTEXITCODE -ne 0) {
        throw 'Law Clarity Logic fixture execution failed.'
    }
    if ($report.capability_id -ne 'assist.law-clarity.v1') {
        throw "Law Clarity Logic returned capability $($report.capability_id)."
    }
    if ([double]$report.law_quality_score -ne 35.75 -or [double]$report.gray_zone_risk_score -ne 74.25 -or [double]$report.manipulation_risk_index -ne 48.11) {
        throw "Law Clarity formula mismatch: Q=$($report.law_quality_score), Z=$($report.gray_zone_risk_score), M=$($report.manipulation_risk_index)."
    }
    if ($report.strict_good_law_gate.status -ne 'FAIL' -or $report.human_rights_gate.status -ne 'FAIL') {
        throw 'Law Clarity non-compensable gates did not fail as expected.'
    }
    if (-not [bool]$report.high_manipulation_trigger -or $report.disposition -ne 'FUNDAMENTAL_REVISION_REQUIRED') {
        throw 'Law Clarity disposition or high-manipulation trigger is incorrect.'
    }
    if ($null -ne $report.user_decision -or [bool]$report.input_receipt.remote_processing -or [bool]$report.input_receipt.persistent_memory) {
        throw 'Law Clarity human-authority or data-boundary invariant failed.'
    }
    Write-Host '  PASS law-clarity-gray-zone.json -> Q=35.75, Z=74.25, M=48.11; rights gate FAIL; no legal verdict' -ForegroundColor Green

    $capitalCorruptionCases = [ordered]@{
        'law-clarity-abusive-capital-corruption-proposal.json' = @{
            Q = 12.0
            Z = 92.75
            M = 92.75
            Strict = 'FAIL'
            Rights = 'FAIL'
            Trigger = $true
            Disposition = 'FUNDAMENTAL_REVISION_REQUIRED'
            PhraseTypes = 6
        }
        'law-clarity-rights-preserving-anti-corruption-revision.json' = @{
            Q = 92.4
            Z = 7.7
            M = 0.01
            Strict = 'PASS'
            Rights = 'PASS'
            Trigger = $false
            Disposition = 'QUALIFIED_REVIEW_REQUIRED'
            PhraseTypes = 0
        }
    }
    foreach ($name in $capitalCorruptionCases.Keys) {
        $caseFixture = Join-Path (Join-Path $moduleRoot 'testdata') $name
        $caseReport = (& $BinaryPath -law-input $caseFixture) | ConvertFrom-Json
        if ($LASTEXITCODE -ne 0) {
            throw "Law Clarity capital-corruption fixture execution failed: $name"
        }
        $want = $capitalCorruptionCases[$name]
        if ([double]$caseReport.law_quality_score -ne $want.Q -or [double]$caseReport.gray_zone_risk_score -ne $want.Z -or [double]$caseReport.manipulation_risk_index -ne $want.M) {
            throw "Law Clarity formula mismatch for ${name}: Q=$($caseReport.law_quality_score), Z=$($caseReport.gray_zone_risk_score), M=$($caseReport.manipulation_risk_index)."
        }
        if ($caseReport.strict_good_law_gate.status -ne $want.Strict -or $caseReport.human_rights_gate.status -ne $want.Rights) {
            throw "Law Clarity gate mismatch for $name."
        }
        if ([bool]$caseReport.high_manipulation_trigger -ne [bool]$want.Trigger -or $caseReport.disposition -ne $want.Disposition) {
            throw "Law Clarity trigger or disposition mismatch for $name."
        }
        if (@($caseReport.visible_phrase_hits).Count -ne $want.PhraseTypes) {
            throw "Law Clarity phrase count mismatch for $name."
        }
        if ($null -ne $caseReport.user_decision -or [bool]$caseReport.input_receipt.remote_processing -or [bool]$caseReport.input_receipt.persistent_memory) {
            throw "Law Clarity authority or data-boundary invariant failed for $name."
        }
        Write-Host "  PASS $name -> Q=$($want.Q), Z=$($want.Z), M=$($want.M); $($want.Disposition); no sentence decision" -ForegroundColor Green
    }
}

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    throw 'Go was not found on PATH. Install Go 1.22 or later.'
}
if (-not (Get-Command gofmt -ErrorAction SilentlyContinue)) {
    throw 'gofmt was not found on PATH.'
}

$cacheRoot = Join-Path ([IO.Path]::GetTempPath()) 'bootx-go-cache'
$tempRoot = Join-Path ([IO.Path]::GetTempPath()) 'bootx-go-tmp'
$null = New-Item -ItemType Directory -Force -Path $cacheRoot
$null = New-Item -ItemType Directory -Force -Path $tempRoot
$env:GOCACHE = $cacheRoot
$env:GOTMPDIR = $tempRoot

Push-Location $moduleRoot
try {
    Write-Host "BootX prototype module: $moduleRoot" -ForegroundColor Cyan
    Write-Host (& go version)

    switch ($Action) {
        'clean' {
            if ($outputRoot -ne $defaultOutputRoot) {
                throw 'Safe clean only permits the default personal-companion\dist directory.'
            }
            if (Test-Path -LiteralPath $outputRoot) {
                Remove-Item -LiteralPath $outputRoot -Recurse -Force
                Write-Host "Removed: $outputRoot" -ForegroundColor Green
            }
            else {
                Write-Host 'Nothing to clean.'
            }
        }
        'test' {
            Invoke-GoTests
            Write-Host 'Tests and vet passed.' -ForegroundColor Green
        }
        'verify' {
            $binary = New-BootXBuild
            Test-Fixtures -BinaryPath $binary
            Test-RealDocumentMode -BinaryPath $binary
            Test-LawClarityMode -BinaryPath $binary
            Write-Host 'Full prototype verification passed.' -ForegroundColor Green
        }
        'run' {
            $binary = New-BootXBuild
            Write-Host 'Starting the BootX terminal UI...' -ForegroundColor Cyan
            & $binary
            if ($LASTEXITCODE -ne 0) {
                throw "BootX terminal UI exited with code $LASTEXITCODE"
            }
        }
        default {
            $null = New-BootXBuild
        }
    }
}
finally {
    Pop-Location
}
