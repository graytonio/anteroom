package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/graytonio/chamber/lib"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:     "ante",
		Short:   "Antechamber is a CLI tool for generating boilerplate code",
		Version: "0.0.1",
	}
)

type Config struct {
	ConfigFile     *ini.File
	ConfigFilePath string
	Verbose        bool
}

var config Config

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	home, _ := os.UserHomeDir()

	rootCmd.PersistentFlags().StringVarP(&config.ConfigFilePath, "config", "c", path.Join(home, ".antecfg"), "Path to Config File")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "Verbose")
}

func initConfig() {
	// Enable Verbose
	if config.Verbose {
		lib.EnableDebugLogger()
	}

	// Load Config File
	cfg, err := ini.Load(config.ConfigFilePath)
	if err != nil {
		lib.LogError(fmt.Sprintf("Could not load config file %s\n%s", config.ConfigFilePath, err.Error()))
		os.Exit(1)
	}

	// Save config to config global
	config.ConfigFile = cfg
}
