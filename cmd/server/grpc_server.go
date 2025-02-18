package server

//
//import (
//	"github.com/nocturna-ta/golib/database/sql"
//	"github.com/nocturna-ta/ums/config"
//	"github.com/spf13/cobra"
//)
//
//var (
//	serveGrpc = &cobra.Command{
//		Use:   "serve-grpc",
//		Short: "User Management System Service gRPC",
//		Long:  "User Management System Service through gRPC",
//		RunE:  nil,
//	}
//)
//
//func ServeGrpcCmd() *cobra.Command {
//	serveGrpcCmd.
//}
//
//func runGrpc(cmd *cobra.Command, args []string) error  {
//	configLocation, _ := cmd.Flags().GetString("config")
//	cfg := &config.MainConfig{}
//	config.ReadConfig(cfg, configLocation)
//
//	database := sql.New(sql.DBConfig{
//		SlaveDSN:        cfg.Database.SlaveDSN,
//		MasterDSN:       cfg.Database.MasterDSN,
//		RetryInterval:   cfg.Database.RetryInterval,
//		MaxIdleConn:     cfg.Database.MaxIdleConn,
//		MaxConn:         cfg.Database.MaxConn,
//		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
//	}, sql.DriverPostgres)
//
//	appContainer := newContainer(&options{
//		Cfg:       cfg,
//		DB:        database,
//	})
//
//	server := grpc.New
//}
