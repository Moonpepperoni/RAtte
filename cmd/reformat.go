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
	"github.com/Moonpepperoni/ratte/dirutil"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// reformatCmd represents the reformat root command it is not directly callable
var reformatCmd = &cobra.Command{
	Use:   "reformat",
	Short: "Converts *asm.sec files to *.asm files",
	Long: `Converts *asm.sec files to *.asm files. 
all/only subcommands changes how many and which files are affected by this command.

Examples: 
  ratte reformat all
  ratte reformat only student_001
  ratte reformat only student_001/assignment1`,
}

func init() {
	rootCmd.AddCommand(reformatCmd)
}

func filterSecFiles(paths []string) []string {
	return dirutil.FilterPaths(paths, func(path string) bool {
		return strings.HasSuffix(path, "asm.sec")
	})
}

func getAllSecFiles(root string) ([]string, error) {
	allFiles, err := dirutil.GetAllFilesRecur(root)
	if err != nil {
		return nil, err
	}
	return filterSecFiles(allFiles), nil
}

func renameToAsm(path string) error {
	return os.Rename(path, strings.Replace(path, "asm.sec", ".asm", 1))
}
