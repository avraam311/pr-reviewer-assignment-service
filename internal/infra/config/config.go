// Package config provides configuration management using Viper.
package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	// ErrLoadEnvFile is returned when loading a .env file fails.
	ErrLoadEnvFile = errors.New("failed to load env file")
	// ErrLoadConfigFile is returned when loading a config file fails.
	ErrLoadConfigFile = errors.New("failed to load config file")
)

// Config wraps an instance of Viper configuration.
type Config struct {
	v *viper.Viper
}

// New creates a new Config instance.
func New() *Config {
	v := viper.New()
	return &Config{v: v}
}

// LoadEnvFiles loads one or more .env files into os.Environ().
func (c *Config) LoadEnvFiles(paths ...string) error {
	for _, path := range paths {
		if err := godotenv.Load(path); err != nil {
			return fmt.Errorf("%w %s: %w", ErrLoadEnvFile, path, err)
		}
	}
	return nil
}

// LoadConfigFiles loads and merges multiple config files.
func (c *Config) LoadConfigFiles(paths ...string) error {
	for _, cfgPath := range paths {
		c.v.SetConfigFile(cfgPath)
		if err := c.v.MergeInConfig(); err != nil {
			return fmt.Errorf("%w %s: %w", ErrLoadConfigFile, cfgPath, err)
		}
	}
	return nil
}

// EnableEnv enables automatic loading of environment variables.
// envPrefix (if set) is used as a prefix for all keys.
func (c *Config) EnableEnv(envPrefix string) {
	c.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if envPrefix != "" {
		c.v.SetEnvPrefix(envPrefix)
	}
	c.v.AutomaticEnv()
}

// GetString gets a string value from the config by key.
func (c *Config) GetString(key string) string {
	return c.v.GetString(key)
}

// GetInt gets an integer value from the config by key.
func (c *Config) GetInt(key string) int {
	return c.v.GetInt(key)
}

// GetInt32 gets an int32 value from the config by key.
func (c *Config) GetInt32(key string) int32 {
	return c.v.GetInt32(key)
}

// GetInt64 gets an int64 value from the config by key.
func (c *Config) GetInt64(key string) int64 {
	return c.v.GetInt64(key)
}

// GetBool gets a boolean value from the config by key.
func (c *Config) GetBool(key string) bool {
	return c.v.GetBool(key)
}

// GetFloat64 gets a float64 value from the config by key.
func (c *Config) GetFloat64(key string) float64 {
	return c.v.GetFloat64(key)
}

// GetTime gets a time.Time value from the config by key.
func (c *Config) GetTime(key string) time.Time {
	return c.v.GetTime(key)
}

// GetDuration gets a time.Duration value from the config by key.
func (c *Config) GetDuration(key string) time.Duration {
	return c.v.GetDuration(key)
}

// GetStringSlice gets a slice of strings from the config by key.
func (c *Config) GetStringSlice(key string) []string {
	return c.v.GetStringSlice(key)
}

// GetIntSlice gets a slice of integers from the config by key.
func (c *Config) GetIntSlice(key string) []int {
	return c.v.GetIntSlice(key)
}
