package workflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureMainJobContentsRead(t *testing.T) {
	t.Run("creates default contents read permissions when empty", func(t *testing.T) {
		result := ensureMainJobContentsRead("", true)
		assert.Equal(t, NewPermissionsContentsRead().RenderToYAML(), result)
	})

	t.Run("adds contents read when required", func(t *testing.T) {
		permissions := "permissions:\n  issues: read\n"

		result := ensureMainJobContentsRead(permissions, true)

		perms := NewPermissionsParser(result).ToPermissions()
		level, exists := perms.Get(PermissionContents)
		assert.True(t, exists)
		assert.Equal(t, PermissionRead, level)
	})

	t.Run("leaves permissions unchanged when not required", func(t *testing.T) {
		permissions := "permissions:\n  issues: read\n"

		result := ensureMainJobContentsRead(permissions, false)

		assert.Equal(t, permissions, result)
	})
}
