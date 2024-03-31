package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	LogFile    string
	DiscordUrl string
	Players    playerMap
	Debug      bool
}

type playerMap map[string]string

func (p *playerMap) Get(player string) string {
	discordPlayer, ok := (*p)[player]
	if ok {
		return discordPlayer
	}
	return player
}

var c *Config

func NewConfig(logFile string, configFileName string, debug bool) (*Config, error) {
	err, parsedConfig := parseConfigFile(configFileName)
	if err != nil {
		return nil, err
	}
	c = &Config{
		LogFile:    logFile,
		DiscordUrl: parsedConfig.DiscordUrl,
		Players:    parsedConfig.Players,
		Debug:      debug,
	}
	return c, nil
}

func GetConfig() *Config {
	return c
}

type configFile struct {
	DiscordUrl string            `json:"discordUrl,omitempty"`
	Players    map[string]string `json:"players,omitempty"`
}

func parseConfigFile(configFilename string) (error, *configFile) {
	var parsedConfig *configFile
	data, err := os.ReadFile(configFilename)
	if err != nil {
		return fmt.Errorf("error while reading config file %s: %s", configFilename, err), nil
	}

	err = json.Unmarshal(data, &parsedConfig)
	if err != nil {
		return fmt.Errorf("error while parsing config file %s as json: %s", configFilename, err), nil
	}
	return nil, parsedConfig
}
