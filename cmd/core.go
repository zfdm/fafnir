package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FafnirCmd = &cobra.Command{
	Use:   "fafnir",
	Short: "Companion to hoard, fafnir manages storage devices.",
	Long: versionInfo + `

Fafnir manages the storage hardware that run the hoard.
`,
}

func init() {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("fafnir")

	cobra.OnInitialize(initConfig)
	FafnirCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/hoard/fafnir.toml or $HOME/.fafnir.toml)")
	RootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")

	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")

	if cfgFile == nil {

		viper.SetConfigName("fafnir")
		viper.AddConfigPath("/etc/hoard/")
		viper.AddConfigPath("./")
	} else {
		viper.SetConfigName(cfgFile)
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("Author is ", viper.GetString("author"))

}

func Execute() {

	if err := FafnirCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
