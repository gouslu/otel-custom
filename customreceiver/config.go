package customreceiver

import (
	"fmt"
	"strconv"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	Port string `mapstructure:"port"`
}

func (cfg *Config) Validate() error {
	port, err := strconv.Atoi(cfg.Port)

	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("'%s' is not a valid port number", cfg.Port)
	}
	return nil
}
