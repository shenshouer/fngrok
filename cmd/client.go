package cmd

import (
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootClientCmd represents the base command when called without any subcommands
var RootClientCmd = &cobra.Command{
	Use:   "fngrok",
	Short: "client for fngrok service",
	Run: func(cmd *cobra.Command, args []string) {
		setLog(cmd)
		log.Infof("This is %s binary for fngrok", currentFlag)
		for {
			// for test
			log.Infoln(time.Now())
			time.Sleep(1 * time.Second)
		}
	},
}

// ExecuteClient adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func ExecuteClient() {
	currentFlag = FlagClient // mark this is client

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

	RootClientCmd.PersistentFlags().String("server_addr", "", "Server address for fngrok service")
	RootClientCmd.PersistentFlags().String("server_port", "", "Server port for fngrok service")
	RootClientCmd.PersistentFlags().String("log_level", "info", "log level")

	viper.BindPFlag("server_addr", RootClientCmd.PersistentFlags().Lookup("server_addr"))
	viper.BindPFlag("server_port", RootClientCmd.PersistentFlags().Lookup("server_port"))
	viper.BindPFlag("log_level", RootClientCmd.PersistentFlags().Lookup("log_level"))

	RootClientCmd.AddCommand(httpCmd)
	RootClientCmd.AddCommand(versionCmd)
}
