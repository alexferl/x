package log

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestConfig_AddFlags(t *testing.T) {
	c := DefaultConfig
	fs := &pflag.FlagSet{}
	c.AddFlags(fs)

	level, _ := fs.GetString("log-level")
	out, _ := fs.GetString("log-output")
	writer, _ := fs.GetString("log-writer")
	assert.Equal(t, DefaultConfig.LogLevel, level)
	assert.Equal(t, DefaultConfig.LogOutput, out)
	assert.Equal(t, DefaultConfig.LogWriter, writer)
}
