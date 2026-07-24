[CmdletBinding()]
param(
    [string]$LogPath = ''
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$moduleRoot = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'personal-companion'))
$buildScript = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'build.ps1'))
$abuseFixture = [IO.Path]::GetFullPath((Join-Path $moduleRoot 'testdata\law-clarity-abusive-capital-corruption-proposal.json'))
$revisionFixture = [IO.Path]::GetFullPath((Join-Path $moduleRoot 'testdata\law-clarity-rights-preserving-anti-corruption-revision.json'))
$testDocument = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot 'TEST_CASE_LAW_CLARITY_CAPITAL_CORRUPTION.md'))
if ([string]::IsNullOrWhiteSpace($LogPath)) {
    $LogPath = Join-Path $PSScriptRoot 'law-clarity-capital-corruption.log'
}
$logFullPath = [IO.Path]::GetFullPath($LogPath)
$logDirectory = Split-Path -Parent $logFullPath
$temporaryBuild = Join-Path ([IO.Path]::GetTempPath()) ('bootx-law-clarity-test-' + [Guid]::NewGuid().ToString('N'))
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

function Invoke-LawFixture {
    param(
        [Parameter(Mandatory = $true)][string]$Label,
        [Parameter(Mandatory = $true)][string]$FixturePath,
        [Parameter(Mandatory = $true)][string]$BinaryPath
    )

    Write-Section "INPUT $Label - COMPLETE JSON REQUEST"
    $inputText = Get-Content -LiteralPath $FixturePath -Raw -Encoding UTF8
    Write-LogBlock $inputText

    Write-Section "PROCESS $Label - OBSERVABLE DETERMINISTIC EVIDENCE"
    $request = $inputText | ConvertFrom-Json
    Write-LogLine "P0 Capability: $($request.capability_id)"
    Write-LogLine "P1 Instrument: $($request.instrument_type); context=$($request.jurisdiction)"
    Write-LogLine "P2 Public/non-sensitive confirmation: $($request.public_non_sensitive_confirmed)"
    Write-LogLine 'P3 Formula path: validate -> Q/strict/fairness -> Z/M/trigger -> phrase scan -> bounded disposition.'
    Write-LogLine 'P4 Authority boundary: no legality, guilt, liability, or sentence decision.'

    $rawOutputLines = @(& $BinaryPath -law-input $FixturePath 2>&1)
    $backendExitCode = $LASTEXITCODE
    $rawOutput = $rawOutputLines -join [Environment]::NewLine
    if ($backendExitCode -ne 0) {
        Write-LogBlock $rawOutput
        throw "$Label backend exited with code $backendExitCode"
    }
    $report = $rawOutput | ConvertFrom-Json
    Write-LogLine "P5 Q=$($report.law_quality_score); Z=$($report.gray_zone_risk_score); M=$($report.manipulation_risk_index)."
    Write-LogLine "P6 Strict=$($report.strict_good_law_gate.status); rights=$($report.human_rights_gate.status); trigger=$($report.high_manipulation_trigger)."
    Write-LogLine "P7 Disposition=$($report.disposition); phrase_types=$(@($report.visible_phrase_hits).Count)."
    $userDecisionState = if ($null -eq $report.user_decision) { 'null' } else { [string]$report.user_decision }
    Write-LogLine "P8 Human authority: user_decision=$userDecisionState."

    return [pscustomobject]@{
        Request = $request
        Report = $report
        RawOutput = $rawOutput
        ExitCode = $backendExitCode
    }
}

foreach ($requiredPath in @($buildScript, $abuseFixture, $revisionFixture, $testDocument)) {
    if (-not (Test-Path -LiteralPath $requiredPath)) {
        throw "Required file not found: $requiredPath"
    }
}
$null = New-Item -ItemType Directory -Force -Path $logDirectory
[IO.File]::WriteAllText($logFullPath, '', $utf8NoBom)

try {
    Write-Section 'TEST METADATA'
    Write-LogLine 'Test suite ID: TC-LAW-CAPITAL-CORRUPTION-001'
    Write-LogLine 'Scenario A: Abusive capital-corruption proposal with scapegoating pathway'
    Write-LogLine 'Scenario B: Rights-preserving non-capital anti-corruption revision'
    Write-LogLine 'Status: FICTIONAL EDUCATIONAL EXERCISE - NOT LEGAL ADVICE OR A SENTENCING DECISION'
    Write-LogLine "Started UTC: $([DateTime]::UtcNow.ToString('o'))"
    Write-LogLine "Abuse fixture: $abuseFixture"
    Write-LogLine "Revision fixture: $revisionFixture"
    Write-LogLine "Test document: $testDocument"
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
    Write-LogLine "Capabilities: $(@($manifest.capabilities) -join ', ')"
    Write-LogLine "Manifest SHA-256: $($manifest.sha256)"
    Write-LogLine "Actual SHA-256:   $actualHash"

    $abuse = Invoke-LawFixture -Label 'CASE A' -FixturePath $abuseFixture -BinaryPath $binaryPath
    $revision = Invoke-LawFixture -Label 'CASE B' -FixturePath $revisionFixture -BinaryPath $binaryPath

    Write-Section 'ASSERTIONS - BUILD AND SHARED BOUNDARIES'
    Assert-Test 'LAW-S01' (-not [bool]$manifest.tests_skipped) 'Build tests and vet were not skipped.'
    Assert-Test 'LAW-S02' ([int]$manifest.external_modules -eq 0) 'No external Go modules are present.'
    Assert-Test 'LAW-S03' (@($manifest.capabilities) -contains 'assist.law-clarity.v1') 'Manifest declares the Law Clarity capability.'
    Assert-Test 'LAW-S04' ($actualHash -eq [string]$manifest.sha256) 'Executable hash matches the manifest.'
    foreach ($sharedCase in @(
        [pscustomobject]@{ ID = 'A'; Value = $abuse },
        [pscustomobject]@{ ID = 'B'; Value = $revision }
    )) {
        $case = $sharedCase.Value
        $id = $sharedCase.ID
        Assert-Test "LAW-S${id}05" ($case.Report.capability_id -eq 'assist.law-clarity.v1') "Case $id capability ID is exact."
        Assert-Test "LAW-S${id}06" (@($case.Report.ai_dna_runtime_checks).Count -eq 9) "Case $id has nine AI DNA runtime checks."
        Assert-Test "LAW-S${id}07" (-not [bool]$case.Report.input_receipt.remote_processing) "Case $id remote processing is false."
        Assert-Test "LAW-S${id}08" (-not [bool]$case.Report.input_receipt.persistent_memory) "Case $id persistent memory is false."
        Assert-Test "LAW-S${id}09" ($null -eq $case.Report.user_decision) "Case $id user decision remains null."
        Assert-Test "LAW-S${id}10" (@($case.Report.blocked_conclusions) -contains 'legal validity or constitutionality') "Case $id blocks legal validity and constitutionality conclusions."
        Assert-Test "LAW-S${id}11" (@($case.Report.blocked_conclusions) -contains 'guilt, liability, eligibility, or entitlement') "Case $id blocks guilt and liability conclusions."
        Assert-Test "LAW-S${id}12" (@($case.Report.blocked_conclusions) -contains 'authorization to enforce, punish, detain, discriminate, or deny rights') "Case $id blocks enforcement and punishment authority."
    }

    Write-Section 'ASSERTIONS - CASE A ABUSIVE PROPOSAL'
    Assert-Test 'LAW-A01' ([double]$abuse.Report.law_quality_score -eq 12.0) 'Q equals 12.00.'
    Assert-Test 'LAW-A02' ([double]$abuse.Report.gray_zone_risk_score -eq 92.75) 'Z equals 92.75.'
    Assert-Test 'LAW-A03' ([double]$abuse.Report.manipulation_risk_index -eq 92.75) 'M equals 92.75 and remains an index, not a probability.'
    Assert-Test 'LAW-A04' ($abuse.Report.strict_good_law_gate.status -eq 'FAIL') 'Strict gate fails.'
    Assert-Test 'LAW-A05' ($abuse.Report.human_rights_gate.status -eq 'FAIL') 'Human-rights fairness gate fails.'
    Assert-Test 'LAW-A06' ([bool]$abuse.Report.high_manipulation_trigger) 'High-manipulation trigger fires.'
    Assert-Test 'LAW-A07' ($abuse.Report.disposition -eq 'FUNDAMENTAL_REVISION_REQUIRED') 'Fundamental revision is required.'
    $abusePhrases = @($abuse.Report.visible_phrase_hits | ForEach-Object { $_.phrase })
    Assert-Test 'LAW-A08' ($abusePhrases.Count -eq 6) 'Six configured phrase types are found.'
    $phraseIndex = 1
    foreach ($phrase in @('appropriate', 'improper', 'disturbing', 'public interest', 'as necessary', 'may take action')) {
        Assert-Test ("LAW-A09-" + $phraseIndex) ($abusePhrases -contains $phrase) "Phrase prompt is present: $phrase."
        $phraseIndex++
    }
    Assert-Test 'LAW-A10' ($abuse.Request.clause_text -match 'higher-ranking organizers') 'Threat model records concealment concerning higher-ranking organizers.'
    Assert-Test 'LAW-A11' ($abuse.Request.clause_text -match 'political opponent') 'Threat model records political-targeting risk.'
    Assert-Test 'LAW-A12' ($abuse.Request.clause_text -match 'presumed responsible unless') 'Threat model records reversed burden of proof.'

    Write-Section 'ASSERTIONS - CASE B RIGHTS-PRESERVING REVISION'
    Assert-Test 'LAW-B01' ([double]$revision.Report.law_quality_score -eq 92.4) 'Q equals 92.40.'
    Assert-Test 'LAW-B02' ([double]$revision.Report.gray_zone_risk_score -eq 7.7) 'Z equals 7.70.'
    Assert-Test 'LAW-B03' ([double]$revision.Report.manipulation_risk_index -eq 0.01) 'M rounds to 0.01 and remains an index, not proof of zero risk.'
    Assert-Test 'LAW-B04' ($revision.Report.strict_good_law_gate.status -eq 'PASS') 'Strict screening gate passes.'
    Assert-Test 'LAW-B05' ($revision.Report.human_rights_gate.status -eq 'PASS') 'Human-rights fairness screening gate passes.'
    Assert-Test 'LAW-B06' (-not [bool]$revision.Report.high_manipulation_trigger) 'High-manipulation trigger does not fire.'
    Assert-Test 'LAW-B07' ($revision.Report.disposition -eq 'QUALIFIED_REVIEW_REQUIRED') 'Qualified review remains required; no legal approval is produced.'
    Assert-Test 'LAW-B08' (@($revision.Report.visible_phrase_hits).Count -eq 0) 'No configured literal phrase is found.'
    Assert-Test 'LAW-B09' ($revision.Request.clause_text -match 'No death penalty may be imposed for a corruption offense') 'Capital punishment for corruption is expressly rejected.'
    Assert-Test 'LAW-B10' ($revision.Request.clause_text -match 'No mandatory life sentence applies') 'Mandatory life imprisonment is expressly rejected.'
    Assert-Test 'LAW-B11' ($revision.Request.clause_text -match 'predicted conduct does not by itself establish criminal responsibility') 'Prediction cannot establish responsibility.'
    Assert-Test 'LAW-B12' ($revision.Request.clause_text -match 'regardless of office or rank') 'The same offense elements and procedures apply regardless of rank.'
    Assert-Test 'LAW-B13' ($revision.Request.clause_text -match 'presumption of innocence') 'Presumption of innocence is explicit.'
    Assert-Test 'LAW-B14' ($revision.Request.clause_text -match 'inculpatory and exculpatory evidence') 'Both inculpatory and exculpatory evidence must be disclosed.'
    Assert-Test 'LAW-B15' ($revision.Request.clause_text -match 'rehabilitation and reintegration') 'Long-custody review preserves rehabilitation and reintegration.'
    Assert-Test 'LAW-B16' ($revision.Request.clause_text -match 'does not erase the conviction, restitution, asset recovery') 'Repentance or rehabilitation does not erase accountability and recovery.'

    Write-Section 'OUTPUT CASE A - COMPLETE JSON SCREENING REPORT'
    Write-LogBlock $abuse.RawOutput
    Write-Section 'OUTPUT CASE B - COMPLETE JSON SCREENING REPORT'
    Write-LogBlock $revision.RawOutput

    Write-Section 'RESULT'
    Write-LogLine 'RESULT: PASS'
    Write-LogLine 'Case A exposed the abusive proposal and required fundamental revision.'
    Write-LogLine 'Case B passed the research screening gates but still required qualified review.'
    Write-LogLine 'No legality, guilt, death, life-imprisonment, release, or enforcement decision was made.'
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
