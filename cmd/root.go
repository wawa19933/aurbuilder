package cmd

import (
	"fmt"
	"os"

	"bytes"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os/exec"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aurbuilder",
	Short: "Runs an AUR builder",
	Long: `Control custom repository with aurutils and repose
aurbuilder build PACKAGE...
aurbuilder update
aurbuilder serve PATH
aurbuilder clean

By default executes 'aurbuilder serve /repository'`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		UpdatePacmanDatabase()
		startPythonServer("/repository")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aurbuilder.yaml)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".aurbuilder" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".aurbuilder")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func UpdatePacmanDatabase() {
	info := color.New(color.FgCyan)
	error := color.New(color.FgRed)
	notice := color.New(color.FgYellow)

	yellow := notice.SprintFunc()
	red := error.SprintFunc()

	info.Print("Updating pacman database...")

	var pacman *exec.Cmd
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	if os.Geteuid() != 0 {
		fmt.Println()

		pacman = exec.Command("sudo", "sh", "-c", "pacman -Sy --noconfirm >/dev/null")
		pacman.Stdin = os.Stdin
		pacman.Stdout = os.Stdout
		pacman.Stderr = os.Stderr
	} else {
		pacman = exec.Command("pacman", "-Sy", "--noconfirm")
		pacman.Stdout = stdout
		pacman.Stderr = stderr
	}

	if err := pacman.Run(); err != nil {
		error.Println("Failed!")
		log.Printf("Error: %s\nStdError:\n%s\nStdOut:\n%s\n", red(err), red(stderr.String()),
			yellow(stdout.String()))
	}
	notice.Println("OK!")
}

func UpdatePacmanPackages() {
	info := color.New(color.FgCyan)
	error := color.New(color.FgRed)
	notice := color.New(color.FgYellow)

	yellow := notice.SprintFunc()
	red := error.SprintFunc()

	info.Print("Updating packages with pacman...")

	var pacman *exec.Cmd
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	if os.Geteuid() != 0 {
		fmt.Println()

		pacman = exec.Command("sudo", "sh", "-c", "pacman -Syu --noconfirm >/dev/null")
		pacman.Stdin = os.Stdin
		pacman.Stdout = os.Stdout
		pacman.Stderr = os.Stderr
	} else {
		pacman = exec.Command("pacman", "-Syu", "--noconfirm")
		pacman.Stdout = stdout
		pacman.Stderr = stderr
	}

	if err := pacman.Run(); err != nil {
		error.Println("Failed!")
		log.Printf("Error: %s\nStdError:\n%s\nStdOut:\n%s\n", red(err), red(stderr.String()),
			yellow(stdout.String()))
	}
	notice.Println("OK!")
}
