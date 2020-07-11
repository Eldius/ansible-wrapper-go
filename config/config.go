package config

import (
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	pyenvFolder   = "pyenv"
	scriptsFolder = "scripts"
)

/*
AppConfig represents the
app configuration
*/
type AppConfig struct {
	Verbose   bool
	Workspace string
}

/*
WorkspaceFolder returns the workspace folder
*/
func (c *AppConfig) WorkspaceFolder() string {
	if wsDir, err := homedir.Expand(c.Workspace); err != nil {
		panic(err.Error())
	} else {
		return wsDir
	}
}

/*
GetPyenvFolder returns the pyenv path
*/
func (c *AppConfig) GetPyenvFolder() string {
	return filepath.Join(c.Workspace, pyenvFolder)
}

/*
GetPyenvBinFolder returns pyenv bin folder
*/
func (c *AppConfig) GetPyenvBinFolder() string {
	return filepath.Join(c.Workspace, pyenvFolder, "bin")
}

/*
GetScriptsFolder returns the scripts folder
*/
func (c *AppConfig) GetScriptsFolder() string {
	return filepath.Join(c.Workspace, scriptsFolder)
}

/*
SaveConfiguration persists the
current configuration
*/
func SaveConfiguration() {
	viper.SafeWriteConfig()
}

/*
GetAppConfig returns the app config
*/
func GetAppConfig() *AppConfig {
	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panic(err.Error())
	}

	if cfg.Verbose {
		log.Println(cfg)
	}
	return &cfg
}
