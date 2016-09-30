package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootClientCmd represents the base command when called without any subcommands
var RootClientCmd = &cobra.Command{
	Use:   "fngrok",
	Short: "client for fngrok service",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// ExecuteClient adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func ExecuteClient() {
	if err := RootClientCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootClientCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fngrok.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootClientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
