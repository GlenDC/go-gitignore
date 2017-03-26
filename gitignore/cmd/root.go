package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// persistent flags
var (
	cfgFile string
	ghtoken string
	ghrepo  string
	pkind   providerKind
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gitignore",
	Short: "Create and list gitignore files.",
	Long: `Create and list gitignore files using gitignore files
from github.com/github/gitignore.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	RootCmd.AddCommand(
		versionCmd,
		listCmd,
		createCmd,
		printCmd,
	)

	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// global flags

	RootCmd.PersistentFlags().StringVar(
		&cfgFile, "config", "",
		"config file (default is $HOME/.go-gitignore.yaml)")
	RootCmd.PersistentFlags().Var(&pkind, "provider",
		"defines the provider to use for getting gitignore content, default: github, options: github")
	RootCmd.PersistentFlags().StringVar(
		&ghtoken, "github-token", "",
		"github token used for some commands in case github provider is used")
	RootCmd.PersistentFlags().StringVar(
		&ghrepo, "github-repo", "github/gitignore",
		"github repo used in case github provider is used")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".go-gitignore") // name of config file (without extension)
	viper.AddConfigPath("$HOME")         // adding home directory as first search path
	viper.AutomaticEnv()                 // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
