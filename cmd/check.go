package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Khoa180806/Port_Detective/internal/lookup"
	"github.com/Khoa180806/Port_Detective/internal/output"
	"github.com/spf13/cobra"
)

var jsonOutput bool

var checkCmd = &cobra.Command{
	Use:   "check [port]",
	Short: "Kiểm tra tiến trình đang chạy trên một port",
	Long: `Kiểm tra và hiển thị thông tin về các tiến trình (processes) 
đang chiếm giữ hoặc lắng nghe trên cổng mạng (port) được chỉ định.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port, err := strconv.Atoi(args[0])
		if err != nil || port < 1 || port > 65535 {
			fmt.Println("Lỗi: Port phải là một số nguyên từ 1 đến 65535")
			os.Exit(2)
		}

		procs, err := lookup.FindProcessByPort(port)
		if err != nil {
			if jsonOutput {
				fmt.Printf(`{"error": "%v"}`+"\n", err)
			} else {
				if errors.Is(err, lookup.ErrPermissionDenied) {
					fmt.Println("Lỗi: Không đủ quyền truy cập để lấy thông tin chi tiết.")
					fmt.Println("Gợi ý: Hãy thử chạy lại lệnh dưới quyền Administrator (hoặc sudo trên Linux/macOS).")
				} else {
					fmt.Printf("Lỗi khi tra cứu: %v\n", err)
				}
			}
			os.Exit(2)
		}

		if len(procs) == 0 {
			if jsonOutput {
				fmt.Println("[]")
			} else {
				fmt.Printf("Không tìm thấy tiến trình nào đang chạy trên port %d\n", port)
			}
			os.Exit(1) // Theo spec: Exit code 1 = port free
		}

		if jsonOutput {
			out, err := output.FormatJSONMultiple(procs)
			if err != nil {
				fmt.Printf(`{"error": "%v"}`+"\n", err)
				os.Exit(2)
			}
			fmt.Println(out)
		} else {
			out := output.FormatTextMultiple(procs)
			fmt.Println(out)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Hiển thị kết quả dưới định dạng JSON")
}
