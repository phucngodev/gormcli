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

var all bool

// downCmd represents the down command
var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback to previous version",
	Long:  "Rollback to previous vesrion",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := cli.InitMigrate()
		if err != nil {
			fmt.Printf("[Rollback] %s\n", err)
			os.Exit(1)
		}
		defer m.Close()

		v, _, err := m.Version()
		if err != nil {
			fmt.Printf("[Rollback] %s\n", err)
			os.Exit(1)
		}

		if v == 1 {
			all = true
		}

		// Rollback all migration
		if all {
			err = m.Down()
			if err != nil {
				fmt.Printf("[Rollback] %s\n", err)
				os.Exit(1)
			}

			fmt.Println("Rollbak all migrations.")
			os.Exit(0)
		}

		// Rollback to previous version
		err = m.Migrate(v - 1)
		if err != nil {
			fmt.Printf("[Rollback] %s\n", err)
			os.Exit(1)
		}

		v, _, err = m.Version()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Rollback to version %d successfully.\n", v)
	},
}

func init() {
	migrateCmd.AddCommand(rollbackCmd)
	rollbackCmd.Flags().BoolVarP(&all, "all", "a", false, "rollback all migrations")
}
