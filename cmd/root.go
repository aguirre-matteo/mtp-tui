package cmd

import (
	"github.com/aguirre-matteo/mtp-tui/app"
	"github.com/aguirre-matteo/mtp-tui/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mtp-tui",
	Short: "A TUI application for easily mounting your MTP devices!",
	Long: `This app uses Jmtpfs for mounting MTP devices,
  and Bubbletea for creating an easy to use UI.`,
	Run: func(cmd *cobra.Command, args []string) {
		user, err := cmd.Flags().GetString("user")
		if err != nil {
			panic(err)
		}

		err = config.InitViper(user)
		if err != nil {
			panic(err)
		}

		err = app.Run()
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("user", "u", "root", "User which will run the app")
}
