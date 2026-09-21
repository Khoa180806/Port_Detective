package cmd

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/Khoa180806/Port_Detective/internal/lookup"
	"github.com/Khoa180806/Port_Detective/internal/output"
	"github.com/Khoa180806/Port_Detective/internal/process"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [start]-[end]",
	Short: i18n.T("scan.short"),
	Long:  i18n.T("scan.long"),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parts := strings.Split(args[0], "-")
		if len(parts) != 2 {
			color.Red(i18n.T("scan.error.format"))
			os.Exit(2)
		}

		startPort, err1 := strconv.Atoi(parts[0])
		endPort, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil || startPort < 1 || endPort > 65535 || startPort > endPort {
			color.Red(i18n.T("scan.error.range"))
			os.Exit(2)
		}

		if endPort-startPort > 5000 {
			color.Red(i18n.T("scan.error.too_large"))
			os.Exit(2)
		}

		var allProcs []process.ProcessInfo
		var mu sync.Mutex
		var permissionErrorFlag bool

		// Task queue
		portChan := make(chan int, endPort-startPort+1)
		var wg sync.WaitGroup

		// Worker pool
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

		// Enqueue ports
		for p := startPort; p <= endPort; p++ {
			portChan <- p
		}
		close(portChan)

		if !jsonOutput {
			fmt.Println(i18n.Tr("scan.scanning", startPort, endPort))
		}

		wg.Wait()

		// Sort results by port
		sort.Slice(allProcs, func(i, j int) bool {
			return allProcs[i].Port < allProcs[j].Port
		})

		if len(allProcs) == 0 {
			if jsonOutput {
				fmt.Println("[]")
			} else {
				color.Green(i18n.T("scan.no_results"))
			}
			os.Exit(1)
		}

		// Output results
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

			// Warn if permission errors were encountered
			if permissionErrorFlag {
				color.Yellow(i18n.T("scan.permission_warn"))
				color.Yellow(i18n.T("scan.permission_hint"))
			}
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, i18n.T("flag.json"))
}
