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
)

// reformatOnlyCmd represents the reformatOnly command
var reformatOnlyCmd = &cobra.Command{
	Use:   "only [directory / path to .sec file]",
	Short: "Only applies the reformat command to one file in the specified directory/path",
	Long: `Only applies the reformat command to one *asm.sec file in the specified directory or path.

The directory is searched recursively until all *asm.sec files are found.
The specified directory must not contain more than one *asm.sec file recursively. 
Calling "reformat only" on a directory that has more than one *asm.sec file will fail.
Thus in order for this condition to be met it might be required that you specify the path all the way to the single file you want to reformat. 

Examples:
  ratte reformat only assignment1
  ratte reformat only assignment1/some_file.sec`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		root := args[0]
		secFiles, err := getAllSecFiles(root)
		if err != nil {
			return err
		}
		if len(secFiles) != 1 {
			return fmt.Errorf("expected %s to lead to exactly 1 file of type *asm.sec but found: %d", root, len(secFiles))
		}
		err = renameToAsm(secFiles[0])
		if err != nil {
			return err
		}
		fmt.Println(format.RenameSuccess(dirutil.RelativePath(secFiles[0])))
		return nil
	},
}

func init() {
	reformatCmd.AddCommand(reformatOnlyCmd)
}
