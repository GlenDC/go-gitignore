package cmd

import (
	"fmt"
	"log"
	"os/user"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// persistent flags
var (
	cfgFile string

	// github mode
	ghtoken string
	ghrepo  string
	pkind   providerKind

	// local mode
	localPath string
)

// config constants
const (
	defProvider   = "github"
	defGithubRepo = "github/gitignore"
)

var (
	defCfgFile = func(home string) string {
		if usr, err := user.Current(); err == nil {
			home = usr.HomeDir
		}

		return path.Join(home, ".go-gitignore.yaml")
	}("")
)

// config names
const (
	cfgProviderKey    = "provider"
	cfgGithubRepoKey  = "github.repository"
	cfgGithubTokenKey = "github.token"
	cfgLocalPathKey   = "local.path"
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
	cobra.OnInitialize(initConfig)

	// global flags

	RootCmd.PersistentFlags().StringVar(
		&cfgFile, "config", "",
		fmt.Sprintf("config file used (default: %s)", defCfgFile))

	RootCmd.PersistentFlags().Var(&pkind, "provider",
		fmt.Sprintf("defines the provider to use for getting gitignore content, options: github, local (default: %s)", defProvider))

	RootCmd.PersistentFlags().StringVar(
		&ghtoken, "github-token", "",
		"github token used for some commands in case github provider is used")
	RootCmd.PersistentFlags().StringVar(
		&ghrepo, "github-repo", "",
		fmt.Sprintf("github repo used in case github provider is used (default: %s)", defGithubRepo))

	RootCmd.PersistentFlags().StringVar(
		&localPath, "local-path", "",
		"local dir used in case local provider is used")
}

func setCfgDefaults() {
	if len(pkind) == 0 {
		pkind = providerKind(defProvider)
	}
	if ghrepo == "" {
		ghrepo = defGithubRepo
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		cfgFile = defCfgFile
	}
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		//fmt.Printf("couldn't load config file %q: %s", cfgFile, err) // TODO: LOG
		setCfgDefaults() // revert to defaults if needed
		return
	}
	//fmt.Println("Using config file:", viper.ConfigFileUsed()) // TODO: LOG

	viper.SetDefault(cfgGithubRepoKey, defGithubRepo)
	viper.SetDefault(cfgProviderKey, defProvider)

	if len(pkind) == 0 {
		pkind = providerKind(viper.GetString(cfgProviderKey))
	}

	if ghtoken == "" {
		ghtoken = viper.GetString(cfgGithubTokenKey)
	}
	if ghrepo == "" {
		ghrepo = viper.GetString(cfgGithubRepoKey)
	}

	if localPath == "" {
		localPath = viper.GetString(cfgLocalPathKey)
	}
}
