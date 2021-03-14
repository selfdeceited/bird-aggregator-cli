package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	apiKey  string
	userId  string

	rootCmd = &cobra.Command{
		Use:   "bird-aggregator-cli",
		Short: "Simple cli interface accessing Flickr API to get birds statistics.",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.cobra)")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "apiKey", "k", "", "Flickr API key")
	rootCmd.PersistentFlags().StringVarP(&userId, "userId", "u", "", "Flickr user ID")

	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("license", "apache")
	viper.SetConfigType("yaml")

	rootCmd.AddCommand(fetchCommand)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		apiKey = viper.GetString("apiKey")
		userId = viper.GetString("userId")
	} else {
		rootCmd.MarkPersistentFlagRequired("apiKey")
		rootCmd.MarkPersistentFlagRequired("userId")
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
