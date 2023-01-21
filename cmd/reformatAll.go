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
	"github.com/Moonpepperoni/ratte/dirutil"
	"github.com/Moonpepperoni/ratte/format"
	"github.com/spf13/cobra"
	"os"
)

// reformatAllCmd represents the all subcommand to of the reformat command
var reformatAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Searches the current directory for *asm.sec files recursively and applies the reformatting",
	Long: `Searches the current directory for *asm.sec files recursively and changes the file extension of all *asm.sec files to *.asm

Examples:
  ratte reformat all`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		root, err := os.Getwd()
		if err != nil {
			return err
		}
		secFiles, err := getAllSecFiles(root)
		if err != nil {
			return err
		}
		if len(secFiles) == 0 {
			return fmt.Errorf("expected %s to lead to at least one file with extension *asm.sec but found none", root)
		}
		for _, secFile := range secFiles {
			err := renameToAsm(secFile)
			if err != nil {
				fmt.Printf("Renaming file %s failed:\n%s\n", dirutil.RelativePath(secFile), err.Error())
			} else {
				fmt.Println(format.RenameSuccess(dirutil.RelativePath(secFile)))
			}
		}
		return nil
	},
}

func init() {
	reformatCmd.AddCommand(reformatAllCmd)
}
