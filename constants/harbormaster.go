package constants

// HarbormasterArtifactType are the types of artifact that can be attached through Harbormaster.
type HarbormasterArtifactType string

const (
	// HarbormasterHost references a host lease from Drydock.
	HarbormasterHost HarbormasterArtifactType = "host"
	// HarbormasterWorkingCopy references a working copy lease from Drydock.
	HarbormasterWorkingCopy = "working-copy"
	// HarbormasterFile stores a reference to file data which has been uploaded to Phabricator.
	HarbormasterFile = "file"
	// HarbormasterURI stores a URI.
	HarbormasterURI = "uri"
)

// HarbormasterMessageType are the types of message that can be sent to sendmessage.
type HarbormasterMessageType string

const (
	HarbormasterPass HarbormasterMessageType = "pass"
	HarbormasterFail HarbormasterMessageType = "fail"
	HarbormasterWork HarbormasterMessageType = "work"
)

// HarbormasterUnitResult are the types of unit test result.
type HarbormasterUnitResult string

const (
	HarbormasterUnitPass    HarbormasterUnitResult = "pass"
	HarbormasterUnitFail    HarbormasterUnitResult = "fail"
	HarbormasterUnitSkip    HarbormasterUnitResult = "skip"
	HarbormasterUnitBroken  HarbormasterUnitResult = "broken"
	HarbormasterUnitUnsound HarbormasterUnitResult = "unsound"
)

// HarbormasterUnitFormat are the types of format unit tests take details in.
type HarbormasterUnitFormat string

const (
	HarbormasterFormatText     HarbormasterUnitFormat = "text"
	HarbormasterFormatRemarkup HarbormasterUnitFormat = "remarkup"
)

// HarbormasterLintSeverity are the accepted severity levels for lint errors.
type HarbormasterLintSeverity string

const (
	HarbormasterSeverityAdvice   HarbormasterLintSeverity = "advice"
	HarbormasterSeverityAutofix  HarbormasterLintSeverity = "autofix"
	HarbormasterSeverityWarning  HarbormasterLintSeverity = "warning"
	HarbormasterSeverityError    HarbormasterLintSeverity = "error"
	HarbormasterSeverityDisabled HarbormasterLintSeverity = "disabled"
)
