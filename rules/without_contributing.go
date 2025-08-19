package rules

// WithoutContributing verifica se o projeto possui um arquivo CONTRIBUTING.md
type WithoutContributing struct{}

// Run executa a verificação da regra
func (r WithoutContributing) Run(p Project) RuleResult {
	if p.ContributingURL == "" {
		return RuleResult{
			Failed:  true,
			Message: "Projeto não possui arquivo CONTRIBUTING.md",
		}
	}

	return RuleResult{Failed: false}
}

// GetLevel retorna o nível da regra
func (r WithoutContributing) GetLevel() string {
	return Warning
}

// GetTitle retorna o título da regra
func (r WithoutContributing) GetTitle() string {
	return "Sem Guia de Contribuição"
}
