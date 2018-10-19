package cmd

import (
	"path/filepath"

	"github.com/blocklayerhq/chainkit/pkg/ui"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		rootDir := getCwd(cmd)
		name := filepath.Base(rootDir)
		start(name, rootDir)
	},
}

func init() {
	startCmd.Flags().String("cwd", ".", "specifies the current working directory")

	rootCmd.AddCommand(startCmd)
}

func start(name, rootDir string) {
	ui.Info("Starting %s", name)
	ui.Success("Application is live at http://localhost:26657/")
	if err := dockerRun(rootDir, name, "start"); err != nil {
		ui.Fatal("Failed to start the application: %v", err)
	}
}
