package migration

import (
	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/log"
	"github.com/codersgarage/black-marlin-api/models"
	"github.com/spf13/cobra"
)

var MigDropCmd = &cobra.Command{
	Use:   "drop",
	Short: "drop drops database tables",
	Run:   drop,
}

func drop(cmd *cobra.Command, args []string) {
	monkey := &models.Monkey{}
	if err := app.DB().DropTable(monkey).Error; err != nil {
		log.Log().Fatalf("Failed to drop table [%s] with error {%v}", monkey.TableName(), err)
	}
}
