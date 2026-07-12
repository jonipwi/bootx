[CmdletBinding()]
param([string]$LogPath = '')

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

if ([string]::IsNullOrWhiteSpace($LogPath)) {
    $LogPath = Join-Path $PSScriptRoot 'ccp-human-rights.log'
}

& (Join-Path $PSScriptRoot 'run-governance-analysis-test.ps1') `
    -FixturePath (Join-Path $PSScriptRoot 'governance-testdata\ccp-human-rights.json') `
    -ExpectedTestID 'TC-DHAI-CCPHR-001' `
    -ExpectedCaseType 'ccp-human-rights' `
    -LogPath $LogPath
exit $LASTEXITCODE

