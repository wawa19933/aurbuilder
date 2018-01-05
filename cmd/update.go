package cmd

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates custom repository with 'aursync -u'",
	Run: func(cmd *cobra.Command, args []string) {
		UpdatePacmanPackages()
		if aurUpdate() != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func aurUpdate() error {
	info := color.New(color.FgCyan)
	error := color.New(color.FgRed)
	notice := color.New(color.FgYellow)

	yellow := notice.SprintFunc()
	red := error.SprintFunc()

	info.Print("Updating AUR packages...")
	aur := exec.Command("aursync", "--no-view", "--no-confirm", "--repo", "wawa19933", "-u")
	out := new(bytes.Buffer)
	aur.Stdout = out
	stderr := new(bytes.Buffer)
	aur.Stderr = stderr

	if err := aur.Run(); err != nil {
		error.Println("Failed!")

		log.Printf("Error: %s\nStdout:\n-------\n%s\nStdError:\n--------\n%s\n",
			red(err), yellow(out.String()), red(stderr.String()))
		return err
	} else {
		notice.Println("OK!")
	}

	return nil
}
