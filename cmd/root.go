package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Khoa180806/Port_Detective/internal/i18n"
	"github.com/spf13/cobra"
)

var langFlag string

var rootCmd = &cobra.Command{
	Use:     "port-detective",
	Short:   i18n.T("root.short"),
	Long:    i18n.T("root.long"),
	Version: "1.1.0",
}

// updateCommandDescriptions updates descriptions for root and all registered subcommands.
func updateCommandDescriptions(cmd *cobra.Command) {
	cmd.Root().Short = i18n.T("root.short")
	cmd.Root().Long = i18n.T("root.long")

	checkCmd.Short = i18n.T("check.short")
	checkCmd.Long = i18n.T("check.long")

	killCmd.Short = i18n.T("kill.short")
	killCmd.Long = i18n.T("kill.long")

	scanCmd.Short = i18n.T("scan.short")
	scanCmd.Long = i18n.T("scan.long")
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&langFlag, "lang", "l", "", i18n.T("flag.lang"))

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Determine language: flag > environment variable > default ("en")
		targetLang := langFlag
		if targetLang == "" {
			targetLang = os.Getenv("PORT_DETECTIVE_LANG")
		}
		if targetLang == "" {
			targetLang = "en"
		}

		if err := i18n.SetLang(targetLang); err != nil {
			return err
		}

		// Dynamically update command descriptions according to selected language
		updateCommandDescriptions(cmd)
		return nil
	}
}

// Execute is the entry point for all commands.
func Execute() {
	// Early language detection from args or env to properly render Help and Usage
	targetLang := os.Getenv("PORT_DETECTIVE_LANG")
	for i, arg := range os.Args {
		if (arg == "--lang" || arg == "-l") && i+1 < len(os.Args) {
			targetLang = os.Args[i+1]
			break
		} else if strings.HasPrefix(arg, "--lang=") {
			targetLang = strings.TrimPrefix(arg, "--lang=")
			break
		}
	}
	if targetLang != "" && i18n.IsSupported(targetLang) {
		_ = i18n.SetLang(targetLang)
		updateCommandDescriptions(rootCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
