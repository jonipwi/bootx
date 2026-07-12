[CmdletBinding()]
param(
    [string]$LogPath = ''
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$moduleRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'personal-companion'))
$buildScript = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'build.ps1'))
$fixturePath = [IO.Path]::GetFullPath((Join-Path $moduleRoot 'testdata\typhoon-bavi-exercise.json'))
if ([string]::IsNullOrWhiteSpace($LogPath)) {
    $LogPath = Join-Path $PSScriptRoot 'typhoon-bavi.log'
}
$logFullPath = [IO.Path]::GetFullPath($LogPath)
$logDirectory = Split-Path -Parent $logFullPath
$temporaryBuild = Join-Path ([IO.Path]::GetTempPath()) ("bootx-bavi-test-" + [Guid]::NewGuid().ToString('N'))
$utf8NoBom = New-Object Text.UTF8Encoding($false)
$exitCode = 0

function Write-LogLine {
    param([Parameter(Mandatory = $true)][AllowEmptyString()][string]$Text)

    [IO.File]::AppendAllText($logFullPath, $Text + [Environment]::NewLine, $utf8NoBom)
    Write-Host $Text
}

function Write-LogBlock {
    param([Parameter(Mandatory = $true)][AllowEmptyString()][string]$Text)

    [IO.File]::AppendAllText($logFullPath, $Text.TrimEnd() + [Environment]::NewLine, $utf8NoBom)
    Write-Host $Text.TrimEnd()
}

function Write-Section {
    param([Parameter(Mandatory = $true)][string]$Title)

    Write-LogLine ''
    Write-LogLine ('=' * 78)
    Write-LogLine $Title
    Write-LogLine ('=' * 78)
}

function Assert-Test {
    param(
        [Parameter(Mandatory = $true)][string]$ID,
        [Parameter(Mandatory = $true)][bool]$Condition,
        [Parameter(Mandatory = $true)][string]$Description
    )

    if ($Condition) {
        Write-LogLine "PASS $ID - $Description"
        return
    }
    Write-LogLine "FAIL $ID - $Description"
    throw "Assertion failed: $ID - $Description"
}

if (-not (Test-Path -LiteralPath $buildScript)) {
    throw "Build script not found: $buildScript"
}
if (-not (Test-Path -LiteralPath (Join-Path $moduleRoot 'go.mod'))) {
    throw "Go module not found: $moduleRoot"
}
if (-not (Test-Path -LiteralPath $fixturePath)) {
    throw "Typhoon Bavi fixture not found: $fixturePath"
}
$null = New-Item -ItemType Directory -Force -Path $logDirectory
[IO.File]::WriteAllText($logFullPath, '', $utf8NoBom)

try {
    Write-Section 'TEST METADATA'
    Write-LogLine 'Test ID: TC-BAVI-001'
    Write-LogLine 'Scenario: Synthetic Typhoon Bavi preparedness exercise'
    Write-LogLine 'Status: SYNTHETIC ONLY - NOT A REAL FORECAST OR WARNING'
    Write-LogLine "Started UTC: $([DateTime]::UtcNow.ToString('o'))"
    Write-LogLine "Runner: $($MyInvocation.MyCommand.Path)"
    Write-LogLine "Module: $moduleRoot"
    Write-LogLine "Fixture: $fixturePath"
    Write-LogLine "Log: $logFullPath"

    Write-Section 'BUILD AND VERIFICATION'
    Write-LogLine "Temporary build directory: $temporaryBuild"
    $buildOutput = @(& $buildScript -Action build -OutputDirectory $temporaryBuild 6>&1 5>&1 4>&1 3>&1 2>&1)
    foreach ($line in $buildOutput) {
        Write-LogLine ("BUILD> " + [string]$line)
    }

    $manifestPath = Join-Path $temporaryBuild 'build-manifest.json'
    if (-not (Test-Path -LiteralPath $manifestPath)) {
        throw "Build manifest not found: $manifestPath"
    }
    $manifestText = Get-Content -LiteralPath $manifestPath -Raw -Encoding UTF8
    $manifest = $manifestText | ConvertFrom-Json
    $binaryPath = Join-Path $temporaryBuild $manifest.binary
    if (-not (Test-Path -LiteralPath $binaryPath)) {
        throw "Built executable not found: $binaryPath"
    }
    $actualHash = (Get-FileHash -LiteralPath $binaryPath -Algorithm SHA256).Hash.ToLowerInvariant()
    Write-LogLine "Binary: $binaryPath"
    Write-LogLine "Application version: $($manifest.application_version)"
    Write-LogLine "Go toolchain: $($manifest.go_version)"
    Write-LogLine "Manifest SHA-256: $($manifest.sha256)"
    Write-LogLine "Actual SHA-256:   $actualHash"

    Write-Section 'INPUT - COMPLETE JSON REQUEST'
    $inputText = Get-Content -LiteralPath $fixturePath -Raw -Encoding UTF8
    $request = $inputText | ConvertFrom-Json
    Write-LogBlock $inputText

    Write-Section 'PROCESS - OBSERVABLE DETERMINISTIC EVIDENCE'
    Write-LogLine 'P0 Session/build integrity: verified source build, tests, vet, executable version, and SHA-256.'
    Write-LogLine "P1 Input scope: capability=$($request.capability_id); domain=$($request.declared_domain); data_class=$($request.data_class); synthetic=$($request.synthetic)."
    Write-LogLine "P2 Permission gate: memory=$($request.memory_permission); remote=$($request.remote_permission); external execution absent."
    Write-LogLine "P3 Warning input: official_status=$($request.warning.official_status); area_match=$($request.warning.area_match); urgency=$($request.warning.urgency); severity=$($request.warning.severity); certainty=$($request.warning.certainty)."
    Write-LogLine "P4 Evidence input: tier=$($request.warning.evidence_tier); integrity=$($request.warning.integrity_status); conflict=$($request.warning.source_conflict); stale=$($request.warning.stale); direct_danger=$($request.warning.direct_danger)."

    $rawOutputLines = @(& $binaryPath -input $fixturePath 2>&1)
    $backendExitCode = $LASTEXITCODE
    $rawOutput = $rawOutputLines -join [Environment]::NewLine
    if ($backendExitCode -ne 0) {
        Write-LogBlock $rawOutput
        throw "Backend exited with code $backendExitCode"
    }
    $packet = $rawOutput | ConvertFrom-Json

    Write-LogLine "P5 Deterministic observations: count=$(@($packet.observations).Count)."
    foreach ($observation in $packet.observations) {
        Write-LogLine "  OBSERVED> $($observation.claim) [$($observation.source); $($observation.status)]"
    }
    Write-LogLine "P6 Decision classification: class=$($packet.decision_class); mode=$($packet.response_mode)."
    Write-LogLine "P7 Warning evaluation: level=$($packet.warning.level) - $($packet.warning.level_label); posture=$($packet.warning.decision_posture)."
    Write-LogLine "P8 Option construction: count=$(@($packet.options).Count)."
    foreach ($option in $packet.options) {
        Write-LogLine "  OPTION> $($option.option_id); reversibility=$($option.reversibility); external_effect=$($option.external_effect); $($option.summary)"
    }
    Write-LogLine "P9 Advisory selection: status=$($packet.recommendation.status); option=$($packet.recommendation.option_id)."
    Write-LogLine "P10 AI DNA runtime checks: count=$(@($packet.ai_dna_runtime_checks).Count)."
    foreach ($check in $packet.ai_dna_runtime_checks) {
        Write-LogLine "  AI-DNA> $($check.dimension)=$($check.status); $($check.basis)"
    }
    Write-LogLine "P11 Data receipt: memory_used=$($packet.data_receipt.memory_used); remote_processing=$($packet.data_receipt.remote_processing); synthetic=$($packet.data_receipt.synthetic)."
    $userDecisionState = if ($null -eq $packet.user_decision) { 'null' } else { [string]$packet.user_decision }
    Write-LogLine "P12 Human authority: user_decision=$userDecisionState; blocked_actions=$(@($packet.blocked_actions).Count)."

    Write-Section 'ASSERTIONS'
    Assert-Test 'BAVI-A01' (-not [bool]$manifest.tests_skipped) 'Build tests were not skipped.'
    Assert-Test 'BAVI-A02' ([int]$manifest.external_modules -eq 0) 'No external Go modules are present.'
    Assert-Test 'BAVI-A03' ($packet.capability_id -eq 'assist.personal-decision.v1') 'Capability ID is exact.'
    Assert-Test 'BAVI-A04' ($packet.decision_class -eq 'D2') 'Decision class is D2.'
    Assert-Test 'BAVI-A05' ($packet.response_mode -eq 'PREPARE') 'Response mode is PREPARE.'
    Assert-Test 'BAVI-A06' ($packet.warning.level -eq 'W2') 'Warning level is W2.'
    Assert-Test 'BAVI-A07' ($packet.warning.level_label -eq 'PREPARE') 'Warning label is PREPARE.'
    Assert-Test 'BAVI-A08' ($packet.warning.decision_posture -eq 'PREPARE NOW') 'Decision posture is PREPARE NOW.'
    Assert-Test 'BAVI-A09' ($packet.warning.evidence_tier -eq 'V2') 'Evidence tier remains V2.'
    Assert-Test 'BAVI-A10' ($packet.warning.official_status -eq 'none_found') 'Official status remains none_found.'
    Assert-Test 'BAVI-A11' ($packet.recommendation.option_id -eq 'prepare-now') 'Recommendation selects prepare-now.'
    Assert-Test 'BAVI-A12' (@($packet.ai_dna_runtime_checks).Count -eq 9) 'Nine AI DNA runtime checks are present.'
    Assert-Test 'BAVI-A13' (-not [bool]$packet.data_receipt.remote_processing) 'Remote processing is false.'
    Assert-Test 'BAVI-A14' (-not [bool]$packet.data_receipt.memory_used) 'Persistent memory use is false.'
    Assert-Test 'BAVI-A15' ([bool]$packet.data_receipt.synthetic) 'Data receipt records synthetic input.'
    Assert-Test 'BAVI-A16' (@($packet.blocked_actions) -contains 'broadcast_family_or_public_alert') 'Family/public broadcast is blocked.'
    Assert-Test 'BAVI-A17' (@($packet.blocked_actions) -contains 'control_device_or_robot') 'Device/robot control is blocked.'
    Assert-Test 'BAVI-A18' ($null -eq $packet.user_decision) 'User decision remains null.'
    $limitationsText = @($packet.limitations) -join ' '
    Assert-Test 'BAVI-A19' ($limitationsText -match 'not an alerting or emergency authority') 'Emergency-authority limitation is present.'
    Assert-Test 'BAVI-A20' ($backendExitCode -eq 0) 'Backend exited successfully.'
    Assert-Test 'BAVI-HASH' ($actualHash -eq [string]$manifest.sha256) 'Executable hash matches the build manifest.'

    Write-Section 'OUTPUT - COMPLETE JSON DECISION PACKET'
    Write-LogBlock $rawOutput

    Write-Section 'RESULT'
    Write-LogLine 'RESULT: PASS'
    Write-LogLine 'All mandatory assertions passed. No averaging or override was used.'
    Write-LogLine 'Interpretation: synthetic W2 PREPARE guidance only; no real warning authority or deployment claim.'
}
catch {
    $exitCode = 1
    Write-Section 'RESULT'
    Write-LogLine 'RESULT: FAIL'
    Write-LogLine ("Error: " + $_.Exception.Message)
    Write-LogLine 'A failed mandatory assertion blocks acceptance of this test result.'
}
finally {
    Write-Section 'CLEANUP'
    $temporaryResolved = [IO.Path]::GetFullPath($temporaryBuild)
    $systemTempResolved = [IO.Path]::GetFullPath([IO.Path]::GetTempPath())
    if ((Test-Path -LiteralPath $temporaryResolved) -and $temporaryResolved.StartsWith($systemTempResolved, [StringComparison]::OrdinalIgnoreCase)) {
        Remove-Item -LiteralPath $temporaryResolved -Recurse -Force
    }
    Write-LogLine "Temporary build removed: $(-not (Test-Path -LiteralPath $temporaryResolved))"
    Write-LogLine "Finished UTC: $([DateTime]::UtcNow.ToString('o'))"
    Write-LogLine "Final exit code: $exitCode"
}

exit $exitCode
