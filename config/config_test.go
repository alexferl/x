package config

import (
	"testing"

	xlog "github.com/alexferl/x/log"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()

	assert.Equal(t, DefaultConfig.AppName, c.AppName)
	assert.Equal(t, DefaultConfig.EnvName, c.EnvName)
	assert.Equal(t, DefaultConfig.EnvVarPrefix, c.EnvVarPrefix)
	assert.Equal(t, DefaultConfig.Logging, xlog.DefaultConfig)
}
