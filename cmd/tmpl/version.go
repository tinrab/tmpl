package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionShort bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		if versionShort {
			fmt.Println(version)
		} else {
			fmt.Printf("tmpl v%s\n", version)
		}
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&versionShort, "short", "s", false, "short form version")
}
