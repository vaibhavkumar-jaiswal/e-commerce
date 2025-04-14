// Package cmd provides command-line interface commands for the application
// This file contains the migrate command for running database migrations
// using GORM. It allows users to specify which models to migrate
// via command-line flags. The command is registered with the root command
// and can be executed from the terminal.
package cmd

import (
	"e-commerce/database/migrations"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var modelNames string

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs database migrations using GORM",
	Run: func(_ *cobra.Command, _ []string) {
		var dbModels []string
		if modelNames != "" {
			dbModels = strings.Split(modelNames, ",")
		}
		err := migrations.RunMigrations(dbModels...)
		if err != nil {
			fmt.Println("Error running migrations:", err)
			return
		}
		fmt.Println("Migrations completed successfully.")
	},
}

func init() {
	migrateCmd.Flags().StringVarP(
		&modelNames,
		"models",
		"m",
		"",
		"Comma-separated model names to migrate (e.g., User, Product)",
	)
	rootCmd.AddCommand(migrateCmd)
}
