package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var checkNetCmd = &cobra.Command{
	Use:   "net",
	Short: "检查网络连通性",
	Run: func(cmd *cobra.Command, args []string) {
		err := CheckNet(backupParams)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("连接 %s:%s 成功", backupParams.Host, backupParams.Port)
	},
}

func init() {
	checkCmd.AddCommand(checkNetCmd)
	checkNetCmd.Flags().StringVarP(&backupParams.Host, "host", "H", "127.0.0.1", "The database host")
	checkNetCmd.Flags().StringVarP(&backupParams.Port, "port", "P", "3306", "The database port")
	checkNetCmd.Flags().StringVarP(&backupParams.Username, "username", "u", "", "Username for database connection")
	checkNetCmd.Flags().StringVarP(&backupParams.Password, "password", "p", "", "Password for database connection")
	checkNetCmd.Flags().StringVarP(&backupParams.Database, "database", "d", "", "Database name")
	checkNetCmd.MarkFlagRequired("username")
	checkNetCmd.MarkFlagRequired("password")
	checkNetCmd.MarkFlagRequired("database")
}

func CheckNet(backupParams BackupParams) error {
	// 检查 MySQL 是否可以连通
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", backupParams.Username, backupParams.Password, backupParams.Host, backupParams.Port, backupParams.Database)
	client, err := sql.Open("mysql", dbConnStr)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL: %v", err)

	}
	defer client.Close()

	err = client.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping MySQL: %v", err)
	}
	return nil
}
