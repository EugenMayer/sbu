package config

import (
	assert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadUserConfig(t *testing.T) {
	config, err := LoadUserConfig("../tests/assets/user-config.yml")
	require.Nil(t, err)
	require.NotNil(t, config)

	assert.Equal(t, config.GlobalConfig.DefaultCredentials.Username, "admin")
	assert.Equal(t, config.GlobalConfig.DefaultCredentials.Password, "verysecret")
}
