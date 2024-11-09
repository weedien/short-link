package config

import (
	"testing"
)

func TestConfigInitialization(t *testing.T) {

	//_ = os.Setenv("CONFIG_NAME", "dev_configs")
	//_ = os.Setenv("ENV", "dev")

	Init()

	t.Logf("config: %+v", Get())
}
