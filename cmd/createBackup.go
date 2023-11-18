/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/spf13/cobra"
)

// createBackupCmd represents the createBackup command
var createBackupCmd = &cobra.Command{
	Use:   "create-backup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		user := cmd.Flag("user").Value.String()
		password := cmd.Flag("password").Value.String()
		backupName := cmd.Flag("backup_name").Value.String()

		conn, err := connect(host, port, user, password)
		if err != nil {
			fmt.Println("Error connecting to database: ", err)
			return
		}

		ctx := context.Background()
		_, err = conn.Query(ctx, "BACKUP TABLE my_table TO Disk('backups', '"+backupName+"')")
		if err != nil {
			fmt.Println("Error creating backup: ", err)
			return
		}

		fmt.Println("Backup created successfully")


	},
}

func init() {
	clickhouseCmd.AddCommand(createBackupCmd)

	createBackupCmd.Flags().StringP("host", "H", "localhost", "Host of the clickhouse server")
	createBackupCmd.Flags().StringP("port", "p", "9000", "Port of the clickhouse server")
	createBackupCmd.Flags().StringP("user", "u", "", "User of the clickhouse server")
	createBackupCmd.Flags().StringP("password", "P", "", "Password of the clickhouse server")
	createBackupCmd.Flags().StringP("backup_name", "", "", "Database of the clickhouse server")

	createBackupCmd.MarkFlagRequired("backup_name")
	createBackupCmd.MarkFlagRequired("user")
	createBackupCmd.MarkFlagRequired("password")
}


func connect(h, p, u, pass string) (driver.Conn, error) {
	var (
		ctx = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{h+":"+p},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: u,
			},
			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
		})
	)

	if err != nil {
		return nil, err
	}

	// check connection
	if err := conn.Ping(ctx); err != nil {
		// check if the exception is a clickhouse exception and print all the info
		// this is a type assertion
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil

}