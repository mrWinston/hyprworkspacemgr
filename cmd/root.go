/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var verbosity int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hyprworkspacemgr",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hyprworkspacemgr.yaml)")
  
  rootCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
    logrus.SetLevel(logrus.Level(verbosity))
    logrus.SetFormatter(&logrus.TextFormatter{
      ForceColors:               true,
      DisableColors:             false,
      ForceQuote:                false,
      DisableQuote:              true,
      EnvironmentOverrideColors: false,
      DisableTimestamp:          false,
      FullTimestamp:             false,
      QuoteEmptyFields:          true,
    })
    switch verbosity {
    case 0:
      logrus.SetLevel(logrus.ErrorLevel)
    case 1:
      logrus.SetLevel(logrus.InfoLevel)
    case 2:
      logrus.SetLevel(logrus.DebugLevel)
    default:
      return fmt.Errorf("verbosity must be one of: [0, 1, 2]. Was: %d", verbosity)
    }
    return nil
  }
  rootCmd.PersistentFlags().IntVarP(&verbosity, "verbosity", "v", 1, "verbosity level: 0, 1, 2")

}
