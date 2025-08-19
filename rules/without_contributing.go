package rules

// WithoutContributing checks if the project has a CONTRIBUTING.md file
type WithoutContributing struct{}

// Run executes the rule check
func (r WithoutContributing) Run(p Project) RuleResult {
	if p.ContributingFile == "" {
		return RuleResult{
			Failed:  true,
			Message: "Project does not have a CONTRIBUTING.md file",
		}
	}

	return RuleResult{Failed: false}
}

// GetLevel returns the rule level
func (r WithoutContributing) GetLevel() string {
	return Warning
}

// GetTitle returns the rule title
func (r WithoutContributing) GetTitle() string {
	return "Without Contributing Guide"
}
