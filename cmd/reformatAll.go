/*
Copyright © 2023 Jan Leonhardt

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
