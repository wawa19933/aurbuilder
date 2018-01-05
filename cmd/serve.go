package cmd

import (
	"os"
	"os/exec"

	"bytes"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve PATH",
	Args:  cobra.MinimumNArgs(1),
	Short: "Will start a simple HTTP server to serve PATH",
	Long: `Example:
		aurbuild serve /srv/repo
would start HTTP server on port 80 with access to /srv/repo`,
	Run: func(cmd *cobra.Command, args []string) {
		startPythonServer(args[1])
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func startPythonServer(dir string) {
	info := color.New(color.FgCyan)
	error := color.New(color.FgRed)
	notice := color.New(color.FgYellow)

	red := error.SprintFunc()

	info.Printf("Starting web server for %s on port 80...", dir)
	server := exec.Command("python", "-m", "http.server", "80")
	server.Dir = dir
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	server.Stdout = stdout
	server.Stderr = stderr
	//server.Stdout = os.Stdout
	//server.Stderr = os.Stderr

	if err := server.Run(); err != nil {
		error.Println("Failed!")
		notice.Println(stdout)
		log.Printf("Error: %s\nStdError:\n%s", red(err), red(stderr))
		os.Exit(1)
	}
	notice.Println(stdout)
}
