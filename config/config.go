package config

import (
	"os"
	"time"
)

const (
	DefaultAddress   = "localhost"
	DefaulftPassword = "minecraft"
	DefaultPort      = 25575

	DefaultTimeout time.Duration = time.Second * 7

	DefaultAddressEnv  = "DEFAULT_RCON_ADDRESS"
	DefaultPasswordEnv = "DEFAULT_RCON_PASSWORD"
	DefaltTimoutEnv    = "DEFAULT_RCON_TIMEOUT"

	Version = "1.4.0"
)

// returns the default address set by ether the config or an environment variable
func GetDefaultAddress() string {
	if address := os.Getenv(DefaultAddressEnv); address != "" {
		return address
	}
	return DefaultAddress
}

// returns the default password set by ether the config or an environment variable
func GetDefaultPassword() string {
	if password := os.Getenv(DefaultPasswordEnv); password != "" {
		return password
	}
	return DefaulftPassword
}

// returns the default timeout set by ether the config or an environment variable
func GetDefaultTimeout() time.Duration {
	if rawTimeout := os.Getenv(DefaltTimoutEnv); rawTimeout != "" {
		timeout, err := time.ParseDuration(rawTimeout)
		if err == nil {
			return timeout
		}
	}
	return DefaultTimeout
}
