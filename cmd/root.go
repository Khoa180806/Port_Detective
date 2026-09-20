package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "port-detective",
	Short: "Port Detective - CLI tool giúp tra cứu và quản lý các cổng mạng (ports)",
	Long: `Port Detective là một công cụ dòng lệnh (CLI) cross-platform được viết bằng Go. 
Công cụ giúp bạn nhanh chóng tìm ra tiến trình (process) nào đang chiếm giữ một cổng mạng (port) cụ thể 
và hỗ trợ buộc dừng (kill) tiến trình đó nếu cần thiết.

Hỗ trợ các hệ điều hành: Windows, macOS, Linux.`,
	Version: "1.0.0",
	// Root command không thực hiện hành động gì nếu không gọi subcommand (check/kill)
}

// Execute là entry point cho tất cả các commands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
