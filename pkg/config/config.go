package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetRootLink() string {
	return viper.GetString("LINK")
}

func RootLinkIsSet() bool {
	return GetRootLink() != ""
}

func ErrRootLinkNotSet() {
	if !RootLinkIsSet() {
		fmt.Println("Error: Cannot execute because link is not set")
		fmt.Println("Please use `shorten config --link` to add a CLI shortener remote link")
		return
	}
}
