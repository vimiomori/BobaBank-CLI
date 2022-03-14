package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
)


func main() {
	dirName := ""

	homePath := os.Getenv("HOME")

	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.Set("root", nil)
			viper.WriteConfigAs("config.yaml")
		} else {
			fmt.Println(err.Error())
		}
	}

	bankPath := viper.GetString("root")
	if bankPath == "" {
		initBank := &survey.Input{
			Message: "Input the path to store your boba bank configurations and data:",
			Default: filepath.Join(homePath, ".boba-bank"),
		}

		err = survey.AskOne(initBank, &dirName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		viper.Set("root", dirName)
		viper.WriteConfigAs("config.yaml")
		fmt.Printf("Your boba bank information will be stored under %s", dirName)
	}
}
