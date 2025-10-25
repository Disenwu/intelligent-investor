package server

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultHomeDir    = ".intelligent-investor"
	defaultConfigName = "api-server.yaml"
)

func onInitialize() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		for _, dir := range searchDirs() {
			viper.AddConfigPath(dir)
		}
		viper.SetConfigType("yaml")
		viper.SetConfigFile(defaultConfigName)
	}
	setupEnvironmentVariables()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read viper configuration file, err: %v", err)
	}
	log.Printf("Using config file: %s", viper.ConfigFileUsed())
}

func setupEnvironmentVariables() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("apiserver")
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

func searchDirs() []string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}

func filePath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defaultHomeDir, defaultConfigName)
}
