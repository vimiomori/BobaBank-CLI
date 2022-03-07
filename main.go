package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
)

// the questions to ask
// var qs = []*survey.Question{
//     {
//         Name:     "name",
//         Prompt:   &survey.Input{Message: "What is your name?"},
//         Validate: survey.Required,
//         Transform: survey.Title,
//     },
//     {
//         Name: "color",
//         Prompt: &survey.Select{
//             Message: "Choose a color:",
//             Options: []string{"red", "blue", "green"},
//             Default: "red",
//         },
//     },
//     {
//         Name: "age",
//         Prompt:   &survey.Input{Message: "How old are you?"},
//     },
// }

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
