package main

import (
	"fmt"
	"github.com/fermayo/shelly-bulk-update/cli"
	"github.com/fermayo/shelly-bulk-update/config"
	"github.com/fermayo/shelly-bulk-update/update"
	"log"
	"os"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli.ParseArgs()

	if *cli.ShowVersion {
		fmt.Printf("sbu %s (%s %s)\n", version, commit, date)
		os.Exit(0)
	}

	if *cli.Password != "" {
		fmt.Printf("Using authentication")
	} else if config.UserConfigExistsInHome() {
		userConfig, err := config.LoadUserConfigFromHome()
		if err != nil {
			log.Fatalln("Failed loading user config", err.Error())
		}
		cli.Username = &userConfig.GlobalConfig.DefaultCredentials.Username
		cli.Password = &userConfig.GlobalConfig.DefaultCredentials.Password
	}

	if cli.Hosts != nil && len(*cli.Hosts) > 0 {
		update.SpecificHosts(*cli.Hosts)
	} else {
		update.AutoDiscoverUsingAndUpdate()
	}
}
