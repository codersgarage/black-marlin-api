package migration

import (
	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/log"
	"github.com/codersgarage/black-marlin-api/models"
	"github.com/spf13/cobra"
)

var MigAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "auto alter database tables",
	Run:   auto,
}

func auto(cmd *cobra.Command, args []string) {
	monkey := &models.Monkey{}
	if err := app.DB().AutoMigrate(monkey).Error; err != nil {
		log.Log().Infof("Failed to auto migrate table [%s] with error {%v}", monkey.TableName(), err)
	}
}
