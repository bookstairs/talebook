package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// This file only used for storing the command line arguments.
const defaultConfigFile = "talebook.yml"

var (
	// The flag variables.
	port       = 0
	workingDir = ""
	libraryDir = ""
	encryptKey = ""
	limit      = 0
	calibreDB  = ""
	convert    = ""
	debug      = false

	// The full config file path.
	configFile = ""

	// The cobra command for executing server.
	rootCmd = &cobra.Command{
		Use:   "talebook",
		Short: "This a fork of github.com/talebook/talebook. Serve as your personal library.",
		Run: func(cmd *cobra.Command, args []string) {
			c := DefaultSeverConfig()

			// Load config from config file.
			if _, err := os.Stat(configFile); err == nil {
				viper.SetConfigFile(configFile)
				viper.AutomaticEnv()
				if err = viper.ReadInConfig(); err != nil {
					log.Fatal(err)
				}
				log.Println("Using config file:", viper.ConfigFileUsed())
				if err = viper.Unmarshal(c); err != nil {
					log.Fatal(err)
				}

			}

			// Override configuration from flags.
			dc := DefaultSeverConfig()
			if port != dc.Port {
				c.Port = port
			}
			if workingDir != dc.WorkingPath {
				c.WorkingPath = workingDir
			}
			if libraryDir != dc.LibraryPath {
				c.LibraryPath = libraryDir
			}
			if encryptKey != dc.EncryptKey {
				c.EncryptKey = encryptKey
			}
			if limit != dc.Limit {
				c.Limit = limit
			}
			if calibreDB != dc.CalibreDB {
				c.CalibreDB = calibreDB
			}
			if convert != dc.Convert {
				c.Convert = convert
			}
			if debug != dc.Debug {
				c.Debug = debug
			}

			// Bootstrap the talebook server.
			StartServer(c)
		},
	}
)

func init() {
	// Create a default configuration with config value.
	c := DefaultSeverConfig()
	configFile = filepath.Join(c.WorkingPath, defaultConfigFile)

	// Register the talebook configuration file.
	rootCmd.Flags().StringVarP(&configFile, "config", "c", configFile, "The configuration file for talebook.")

	// Register the talebook configuration flags. These would override the configurations in file.
	rootCmd.Flags().IntVarP(&port, "port", "p", c.Port, "The http port for talebook.")
	rootCmd.Flags().StringVarP(&workingDir, "working-dir", "w", c.WorkingPath, "The working directory for talebook.")
	rootCmd.Flags().StringVarP(&libraryDir, "library-dir", "l", c.LibraryPath, "The calibre library directory.")
	rootCmd.Flags().StringVarP(&encryptKey, "encrypt-key", "e", c.EncryptKey, "The key for encrypting the cookie content.")
	rootCmd.Flags().IntVarP(&limit, "ratelimit", "r", c.Limit, "Add limit for the requests per second.")
	rootCmd.Flags().StringVarP(&calibreDB, "calibredb", "", c.CalibreDB, "The full path for calibredb(.exe).")
	rootCmd.Flags().StringVarP(&convert, "convert", "", c.Convert, "The full path for ebook-convert(.exe).")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", c.Debug, "Enable some functions for debug purpose. This shouldn't be enable in production.")
}

func main() {
	// Paring the args and start applications.
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
