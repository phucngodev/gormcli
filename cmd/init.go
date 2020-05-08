/*
Copyright © 2020 Phuc

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
	"os"

	"github.com/spf13/cobra"
)

var dbType string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup migration feature for current project",
	Long:  "Setup migration feature for current project",
	Run: func(cmd *cobra.Command, args []string) {
		if dbType == "" {
			dbType = "mysql"
		}
		err := createConfigFile(dbType)
		if err != nil {
			fmt.Printf("[Init] failed to create app config file: %+v\n", err)
			os.Exit(1)
		}

		err = createMigrationDirectory()
		if err != nil {
			fmt.Printf("[Init] failed to create migration directory: %+v\n", err)
			os.Exit(1)
		}

		fmt.Println("Database migration init successfully, run gorm --help form available commands.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&dbType, "type", "t", "mysql", "initialize project using gorm in current directory")
}
