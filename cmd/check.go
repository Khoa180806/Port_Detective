package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/Khoa180806/Port_Detective/internal/lookup"
	"github.com/Khoa180806/Port_Detective/internal/output"
	"github.com/spf13/cobra"
)

var jsonOutput bool

var checkCmd = &cobra.Command{
	Use:   "check [port]",
	Short: i18n.T("check.short"),
	Long:  i18n.T("check.long"),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port, err := strconv.Atoi(args[0])
		if err != nil || port < 1 || port > 65535 {
			fmt.Println(i18n.T("check.error.invalid_port"))
			os.Exit(2)
		}

		procs, err := lookup.FindProcessByPort(port)
		if err != nil {
			if jsonOutput {
				fmt.Printf(`{"error": "%v"}`+"\n", err)
			} else {
				if errors.Is(err, lookup.ErrPermissionDenied) {
					fmt.Println(i18n.T("check.error.permission"))
					fmt.Println(i18n.T("check.error.permission_hint"))
				} else {
					fmt.Println(i18n.Tr("check.error.lookup", err))
				}
			}
			os.Exit(2)
		}

		if len(procs) == 0 {
			if jsonOutput {
				fmt.Println("[]")
			} else {
				fmt.Println(i18n.Tr("check.not_found", port))
			}
			os.Exit(1) // Exit code 1: port is free
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
	checkCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, i18n.T("flag.json"))
}
