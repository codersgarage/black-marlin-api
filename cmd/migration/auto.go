package migration

import (
	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/log"
	"github.com/codersgarage/black-marlin-api/models"
	"github.com/codersgarage/black-marlin-api/utils"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var MigAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "auto alter database tables",
	Run:   auto,
}

func auto(cmd *cobra.Command, args []string) {
	success := true

	admin := &models.Admin{}
	if err := app.DB().AutoMigrate(admin).Error; err != nil {
		log.Log().Infof("Failed to auto migrate table [%s] with error {%v}", admin.TableName(), err)
		success = false
	}
	adminSession := &models.AdminSession{}
	if err := app.DB().AutoMigrate(adminSession).Error; err != nil {
		log.Log().Infof("Failed to auto migrate table [%s] with error {%v}", adminSession.TableName(), err)
		success = false
	}
	a := &models.App{}
	if err := app.DB().AutoMigrate(a).Error; err != nil {
		log.Log().Infof("Failed to auto migrate table [%s] with error {%v}", a.TableName(), err)
		success = false
	}
	session := &models.Session{}
	if err := app.DB().AutoMigrate(session).Error; err != nil {
		log.Log().Infof("Failed to auto migrate table [%s] with error {%v}", session.TableName(), err)
		success = false
	}
	sub := &models.Subscription{}
	if err := app.DB().AutoMigrate(sub).Error; err != nil {
		log.Log().Infof("Failed to auto migrate table [%s] with error {%v}", sub.TableName(), err)
		success = false
	}
	user := &models.User{}
	if err := app.DB().AutoMigrate(user).Error; err != nil {
		log.Log().Infof("Failed to auto migrate table [%s] with error {%v}", user.TableName(), err)
		success = false
	}
	settings := &models.Settings{}
	if err := app.DB().AutoMigrate(settings).Error; err != nil {
		log.Log().Infof("Failed to auto migrate table [%s] with error {%v}", settings.TableName(), err)
		success = false
	}

	if success {
		settings := models.Settings{}
		settings.ID = 1
		settings.ProjectName = "Black Marlin"
		settings.ProjectDescription = "Black Marlin is a highly scalable MQTT based push notification server."
		settings.SystemKey = os.Getenv("BLACK_MARLIN_SYSTEM_KEY")
		settings.ServiceKey = os.Getenv("BLACK_MARLIN_SERVICE_KEY")

		if err := app.DB().Create(&settings).Error; err != nil {
			log.Log().Infof("Failed to create default settings with error {%v}", err)
		}

		as := models.Admin{}
		as.ID = 1
		as.UUID = utils.UUID()
		as.Name = "Black Marlin Admin"
		as.Email = "black.marlin.admin@example"
		as.Password = utils.CreateHashedPassword("@1234567890")
		as.CreatedAt = time.Now()
		as.UpdatedAt = time.Now()
		as.IsActive = true

		if err := app.DB().Create(&as).Error; err != nil {
			log.Log().Infof("Failed to create default user with error {%v}", err)
		}
	}
}
