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

// playerMap holds the matches between a Minecraft player name and the corresponding Discord ID
type playerMap map[string]string

// Get returns the Discord ID of the player if found, otherwise returns the Minecraft player name
// The discord ID is formatted as required by discord to ping the player
func (p *playerMap) Get(player string) string {
	discordPlayer, ok := (*p)[player]
	if ok {
		return fmt.Sprintf("<@%s>", discordPlayer)
	}
	return player
}

var c *Config

// NewConfig returns the struct holding the global configuration
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

// configFile holds the structs used to parse the json configuration file
type configFile struct {
	DiscordUrl string            `json:"discordUrl,omitempty"`
	Players    map[string]string `json:"players,omitempty"`
}

// parseConfigFile parses the configuration file and return the corresponding struct
// returns an error if the file parsing fails
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
