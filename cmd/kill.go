package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/Khoa180806/Port_Detective/internal/lookup"
	"github.com/Khoa180806/Port_Detective/internal/output"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var force bool
var dryRun bool

var killCmd = &cobra.Command{
	Use:   "kill [port]",
	Short: i18n.T("kill.short"),
	Long:  i18n.T("kill.long"),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port, err := strconv.Atoi(args[0])
		if err != nil || port < 1 || port > 65535 {
			fmt.Println(i18n.T("kill.error.invalid_port"))
			os.Exit(2)
		}

		procs, err := lookup.FindProcessByPort(port)
		if err != nil {
			if errors.Is(err, lookup.ErrPermissionDenied) {
				color.Red(i18n.T("check.error.permission"))
				color.Yellow(i18n.T("check.error.permission_hint"))
			} else {
				fmt.Println(i18n.Tr("kill.error.lookup", err))
			}
			os.Exit(2)
		}

		if len(procs) == 0 {
			fmt.Println(i18n.Tr("kill.not_found", port))
			os.Exit(1)
		}

		// Display identified processes occupying the port
		color.New(color.FgRed, color.Bold).Println(i18n.T("kill.warning"))
		fmt.Println(output.FormatTextMultiple(procs))
		fmt.Println()

		if dryRun {
			fmt.Println(i18n.T("kill.dry_run"))
			os.Exit(0)
		}

		if !force {
			fmt.Print(i18n.T("kill.confirm"))
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\n" + i18n.T("kill.cancelled"))
				os.Exit(1)
			}
			response = strings.TrimSpace(strings.ToLower(response))
			if response != "y" && response != "yes" {
				fmt.Println(i18n.T("kill.cancelled"))
				os.Exit(0)
			}
		}

		hasError := false
		for _, p := range procs {
			fmt.Print(i18n.Tr("kill.killing", p.PID, p.Name))
			err := lookup.KillProcess(p.PID)
			if err != nil {
				if errors.Is(err, lookup.ErrPermissionDenied) {
					color.Red(i18n.T("kill.failed.permission"))
				} else {
					color.Red(i18n.Tr("kill.failed", err))
				}
				hasError = true
			} else {
				color.Green(i18n.T("kill.success"))
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
	killCmd.Flags().BoolVarP(&force, "force", "f", false, i18n.T("flag.force"))
	killCmd.Flags().BoolVar(&dryRun, "dry-run", false, i18n.T("flag.dry_run"))
}
