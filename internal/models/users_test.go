package models

import (
	"testing"

	"github.com/jagfiend/snippetbox/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			"valid id",
			1,
			true,
		},
		{
			"zero id",
			0,
			false,
		},
		{
			"non-existent id",
			42,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDb(t)

			m := UserModel{db}

			exists, err := m.Exists(tt.userID)

			assert.Equal(t, tt.want, exists)
			assert.NilError(t, err)
		})
	}
}
