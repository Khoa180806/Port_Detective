package cmd

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/Khoa180806/Port_Detective/internal/lookup"
	"github.com/Khoa180806/Port_Detective/internal/output"
	"github.com/Khoa180806/Port_Detective/internal/process"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [start]-[end]",
	Short: "Quét một dải cổng mạng (port range) để tìm các tiến trình đang hoạt động",
	Long: `Lệnh scan giúp bạn kiểm tra hàng loạt cổng mạng trong một dải cụ thể.
Ví dụ: port-detective scan 3000-3010
Dải cổng hợp lệ là từ 1 đến 65535. Để bảo vệ hệ thống, giới hạn tối đa 5000 cổng mỗi lần quét.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parts := strings.Split(args[0], "-")
		if len(parts) != 2 {
			color.Red("Lỗi: Sai định dạng dải port. Vui lòng sử dụng định dạng <start>-<end> (vd: 3000-3010).")
			os.Exit(2)
		}

		startPort, err1 := strconv.Atoi(parts[0])
		endPort, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil || startPort < 1 || endPort > 65535 || startPort > endPort {
			color.Red("Lỗi: Dải port không hợp lệ. (1 <= start <= end <= 65535).")
			os.Exit(2)
		}

		if endPort-startPort > 5000 {
			color.Red("Lỗi: Khoảng quét quá lớn (tối đa 5000 port mỗi lần) để tránh quá tải hệ thống.")
			os.Exit(2)
		}

		var allProcs []process.ProcessInfo
		var mu sync.Mutex
		var permissionErrorFlag bool // Cờ đánh dấu nếu gặp lỗi quyền truy cập

		// Hàng đợi công việc
		portChan := make(chan int, endPort-startPort+1)
		var wg sync.WaitGroup

		// Số lượng worker (goroutines)
		numWorkers := 20
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for p := range portChan {
					procs, err := lookup.FindProcessByPort(p)
					if err != nil {
						if errors.Is(err, lookup.ErrPermissionDenied) {
							mu.Lock()
							permissionErrorFlag = true
							mu.Unlock()
						}
					} else if len(procs) > 0 {
						mu.Lock()
						allProcs = append(allProcs, procs...)
						mu.Unlock()
					}
				}
			}()
		}

		// Nạp các port vào channel
		for p := startPort; p <= endPort; p++ {
			portChan <- p
		}
		close(portChan)

		// Hiển thị tiến trình giả lập (hoặc có thể dùng thư viện spinner, ở đây giữ đơn giản)
		if !jsonOutput {
			fmt.Printf("Đang quét dải port từ %d đến %d...\n", startPort, endPort)
		}

		wg.Wait()

		// Sắp xếp lại theo số port
		sort.Slice(allProcs, func(i, j int) bool {
			return allProcs[i].Port < allProcs[j].Port
		})

		if len(allProcs) == 0 {
			if jsonOutput {
				fmt.Println("[]")
			} else {
				color.Green("Tuyệt vời! Không có tiến trình nào đang chiếm dụng dải port này.")
			}
			os.Exit(1)
		}

		// In ra màn hình
		if jsonOutput {
			out, err := output.FormatJSONMultiple(allProcs)
			if err != nil {
				fmt.Printf(`{"error": "%v"}`+"\n", err)
				os.Exit(2)
			}
			fmt.Println(out)
		} else {
			out := output.FormatTextMultiple(allProcs)
			fmt.Println(out)
			
			// Cảnh báo nếu gặp thiếu quyền ở bất kỳ port nào
			if permissionErrorFlag {
				color.Yellow("\n[Chú ý] Quá trình quét bị từ chối truy cập (Permission Denied) ở một số port.")
				color.Yellow("Gợi ý: Hãy chạy lại công cụ dưới quyền Administrator/sudo để có danh sách đầy đủ nhất.")
			}
		}
		
		// Khác với check/kill (bị exit do OS), ở đây exit 0
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	// Tái sử dụng cờ jsonOutput đã được định nghĩa trong check.go
	// Nhưng vì cobra gán cờ theo từng command, ta phải định nghĩa lại. 
	// (Hoặc bind vào PersistentFlags của root). 
	// Do biến jsonOutput dùng chung (var ở root/check), ta có thể gắn thẳng vào cờ của scanCmd.
	scanCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Hiển thị kết quả dưới định dạng JSON")
}
