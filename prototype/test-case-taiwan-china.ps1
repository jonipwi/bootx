[CmdletBinding()]
param([string]$LogPath = '')

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

if ([string]::IsNullOrWhiteSpace($LogPath)) {
    $LogPath = Join-Path $PSScriptRoot 'taiwan-china.log'
}

& (Join-Path $PSScriptRoot 'run-governance-analysis-test.ps1') `
    -FixturePath (Join-Path $PSScriptRoot 'governance-testdata\taiwan-china.json') `
    -ExpectedTestID 'TC-DHAI-TWC-001' `
    -ExpectedCaseType 'taiwan-china' `
    -LogPath $LogPath
exit $LASTEXITCODE

