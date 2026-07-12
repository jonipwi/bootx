[CmdletBinding()]
param([string]$LogPath = '')

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

if ([string]::IsNullOrWhiteSpace($LogPath)) {
    $LogPath = Join-Path $PSScriptRoot 'israel-palestine.log'
}

& (Join-Path $PSScriptRoot 'run-governance-analysis-test.ps1') `
    -FixturePath (Join-Path $PSScriptRoot 'governance-testdata\israel-palestine.json') `
    -ExpectedTestID 'TC-DHAI-ILPS-001' `
    -ExpectedCaseType 'israel-palestine' `
    -LogPath $LogPath
exit $LASTEXITCODE

