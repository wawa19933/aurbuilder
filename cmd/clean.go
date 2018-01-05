package cmd

import (
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans the build cache after AUR build and pacman update",
	Run: func(cmd *cobra.Command, args []string) {
		if rmCache() != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func rmCache() error {
	info := color.New(color.FgCyan)
	error := color.New(color.FgRed)
	notice := color.New(color.FgYellow)

	yellow := notice.SprintFunc()
	red := error.SprintFunc()

	info.Print("Cleaning cache...")
	clean := exec.Command("rm", "-rf", "/var/cache/pacman/pkg", "/tmp/makepkg")
	if err := clean.Run(); err != nil {
		error.Println("Failed!")
		log.Printf("Error: %s\nStdOut:\n---------\n%s\nStdError:\n---------\n%s\n",
			red(err), yellow(clean.Stdout), red(clean.Stderr))
		return err
	} else {
		notice.Println("OK!")
	}
	return nil
}
