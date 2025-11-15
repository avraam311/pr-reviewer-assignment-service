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
	ErrLoadEnvFile    = errors.New("failed to load env file")
	ErrLoadConfigFile = errors.New("failed to load config file")
)

type Config struct {
	v *viper.Viper
}

func New() *Config {
	v := viper.New()
	return &Config{v: v}
}

func (c *Config) LoadEnvFiles(paths ...string) error {
	for _, path := range paths {
		if err := godotenv.Load(path); err != nil {
			return fmt.Errorf("%w %s: %w", ErrLoadEnvFile, path, err)
		}
	}
	return nil
}

func (c *Config) LoadConfigFiles(paths ...string) error {
	for _, cfgPath := range paths {
		c.v.SetConfigFile(cfgPath)
		if err := c.v.MergeInConfig(); err != nil {
			return fmt.Errorf("%w %s: %w", ErrLoadConfigFile, cfgPath, err)
		}
	}
	return nil
}

func (c *Config) EnableEnv(envPrefix string) {
	c.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if envPrefix != "" {
		c.v.SetEnvPrefix(envPrefix)
	}
	c.v.AutomaticEnv()
}

func (c *Config) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *Config) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *Config) GetInt32(key string) int32 {
	return c.v.GetInt32(key)
}

func (c *Config) GetInt64(key string) int64 {
	return c.v.GetInt64(key)
}

func (c *Config) GetBool(key string) bool {
	return c.v.GetBool(key)
}

func (c *Config) GetFloat64(key string) float64 {
	return c.v.GetFloat64(key)
}

func (c *Config) GetTime(key string) time.Time {
	return c.v.GetTime(key)
}

func (c *Config) GetDuration(key string) time.Duration {
	return c.v.GetDuration(key)
}

func (c *Config) GetStringSlice(key string) []string {
	return c.v.GetStringSlice(key)
}

func (c *Config) GetIntSlice(key string) []int {
	return c.v.GetIntSlice(key)
}
