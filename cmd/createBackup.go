package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var createBackupCmd = &cobra.Command{
	Use:   "create",
	Short: "创建备份文件",
	Run: func(cmd *cobra.Command, args []string) {
		err := CheckNet(backupParams)
		if err != nil {
			log.Fatal(err)
		}
		CreateBackup(backupParams)
	},
}

func init() {
	rootCmd.AddCommand(createBackupCmd)
	createBackupCmd.Flags().StringVarP(&backupParams.Host, "host", "H", "127.0.0.1", "The database host")
	createBackupCmd.Flags().StringVarP(&backupParams.Port, "port", "P", "3306", "The database port")
	createBackupCmd.Flags().StringVarP(&backupParams.Username, "username", "u", "", "Username for database connection")
	createBackupCmd.Flags().StringVarP(&backupParams.Password, "password", "p", "", "Password for database connection")
	createBackupCmd.Flags().StringVarP(&backupParams.Database, "database", "d", "", "Database name")
	createBackupCmd.MarkFlagRequired("username")
	createBackupCmd.MarkFlagRequired("password")
	createBackupCmd.MarkFlagRequired("database")
}

func CreateBackup(backupParams BackupParams) {

	backupDir := "."
	backupFileName := fmt.Sprintf("%s-%s.sql", backupParams.Database, time.Now().Format("20060102150405"))
	command := fmt.Sprintf("mysqldump -h%s -P%s -u%s -p%s %s", backupParams.Host, backupParams.Port, backupParams.Username, backupParams.Password, backupParams.Database)
	backupCmd := exec.Command("bash", "-c", command)

	// 打开备份文件
	os.Chdir(backupDir)
	backupFile, err := os.Create(backupFileName)
	if err != nil {
		log.Fatalf("Failed to create backup file: %v", err)
	}
	defer backupFile.Close()
	backupCmd.Stdout = backupFile

	// 执行备份命令
	err = backupCmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute backup command: %v", err)
	}

	log.Printf("Database backup completed. File saved as %s", backupFileName)
}
