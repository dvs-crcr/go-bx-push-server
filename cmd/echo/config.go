package echo

import (
    "flag"
    "os"
)

var (
    DefaultLogLevel = "info"
    DefaultAddress  = "localhost:8080"
)

const (
    TextAddressUsage  = "server endpoint address"
    TextLogLevelUsage = "logging level"
)

type Config struct {
    LogLevel string
    Address  string
}

// NewDefaultConfig returns default Config.
func NewDefaultConfig() *Config {
    return &Config{
        LogLevel: DefaultLogLevel,
        Address:  DefaultAddress,
    }
}

// parseVariables uses to parse cli flags and environment variables.
func (c *Config) parseVariables() {
    flag.StringVar(
        &c.Address, "a", c.Address, TextAddressUsage,
    )
    flag.StringVar(
        &c.LogLevel, "l", c.LogLevel, TextLogLevelUsage,
    )

    flag.Parse()

    if addrEnv := os.Getenv("ADDRESS"); addrEnv != "" {
        c.Address = addrEnv
    }

    if logLevelEnv := os.Getenv("LOG_LEVEL"); logLevelEnv != "" {
        c.LogLevel = logLevelEnv
    }
}
