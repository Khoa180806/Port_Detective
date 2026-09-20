package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/lookup"
	"github.com/Khoa180806/Port_Detective/internal/output"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var force bool
var dryRun bool

var killCmd = &cobra.Command{
	Use:   "kill [port]",
	Short: "Buộc dừng các tiến trình đang chiếm giữ port",
	Long: `Kiểm tra xem cổng mạng (port) nào đang bị chiếm giữ và buộc dừng (kill)
tất cả các tiến trình (processes) đang sử dụng cổng mạng đó.

Mặc định lệnh này sẽ hiển thị thông tin tiến trình và hỏi bạn (y/N) trước khi kill.
Bạn có thể dùng cờ --force (-f) để bỏ qua câu hỏi xác nhận.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port, err := strconv.Atoi(args[0])
		if err != nil || port < 1 || port > 65535 {
			fmt.Println("Lỗi: Port phải là một số nguyên từ 1 đến 65535")
			os.Exit(2)
		}

		procs, err := lookup.FindProcessByPort(port)
		if err != nil {
			fmt.Printf("Lỗi khi tra cứu: %v\n", err)
			os.Exit(2)
		}

		if len(procs) == 0 {
			fmt.Printf("Không tìm thấy tiến trình nào đang chạy trên port %d\n", port)
			os.Exit(1)
		}

		// Hiển thị thông tin tiến trình
		color.New(color.FgRed, color.Bold).Println("CẢNH BÁO: Phát hiện các tiến trình sau đang chiếm dụng port:")
		fmt.Println(output.FormatTextMultiple(procs))
		fmt.Println()

		if dryRun {
			fmt.Println("[DRY RUN] Sẽ không thực hiện lệnh kill nào.")
			os.Exit(0)
		}

		if !force {
			fmt.Print("Bạn có chắc chắn muốn KILL tất cả các tiến trình trên không? [y/N]: ")
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\nĐã hủy.")
				os.Exit(1)
			}
			response = strings.TrimSpace(strings.ToLower(response))
			if response != "y" && response != "yes" {
				fmt.Println("Đã hủy thao tác.")
				os.Exit(0)
			}
		}

		hasError := false
		for _, p := range procs {
			fmt.Printf("Đang kill PID %d (%s)... ", p.PID, p.Name)
			err := lookup.KillProcess(p.PID)
			if err != nil {
				color.Red("THẤT BẠI: %v", err)
				hasError = true
			} else {
				color.Green("THÀNH CÔNG")
			}
		}

		if hasError {
			os.Exit(2)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(killCmd)
	killCmd.Flags().BoolVarP(&force, "force", "f", false, "Buộc dừng tiến trình không cần xác nhận (Bỏ qua y/N)")
	killCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Chỉ hiển thị các tiến trình sẽ bị kill mà không thực thi")
}
