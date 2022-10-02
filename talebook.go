package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/bookstairs/talebook/config"
	"github.com/bookstairs/talebook/handler"
)

const (
	// The name of our config file, without the file extension because viper supports many config file languages.
	defaultConfigFile = "talebook"

	// The environment variable prefix of all environment variables bound to our command line flags.
	envPrefix = "TALE"
)

// The final configuration would be generated by the following order:
// flags > environment variables > configuration files > flag defaults
// See https://carolynvanslyck.com/blog/2020/08/sting-of-the-viper/ for how we achieve this.
func main() {
	// Paring the args and start applications.
	cmd := NewRootCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// NewRootCommand will build the cobra command that handles our command line tool.
func NewRootCommand() *cobra.Command {
	// Store the result of binding cobra flags and viper config. In a
	// real application, these would be data structures, most likely
	// custom structs per command. This is simplified for the demo app and is
	// not recommended that you use one-off variable. The point is that we
	// aren't retrieving the values directly from viper or flags, we read the values
	// from standard Go data structures.
	port := 0
	workingPath := ""
	libraryPath := ""
	encryptKey := ""
	limit := 0
	calibreDB := ""
	convert := ""
	fileCache := 0
	debug := false

	// Create a default configuration with config value.
	dc := config.DefaultSeverConfig()

	// The cobra command for executing server.
	rootCmd := &cobra.Command{
		Use: "talebook",
		Short: "Talebook (in Golang)\n\n" +
			"This a fork of github.com/talebook/talebook. Serve as your personal library.\n\n" +
			config.TalebookVersion().String() +
			"\n",
		Version: config.TalebookVersion().String(),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initializeConfig(cmd, workingPath)
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Make sure we use the proper library path.
			if libraryPath == dc.LibraryPath && workingPath != dc.WorkingPath {
				libraryPath = config.DefaultLibraryPath(workingPath)
			}

			c := &config.ServerConfig{
				Port:        port,
				WorkingPath: workingPath,
				LibraryPath: libraryPath,
				EncryptKey:  encryptKey,
				Limit:       limit,
				CalibreDB:   calibreDB,
				Convert:     convert,
				FileCache:   fileCache,
				Debug:       debug,
				Frontend:    frontend,
			}

			// Create working directories and perform other checks.
			initRuntime(c)

			// Bootstrap the talebook server.
			handler.StartServer(c)
		},
	}

	// Register the talebook configuration flags. These would override the configurations in file.
	rootCmd.Flags().IntVarP(&port, "port", "p", dc.Port, "The http port for talebook.")
	rootCmd.Flags().StringVarP(&workingPath, "working-path", "w", dc.WorkingPath, "The working directory for talebook.")
	rootCmd.Flags().StringVarP(&libraryPath, "library-path", "l", dc.LibraryPath, "The calibre library directory.")
	rootCmd.Flags().StringVarP(&encryptKey, "encrypt-key", "e", dc.EncryptKey, "The key for encrypting the cookie content.")
	rootCmd.Flags().IntVarP(&limit, "ratelimit", "r", dc.Limit, "Add limit for the requests per second.")
	rootCmd.Flags().StringVarP(&calibreDB, "calibredb", "", dc.CalibreDB, "The full path for calibredb(.exe).")
	rootCmd.Flags().StringVarP(&convert, "convert", "", dc.Convert, "The full path for ebook-convert(.exe).")
	rootCmd.Flags().IntVarP(&fileCache, "file-cache", "", dc.FileCache, "The file cache (MB) in memory.")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", dc.Debug, "This shouldn't be enable in production.")

	// Add version flag.
	rootCmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)
	rootCmd.InitDefaultVersionFlag()

	return rootCmd
}

func initializeConfig(cmd *cobra.Command, dir string) error {
	v := viper.New()

	// Set the base name of the config file, without the file extension.
	v.SetConfigName(defaultConfigFile)

	// Set as many paths as you like where viper should look for the
	// config file. We are only looking in the current working directory.
	v.AddConfigPath(".")
	if dir != "" {
		v.AddConfigPath(dir)
	}

	// Attempt to read the config file, gracefully ignoring errors
	// caused by a config file not being found. Return an error
	// if we cannot parse the config file.
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable STING_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix(envPrefix)

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent keys with underscores.
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			_ = v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and the viper has a value.
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
