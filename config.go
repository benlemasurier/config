package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	// environment variable prefix
	prefix string
}

// New creates a new Config, working within the systems environment
// variables under the given prefix.
func New(prefix string) Config {
	return Config{
		prefix: prefix,
	}
}

// Get retrieves the request configuration value from the current environment.
func (c Config) Get(key string) string {
	return os.Getenv(fmt.Sprintf("%s_%s", c.prefix, key))
}

// Set saves a new configuration value within the current environment.
func (c Config) Set(key, value string) error {
	return os.Setenv(fmt.Sprintf("%s_%s", c.prefix, key), value)
}

// Unset unsets the underlying environment variable.
func (c Config) Unset(key string) error {
	return os.Unsetenv(fmt.Sprintf("%s_%s", c.prefix, key))
}

// Require checks for the existence of each provided key.
// Because the underlying configuration storage mechanism is backed
// by environment variables, empty values are considered to be missing.
func (c Config) Require(keys ...string) error {
	var missing = make([]string, 0)
	for _, k := range keys {
		fullKey := fmt.Sprintf("%s_%s", c.prefix, k)
		if os.Getenv(fullKey) == "" {
			missing = append(missing, fullKey)
		}
	}

	if len(missing) != 0 {
		message := fmt.Sprintf("missing %s environment variable(s)", strings.Join(missing, ", "))
		return errors.New(message)
	}

	return nil
}
