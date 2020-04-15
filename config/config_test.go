package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigCases(t *testing.T) {
	v := Viper()
	cases := []struct {
		key      string
		callFunc func() string
	}{
		{
			key:      "log.level",
			callFunc: LogLevel,
		},
	}

	for _, c := range cases {
		assert.Equal(t, v.GetString(c.key), c.callFunc())
	}
}

func TestConfigCasesWithFeatureToggle(t *testing.T) {
	v := Viper()
	cases := []struct {
		key      string
		callFunc func() bool
	}{
		{
			key:      "debug",
			callFunc: DebugModeEnabled,
		},
	}
	for _, c := range cases {
		assert.Equal(t, v.GetBool(c.key), c.callFunc())
	}

	assert.Equal(t, false, v.GetBool("nonexistingkey"))
}
