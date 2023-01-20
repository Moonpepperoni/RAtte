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
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ratte",
	Short: "ratte is a cli to build and test simple I/O bound assembly code",
	Long: `ratte is a cli to build and test simple I/O bound assembly code via predefined interactions
Example calls: 
ratte reformat all
ratte build all
ratte test only student_001 test3.json
ratte test only assignment1 test1.json

All commands need to be run with either subcommand 'all' or subcommand 'only'

Prerequisites: 
Requires commands gcc and nasm to be callable under precisely those names.
Requires SHELL environment variable to be set correctly`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
