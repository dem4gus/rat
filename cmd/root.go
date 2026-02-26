package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "rat",
		Long: `Repository Audit Tool scans and reports on the configuration
of remotely hosted code repositories.

It brings the 'verify' to 'trust but verify'.`,
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_DIR/rat/config.toml)")

	cmd.AddCommand(AuditCommand())
	return cmd
}

func Execute() {
	rootCmd := RootCmd()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find config directory.
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(fmt.Sprintf("%s/rat", configDir))
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
