package server

import (
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/handler/grpc"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var (
	serveGrpcCmd = &cobra.Command{
		Use:   "serve-grpc",
		Short: "User Management System Service gRPC",
		Long:  "User Management System Service through gRPC",
		RunE:  runGrpc,
	}
)

func ServeGrpcCmd() *cobra.Command {
	serveGrpcCmd.Flags().StringP("config", "c", "", "Config Path, both relative or absolute. i.e: /usr/local/bin/config/files")
	return serveGrpcCmd
}

func runGrpc(cmd *cobra.Command, args []string) error {
	configLocation, _ := cmd.Flags().GetString("config")
	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, configLocation)

	database := sql.New(sql.DBConfig{
		SlaveDSN:        cfg.Database.SlaveDSN,
		MasterDSN:       cfg.Database.MasterDSN,
		RetryInterval:   cfg.Database.RetryInterval,
		MaxIdleConn:     cfg.Database.MaxIdleConn,
		MaxConn:         cfg.Database.MaxConn,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	}, sql.DriverPostgres)

	appContainer := newContainer(&options{
		Cfg: cfg,
		DB:  database,
	})

	server := grpc.New(&grpc.Options{
		Cfg:    appContainer.Cfg,
		AuthUc: appContainer.AuthUc,
	})

	go server.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Info("Exiting gracefully....")
	}

	return nil
}
