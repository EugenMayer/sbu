package gen2plus

import (
	"fmt"
	"github.com/fermayo/shelly-bulk-update/cli"
	"time"
)

func UpdateShelly(name, address string) {
	prefix := fmt.Sprintf("[%s/%s/gen2+]", name, address)
	// First, we trigger a check for updates
	fmt.Printf("%s checking for updates...\n", prefix)
	updates, err := doCheckForUpdateRequest(address)
	if err != nil {
		fmt.Printf("%s failed to check for updates: %s, aborting...\n", prefix, err)
		return
	}

	updateVersion := updates.Stable.Version
	if *cli.Channel == "beta" {
		updateVersion = updates.Beta.Version
	}
	if updateVersion == "" {
		fmt.Printf("%s already up to date\n", prefix)
		return
	}
	newVersion := updateVersion

	fmt.Printf("%s updating to version %s...\n", prefix, updateVersion)
	err = doUpdateRequest(address, *cli.Channel)
	if err != nil {
		fmt.Printf("%s failed to update: %s, aborting...\n", prefix, err)
		return
	}

	// wait for update to complete
	tries := 0
	for updateVersion != "" {
		tries++
		if tries > 12 {
			fmt.Printf("%s failed to check if update completed successfully", prefix)
			return
		}
		time.Sleep(time.Second * 5)
		updates, err := doCheckForUpdateRequest(address)
		if err != nil {
			fmt.Printf("%s failed to query update status: %s, retrying...\n", prefix, err)
			continue
		}

		if *cli.Channel == "beta" {
			updateVersion = updates.Beta.Version
		} else {
			updateVersion = updates.Stable.Version
		}
	}
	fmt.Printf("%s device updated to %s!\n", prefix, newVersion)
}
