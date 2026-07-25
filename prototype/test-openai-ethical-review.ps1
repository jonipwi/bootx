[CmdletBinding()]
param(
    [string]$InputFile = '',
    [string]$Model = '',
    [switch]$ShowFullOutput
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$moduleRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'personal-companion'))
$repositoryRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot '..'))
if ([string]::IsNullOrWhiteSpace($InputFile)) {
    $InputFile = Join-Path $moduleRoot 'testdata\ethical-review-synthetic-publication.json'
}
$inputPath = [IO.Path]::GetFullPath($InputFile)
if (-not (Test-Path -LiteralPath $inputPath -PathType Leaf)) {
    throw "Review input not found: $inputPath"
}
if ([string]::IsNullOrWhiteSpace($env:OPENAI_API_KEY)) {
    throw 'OPENAI_API_KEY is not available to this PowerShell process. Never put the key in a fixture or command-line argument.'
}

& (Join-Path $PSScriptRoot 'build.ps1') -Action build
if ($LASTEXITCODE -ne 0) {
    throw "BootX build failed with exit code $LASTEXITCODE"
}

$goos = (& go env GOOS).Trim()
$goarch = (& go env GOARCH).Trim()
$extension = if ($goos -eq 'windows') { '.exe' } else { '' }
$binary = Join-Path $moduleRoot "dist\bootx-companion-$goos-$goarch$extension"

$arguments = @('-review-input', $inputPath)
if (-not [string]::IsNullOrWhiteSpace($Model)) {
    $arguments += @('-openai-model', $Model)
}
$jsonLines = @(& $binary @arguments)
if ($LASTEXITCODE -ne 0) {
    throw "OpenAI ethical-review smoke test failed with exit code $LASTEXITCODE"
}
$jsonText = $jsonLines -join [Environment]::NewLine
$result = $jsonText | ConvertFrom-Json

if ($result.capability_id -ne 'assist.ethical-review.v1') {
    throw "Unexpected capability: $($result.capability_id)"
}
if ([bool]$result.remote_receipt.store_requested -or
    [bool]$result.remote_receipt.tools_enabled -or
    [bool]$result.remote_receipt.external_actions_enabled -or
    [bool]$result.remote_receipt.application_persistence) {
    throw 'Remote safety receipt failed: storage, tools, actions, or application persistence was enabled.'
}
if ($null -ne $result.user_decision) {
    throw 'Human-authority invariant failed: user_decision must remain null.'
}

Write-Host 'BootX OpenAI ethical-review smoke test passed.' -ForegroundColor Green
Write-Host "  Model requested: $($result.remote_receipt.model_requested)"
Write-Host "  Model returned:  $($result.remote_receipt.model_returned)"
Write-Host "  API status:      $($result.openai_advisory.status)"
Write-Host "  Warning:         $($result.deterministic_preflight.warning_level)"
Write-Host "  Local risk index:$($result.deterministic_preflight.review_risk_index) (review-priority index, not probability)"
if ($null -ne $result.openai_advisory.advice) {
    Write-Host "  AI posture:      $($result.openai_advisory.advice.suggested_posture)"
}
Write-Host '  Tools/actions:   disabled'
Write-Host '  Human decision:  required'

if ($ShowFullOutput) {
    Write-Output $jsonText
}
