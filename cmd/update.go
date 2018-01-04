// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os/exec"
	"log"
	"os"
	"bytes"
	"github.com/fatih/color"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pacmanUpdate()
		if aurUpdate() != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func pacmanUpdate () error {
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

func aurUpdate () error {
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