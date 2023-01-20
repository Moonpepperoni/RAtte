/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// reformatAllCmd represents the all subcommand to of the reformat command
var reformatAllCmd = &cobra.Command{
	Use:   "all [directory]",
	Short: "Applies the reformat command to all files in the optionally specified directory.",
	Long: `Applies the reformat command to all .sec files in the specified directory.
directory is an optional argument. Calling reformat all without specifying directory explicitly will run reformat in the currently selected directory: e.g. ./  

Examples:
  ratte reformat all
  ratte reformat all assignment1`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("called all subcommand of reformat")
	},
}

func init() {
	reformatCmd.AddCommand(reformatAllCmd)
}
