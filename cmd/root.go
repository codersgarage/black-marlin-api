package cmd

import (
	"fmt"
	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/config"
	"github.com/codersgarage/black-marlin-api/log"
	"github.com/codersgarage/black-marlin-api/mq"
	"github.com/codersgarage/black-marlin-api/worker"
	"os"

	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command of boot-kit service
	RootCmd = &cobra.Command{
		Use:   "quest",
		Short: "quest is a grpc/http service",
		Long:  `An gRPC/HTTP JSON API backend service of boot-kit`,
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(migrationCmd)
}

// Execute executes the root command
func Execute() {
	if err := config.LoadConfig(); err != nil {
		fmt.Println("Failed to read config : ", err)
		os.Exit(1)
	}
	log.SetupLog()

	if err := app.ConnectDB(); err != nil {
		log.Log().Fatalf("Failed to connect to database : %v\n", err)
		os.Exit(-1)
	}
	if err := mq.ConnectMQ(); err != nil {
		log.Log().Fatalf("Failed to connect to queue server: %v\n", err)
	}
	worker.RunWorker()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
