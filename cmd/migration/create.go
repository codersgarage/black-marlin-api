package migration

import (
	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/log"
	"github.com/codersgarage/black-marlin-api/models"
	"github.com/spf13/cobra"
)

var MigCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create creates database tables",
	Run:   create,
}

func create(cmd *cobra.Command, args []string) {
	monkey := &models.Monkey{}
	if err := app.DB().CreateTable(monkey).Error; err != nil {
		log.Log().Fatalf("Failed to create table [%s] with error {%v}", monkey.TableName(), err)
	}
}
