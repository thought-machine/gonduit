package entities

import "github.com/samwestmoreland/gonduit/constants"

// A UnitResult represents the result of a unit test, as sent to Harbormaster.
type UnitResult struct {
	Name      string                           `json:"name"`
	Result    constants.HarbormasterUnitResult `json:"result"`
	Namespace string                           `json:"namespace,omitempty"`
	Engine    string                           `json:"engine,omitempty"`
	Duration  float64                          `json:"duration"`
	Path      string                           `json:"path,omitempty"`
	Coverage  map[string]string                `json:"coverage,omitempty"`
	Details   string                           `json:"details,omitempty"`
	Format    constants.HarbormasterUnitFormat `json:"format"`
}

// A LintResult represents the result of a lint warning, as sent to Harbormaster.
type LintResult struct {
	Name        string                             `json:"name"`
	Code        string                             `json:"code"`
	Severity    constants.HarbormasterLintSeverity `json:"severity"`
	Path        string                             `json:"path"`
	Line        int                                `json:"line"`
	Char        int                                `json:"char"`
	Description string                             `json:"description"`
}
