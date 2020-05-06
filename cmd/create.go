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

	"github.com/phucngodev/gormcli/internal/cli"
	"github.com/spf13/cobra"
)

var FileName string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create migration file",
	Long:  "Create migration file",
	Run: func(cmd *cobra.Command, args []string) {
		if FileName == "" {
			fmt.Println("[Create] migration name is required")
			os.Exit(1)
		}
		err := cli.CreateCmd(FileName)
		if err != nil {
			fmt.Printf("[Create] %s\n", err)
		}
	},
}

func init() {
	migrateCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&FileName, "name", "n", "", "gorm db create --name create_user_table")
}
