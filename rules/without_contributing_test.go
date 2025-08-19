package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutContributing(t *testing.T) {
	rule := WithoutContributing{}

	t.Run("when project has CONTRIBUTING.md", func(t *testing.T) {
		project := mockProject()
		project.ContributingFile = "# Contributing Guide"
		
		result := rule.Run(project)
		assert.False(t, result.Failed)
		assert.Empty(t, result.Message)
	})

	t.Run("when project does not have CONTRIBUTING.md", func(t *testing.T) {
		project := mockProject()
		project.ContributingFile = ""
		
		result := rule.Run(project)
		assert.True(t, result.Failed)
		assert.Equal(t, "Project does not have a CONTRIBUTING.md file", result.Message)
	})
}
