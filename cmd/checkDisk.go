package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "检测磁盘",
	Run: func(cmd *cobra.Command, args []string) {
		err := CheckDisk(".")
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("磁盘空间充足")
		}
	},
}

func init() {
	checkCmd.AddCommand(diskCmd)
}

func CheckDisk(path string) error {

	absPath, _ := filepath.Abs(path)

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return fmt.Errorf("failed to get disk space information: %v", err)
	}

	// 获取分区剩余空间的字节数
	// 使用 uint64(fs.Bavail) * uint64(fs.Bsize) 计算分区剩余空间（字节数）
	// Bavail 代表非超级用户可用的文件系统空间块数
	// Bsize 代表文件系统的块大小（字节数）
	availableSpace := uint64(fs.Bavail) * uint64(fs.Bsize)
	availableSpaceGB := float64(availableSpace) / float64(1024) / float64(1024) / float64(1024) // 将字节数转换为GB

	minSpace := float64(5 * 1024 * 1024 * 1024)                                     // 10G
	minSpaceGB := float64(minSpace) / float64(1024) / float64(1024) / float64(1024) // // 将字节数转换为GB

	if availableSpaceGB < minSpaceGB {
		return fmt.Errorf("目录 %s 所在分区磁盘空间 %dGB 小于额定空间 %dGB。磁盘不足，谨慎备份。", absPath, int(availableSpaceGB), int(minSpaceGB))
	}

	return nil
}
