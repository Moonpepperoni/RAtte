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
	Use:   "all",
	Short: "Searches the current directory for .sec files recursively and applies the reformatting",
	Long: `Searches the current directory for .sec files recursively and changes the file extension of all *asm.sec files to *.asm

Examples:
  ratte reformat all`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("called all subcommand of reformat")
	},
}

func init() {
	reformatCmd.AddCommand(reformatAllCmd)
}
