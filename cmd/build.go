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
	"bytes"
	"os"
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build PACKAGES...",
    Args: cobra.MinimumNArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        if aurBuild(args) != nil {
            os.Exit(1)
        }
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
