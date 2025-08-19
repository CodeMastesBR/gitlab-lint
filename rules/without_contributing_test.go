package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutContributing(t *testing.T) {
	rule := WithoutContributing{}

	t.Run("quando o projeto tem arquivo CONTRIBUTING.md", func(t *testing.T) {
		project := mockProject()
		project.ContributingURL = "https://gitlab.com/example/project/-/blob/master/CONTRIBUTING.md"
		
		result := rule.Run(project)
		assert.False(t, result.Failed)
		assert.Empty(t, result.Message)
	})

	t.Run("quando o projeto não tem arquivo CONTRIBUTING.md", func(t *testing.T) {
		project := mockProject()
		project.ContributingURL = ""
		
		result := rule.Run(project)
		assert.True(t, result.Failed)
		assert.Equal(t, "Projeto não possui arquivo CONTRIBUTING.md", result.Message)
	})
}
