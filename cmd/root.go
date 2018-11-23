package cmd

import (
	"fmt"
	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/config"
	"github.com/codersgarage/black-marlin-api/log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command of black-marlin-api service
	RootCmd = &cobra.Command{
		Use:   "black-marlin-api",
		Short: "black-marlin-api is a grpc/http service",
		Long:  `An gRPC/HTTP JSON API backend service of black-marlin-api`,
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

	if err := RootCmd.Execute(); err != nil {
		log.Log().Fatalf("Failed to execute root command : %v\n", err)
		os.Exit(1)
	}
}
