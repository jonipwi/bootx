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
        application_version = $appVersion
        build_utc = [DateTime]::UtcNow.ToString('o')
        go_version = $goVersion
        goos = $goos
        goarch = $goarch
        binary = $binaryName
        sha256 = $hash
        tests_skipped = [bool]$SkipTests
        external_modules = $externalModuleCount
        safety_status = 'DEV-1 deterministic prototype; not deployed or safety certified'
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
