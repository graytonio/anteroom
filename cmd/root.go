package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path"

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
var debug_logger *log.Logger
var error_logger *log.Logger
var status_logger *log.Logger

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	error_logger = log.New(os.Stderr, "", 0)
	debug_logger = log.New(ioutil.Discard, "", log.LstdFlags)
	status_logger = log.New(os.Stdout, "", 0)

	home, _ := os.UserHomeDir()

	rootCmd.PersistentFlags().StringVarP(&config.ConfigFilePath, "config", "c", path.Join(home, ".antecfg"), "Path to Config File")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "Verbose")
}

func initConfig() {
	// Enable Verbose
	if config.Verbose {
		debug_logger.SetOutput(os.Stdout)
	}

	// Load Config File
	cfg, err := ini.Load(config.ConfigFilePath)
	if err != nil {
		error_logger.Fatalf("Could not load config file %s\n%s", config.ConfigFilePath, err.Error())
	}

	// Save config to config global
	config.ConfigFile = cfg
}
