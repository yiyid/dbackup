/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type BackupParams struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var backupParams BackupParams

var rootCmd = &cobra.Command{
	Use:   "dbackup",
	Short: "数据库备份工具",
}

func init() {

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
