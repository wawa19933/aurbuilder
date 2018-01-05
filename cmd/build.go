package cmd

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build PACKAGES...",
	Args:  cobra.MinimumNArgs(1),
	Short: "Build packages with 'aursync'",
	Long: `Example:
		aurbuilder build radare2-git radare2-cutter-git`,
	Run: func(cmd *cobra.Command, args []string) {
		if aurBuild(args) != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func aurBuild(packages []string) error {
	info := color.New(color.FgCyan)
	error := color.New(color.FgRed)
	notice := color.New(color.FgYellow)

	yellow := notice.SprintFunc()
	red := error.SprintFunc()

	info.Print("Build AUR packages...")
	packages = append([]string{"--no-view", "--no-confirm", "--repo", "wawa19933"}, packages...)
	aur := exec.Command("aursync", packages...)
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
