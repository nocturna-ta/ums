package cmd

import (
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/ums/cmd/server"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Votechain : User Management Service",
		Short: "Votechain : User Management Service",
	}
)

func Execute() {
	log.SetFormatter("json")
	rootCmd.AddCommand(server.ServeHttpCmd())
	rootCmd.AddCommand(server.ServeGrpcCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Error: ", err.Error())
		os.Exit(-1)
	}
}
