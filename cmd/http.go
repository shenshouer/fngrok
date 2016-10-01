package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// httpCmd start http proxy
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start http proxy ",
	Run: func(cmd *cobra.Command, args []string) {
		setLog(cmd)
		fmt.Println(currentFlag, "Start Http Proxy")
		switch currentFlag {
		case FlagServer:
			break
		case FlagClient:
			break
		}
	},
}
