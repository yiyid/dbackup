package cmd

import (
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "检测网络和磁盘",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(&backupParams.Host, "host", "H", "127.0.0.1", "The database host")
	checkCmd.Flags().StringVarP(&backupParams.Port, "port", "P", "3306", "The database port")
	checkCmd.Flags().StringVarP(&backupParams.Username, "username", "u", "", "Username for database connection")
	checkCmd.Flags().StringVarP(&backupParams.Password, "password", "p", "", "Password for database connection")
	checkCmd.Flags().StringVarP(&backupParams.Database, "database", "d", "", "Database name")
	checkCmd.MarkFlagRequired("username")
	checkCmd.MarkFlagRequired("password")
	checkCmd.MarkFlagRequired("database")
}
