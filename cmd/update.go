package cmd

import (
	"fmt"

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
		pacmanUpdate()
		if aurUpdate() != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func pacmanUpdate() error {
	info := color.New(color.FgCyan)
	error := color.New(color.FgRed)
	notice := color.New(color.FgYellow)

	yellow := notice.SprintFunc()
	red := error.SprintFunc()

	info.Print("Updating packages...")
	var pacman *exec.Cmd
	if os.Geteuid() != 0 {
		fmt.Println()

		pacman = exec.Command("sudo", "sh", "-c", "pacman -Syu --noconfirm >/dev/null")
		pacman.Stdin = os.Stdin
		pacman.Stdout = os.Stdout
		pacman.Stderr = os.Stderr
	} else {
		pacman = exec.Command("pacman", "-Syu", "--noconfirm")
	}

	if err := pacman.Run(); err != nil {
		error.Println("Failed!")

		log.Printf("Error: %s!\nStdOut:\n-------\n%s\nStdError:\n--------\n%s\n",
			err, yellow(pacman.Stdout), red(pacman.Stderr))
		return err
	} else {
		notice.Println("OK!")
	}

	return nil
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
