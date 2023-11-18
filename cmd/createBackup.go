/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/ClickHouse/clickhouse-go"
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
		fmt.Println("createBackup called")

		// Connect to clichouse database
		connect, err := sql.Open("clickhouse", "tcp://localhost:9000?debug=true")
		if err != nil {
			fmt.Println("Error connecting to database: ", err)
			return
		}

		// check connection
		if err := connect.Ping(); err != nil {
			// check if the exception is a clickhouse exception and print all the info
			// this is a type assertion
			if exception, ok := err.(*clickhouse.Exception); ok {
				fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
			} else {
				fmt.Println(err)
			}
			return
		}

		_, err = connect.Exec("BACKUP TABLE my_table TO Disk('backups', '1.zip')")
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
	createBackupCmd.Flags().StringP("user", "u", "default", "User of the clickhouse server")
}
