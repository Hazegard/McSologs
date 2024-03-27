package config

import (
	"fmt"
	"os"
)

type Config struct {
	LogFile    string
	DiscordUrl string
}

func NewConfig(logFile string, configFile string) (*Config, error) {
	notifyUrl, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading notify URL file %s: %s", configFile, err)
	}
	return &Config{
		LogFile:    logFile,
		DiscordUrl: string(notifyUrl),
	}, nil
}
