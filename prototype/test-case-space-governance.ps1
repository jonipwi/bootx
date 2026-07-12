[CmdletBinding()]
param(
    [string]$LogPath = ''
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$repositoryRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot '..'))
$moduleRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'personal-companion'))
$buildScript = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'build.ps1'))
$fixturePath = [IO.Path]::GetFullPath((Join-Path $moduleRoot 'testdata\space-governance-readiness-exercise.json'))
$documentPath = [IO.Path]::GetFullPath((Join-Path $repositoryRoot 'docs\research\space\civilization-governance-before-space-expansion.md'))
if ([string]::IsNullOrWhiteSpace($LogPath)) {
    $LogPath = Join-Path $PSScriptRoot 'space-governance.log'
}
$logFullPath = [IO.Path]::GetFullPath($LogPath)
$logDirectory = Split-Path -Parent $logFullPath
$temporaryBuild = Join-Path ([IO.Path]::GetTempPath()) ('bootx-space-governance-test-' + [Guid]::NewGuid().ToString('N'))
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

foreach ($requiredPath in @($buildScript, $fixturePath, $documentPath)) {
    if (-not (Test-Path -LiteralPath $requiredPath)) {
        throw "Required file not found: $requiredPath"
    }
}
$null = New-Item -ItemType Directory -Force -Path $logDirectory
[IO.File]::WriteAllText($logFullPath, '', $utf8NoBom)

try {
    Write-Section 'TEST METADATA'
    Write-LogLine 'Test ID: TC-SPACE-GOV-001'
    Write-LogLine 'Scenario: Fictional permanent lunar settlement governance review'
    Write-LogLine 'Status: SYNTHETIC ONLY - NOT LEGAL ADVICE OR MISSION APPROVAL'
    Write-LogLine "Started UTC: $([DateTime]::UtcNow.ToString('o'))"
    Write-LogLine "Document: $documentPath"
    Write-LogLine "Fixture: $fixturePath"
    Write-LogLine "Log: $logFullPath"

    Write-Section 'BUILD AND VERIFICATION'
    Write-LogLine "Temporary build directory: $temporaryBuild"
    $buildOutput = @(& $buildScript -Action build -OutputDirectory $temporaryBuild 6>&1 5>&1 4>&1 3>&1 2>&1)
    foreach ($line in $buildOutput) {
        Write-LogLine ('BUILD> ' + [string]$line)
    }
    $manifestPath = Join-Path $temporaryBuild 'build-manifest.json'
    if (-not (Test-Path -LiteralPath $manifestPath)) {
        throw "Build manifest not found: $manifestPath"
    }
    $manifest = Get-Content -LiteralPath $manifestPath -Raw -Encoding UTF8 | ConvertFrom-Json
    $binaryPath = Join-Path $temporaryBuild $manifest.binary
    if (-not (Test-Path -LiteralPath $binaryPath)) {
        throw "Built executable not found: $binaryPath"
    }
    $actualHash = (Get-FileHash -LiteralPath $binaryPath -Algorithm SHA256).Hash.ToLowerInvariant()
    Write-LogLine "Application version: $($manifest.application_version)"
    Write-LogLine "Manifest SHA-256: $($manifest.sha256)"
    Write-LogLine "Actual SHA-256:   $actualHash"

    Write-Section 'DOCUMENT PROCESS - VISIBLE POLICY EVIDENCE'
    $documentText = Get-Content -LiteralPath $documentPath -Raw -Encoding UTF8
    Write-LogLine "Document bytes: $($utf8NoBom.GetByteCount($documentText))"
    Write-LogLine "Document SHA-256: $((Get-FileHash -LiteralPath $documentPath -Algorithm SHA256).Hash.ToLowerInvariant())"
    Write-LogLine 'D0 Check evidence maturity and fact/proposal separation.'
    Write-LogLine 'D1 Check all six mandatory non-compensable gates.'
    Write-LogLine 'D2 Check rejection of empirical likelihood and readiness-certification claims.'
    Write-LogLine 'D3 Check human rights, bounded AI, pluralism, and official references.'

    Write-Section 'INPUT - COMPLETE JSON REQUEST'
    $inputText = Get-Content -LiteralPath $fixturePath -Raw -Encoding UTF8
    $request = $inputText | ConvertFrom-Json
    Write-LogBlock $inputText

    Write-Section 'MVP PROCESS - OBSERVABLE DETERMINISTIC EVIDENCE'
    Write-LogLine "P0 Scope: capability=$($request.capability_id); domain=$($request.declared_domain); synthetic=$($request.synthetic)."
    Write-LogLine "P1 Permissions: memory=$($request.memory_permission); remote=$($request.remote_permission); external execution absent."
    Write-LogLine 'P2 Boundary expectation: legal/governance certification requires qualified multidisciplinary judgment.'
    $rawOutputLines = @(& $binaryPath -input $fixturePath 2>&1)
    $backendExitCode = $LASTEXITCODE
    $rawOutput = $rawOutputLines -join [Environment]::NewLine
    if ($backendExitCode -ne 0) {
        Write-LogBlock $rawOutput
        throw "Backend exited with code $backendExitCode"
    }
    $packet = $rawOutput | ConvertFrom-Json
    Write-LogLine "P3 Classification: class=$($packet.decision_class); mode=$($packet.response_mode)."
    Write-LogLine "P4 Recommendation: status=$($packet.recommendation.status); option=$($packet.recommendation.option_id)."
    Write-LogLine "P5 AI DNA checks: count=$(@($packet.ai_dna_runtime_checks).Count)."
    Write-LogLine "P6 External actions blocked: count=$(@($packet.blocked_actions).Count)."
    $userDecisionState = if ($null -eq $packet.user_decision) { 'null' } else { [string]$packet.user_decision }
    Write-LogLine "P7 Human authority: user_decision=$userDecisionState."

    Write-Section 'ASSERTIONS - DOCUMENTATION'
    Assert-Test 'SPACE-D01' ($documentText -match 'Evidence maturity:\*\* `E1 - Hypothesis`') 'Central thesis is E1 - Hypothesis.'
    Assert-Test 'SPACE-D02' ($documentText -match '### 3\.1 Established foundation') 'Established sources have a separate section.'
    Assert-Test 'SPACE-D03' ($documentText -match '### 3\.2 BootX hypotheses and proposals') 'BootX hypotheses and proposals have a separate section.'
    foreach ($gate in 1..6) {
        Assert-Test ("SPACE-D0" + ($gate + 3)) ($documentText -match ("### Gate G" + $gate + " -")) "Mandatory gate G$gate is present."
    }
    Assert-Test 'SPACE-D10' ($documentText -match 'not empirical likelihood estimates') 'Risk priorities are not presented as empirical likelihoods.'
    Assert-Test 'SPACE-D11' ($documentText -match 'not a certification score') 'Readiness profile rejects certification use.'
    Assert-Test 'SPACE-D12' ($documentText -match 'does \*\*not\*\* adopt those thresholds') 'Earlier 0-to-36 thresholds are explicitly not adopted.'
    Assert-Test 'SPACE-D13' (($documentText -match '`PASS`') -and ($documentText -match '`CONDITIONAL`') -and ($documentText -match '`FAIL`') -and ($documentText -match '`NOT ESTABLISHED`')) 'Four gate evidence states are present.'
    Assert-Test 'SPACE-D14' ($documentText -match 'may not be the final sovereign') 'AI cannot be final sovereign or unreviewable controller.'
    Assert-Test 'SPACE-D15' ($documentText -match 'must not claim divine authorization') 'Creator-centered belief is separated from divine policy claims.'
    Assert-Test 'SPACE-D16' (($documentText -match 'unoosa\.org') -and ($documentText -match 'nasa\.gov')) 'Official UNOOSA and NASA references are present.'

    Write-Section 'ASSERTIONS - MVP'
    Assert-Test 'SPACE-A01' (-not [bool]$manifest.tests_skipped) 'Build tests were not skipped.'
    Assert-Test 'SPACE-A02' ([int]$manifest.external_modules -eq 0) 'No external Go modules are present.'
    Assert-Test 'SPACE-A03' ($packet.capability_id -eq 'assist.personal-decision.v1') 'Capability ID is exact.'
    Assert-Test 'SPACE-A04' ($packet.decision_class -eq 'D3') 'Decision class is D3.'
    Assert-Test 'SPACE-A05' ($packet.response_mode -eq 'ABSTAIN') 'Response mode is ABSTAIN.'
    Assert-Test 'SPACE-A06' ($packet.recommendation.option_id -eq 'prepare-qualified-review') 'Recommendation selects qualified review.'
    Assert-Test 'SPACE-A07' ($packet.recommendation.basis -match 'requires qualified judgment') 'Recommendation states the competence boundary.'
    $warningProperty = $packet.PSObject.Properties['warning']
    Assert-Test 'SPACE-A08' ($null -eq $warningProperty) 'No disaster warning card is generated.'
    $externalEffects = @($packet.options | ForEach-Object { $_.external_effect } | Select-Object -Unique)
    Assert-Test 'SPACE-A09' (($externalEffects.Count -eq 1) -and ($externalEffects[0] -eq 'none')) 'All options have no external effect.'
    Assert-Test 'SPACE-A10' (@($packet.ai_dna_runtime_checks).Count -eq 9) 'Nine AI DNA runtime checks are present.'
    Assert-Test 'SPACE-A11' (-not [bool]$packet.data_receipt.remote_processing) 'Remote processing is false.'
    Assert-Test 'SPACE-A12' (-not [bool]$packet.data_receipt.memory_used) 'Persistent memory use is false.'
    Assert-Test 'SPACE-A13' ([bool]$packet.data_receipt.synthetic) 'Data receipt records synthetic input.'
    Assert-Test 'SPACE-A14' (@($packet.blocked_actions) -contains 'control_device_or_robot') 'Device and robot control is blocked.'
    Assert-Test 'SPACE-A15' (@($packet.blocked_actions) -contains 'broadcast_family_or_public_alert') 'Family and public broadcast is blocked.'
    Assert-Test 'SPACE-A16' ($null -eq $packet.user_decision) 'User decision remains null.'
    Assert-Test 'SPACE-A17' ($actualHash -eq [string]$manifest.sha256) 'Executable hash matches the manifest.'
    Assert-Test 'SPACE-A18' ($backendExitCode -eq 0) 'Backend exited successfully.'

    Write-Section 'OUTPUT - COMPLETE JSON DECISION PACKET'
    Write-LogBlock $rawOutput

    Write-Section 'RESULT'
    Write-LogLine 'RESULT: PASS'
    Write-LogLine 'All mandatory documentation and MVP assertions passed. No score, averaging, or override was used.'
    Write-LogLine 'Interpretation: the prototype abstained; this is not legal validation, mission approval, or civilization certification.'
}
catch {
    $exitCode = 1
    Write-Section 'RESULT'
    Write-LogLine 'RESULT: FAIL'
    Write-LogLine ('Error: ' + $_.Exception.Message)
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
