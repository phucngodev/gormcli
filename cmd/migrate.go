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

// Migrate
var version uint

// dbCmd represents the db command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	Long:  "Run database migration",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := cli.InitMigrate()
		if err != nil {
			fmt.Printf("[Migrate] %s\n", err)
			os.Exit(1)
		}
		defer m.Close()
		if version == 0 {
			err = m.Up()
		} else {
			err = m.Migrate(version)
		}

		if err != nil {
			fmt.Printf("[Migrate] %s\n", err)
			os.Exit(1)
		}
		v, _, err := m.Version()
		if err != nil {
			fmt.Printf("[Migrate] %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Migrated to version %d successfully.\n", v)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().UintVarP(&version, "version", "v", 0, "migrate database up/down to version")
}
