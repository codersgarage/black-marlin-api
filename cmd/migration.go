package cmd

import (
	"github.com/codersgarage/black-marlin-api/cmd/migration"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration migrates database schemas",
}

func init() {
	migrationCmd.AddCommand(migration.MigCreateCmd)
	migrationCmd.AddCommand(migration.MigDropCmd)
	migrationCmd.AddCommand(migration.MigAutoCmd)
}
