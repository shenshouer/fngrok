package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootServerCmd represents the base command when called without any subcommands
var RootServerCmd = &cobra.Command{
	Use:   "fngrok",
	Short: "server for fngrok service",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This is %s binary for fngrok", currentFlag)
	},
}

// ExecuteServer adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func ExecuteServer() {
	currentFlag = FlagServer // mark this is server

	if err := RootServerCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootServerCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fngrok.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootServerCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")

	RootServerCmd.AddCommand(httpCmd)
	RootServerCmd.AddCommand(versionCmd)
}
