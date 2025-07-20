package gen1

import (
	"fmt"
	"github.com/fermayo/shelly-bulk-update/api"
	"github.com/fermayo/shelly-bulk-update/cli"
	"time"
)

func UpdateShelly(name, address string) {
	prefix := fmt.Sprintf("[%s/%s/gen1]", name, address)
	// First, we trigger a check for updates
	fmt.Printf("%s checking for updates...\n", prefix)
	_, err := triggerShellyUpdateCheck(address)
	if err != nil {
		fmt.Printf("%s failed to check for updates: %s, aborting...\n", prefix, err)
		return
	}

	// Check for updates is asynchronous, so we need to wait a bit
	time.Sleep(time.Second * 5)

	// Then, we check if there are any updates available
	updateStatus, err := checkShellyUpdateStatus(address)
	if err != nil {
		fmt.Printf("%s failed to query update status: %s, aborting...\n", prefix, err)
		return
	}

	// If there's an update available, trigger the update
	if (*cli.Channel == "stable" && updateStatus.HasUpdate) ||
		(*cli.Channel == "beta" && updateStatus.OldVersion != updateStatus.BetaVersion) {
		newVersion := updateStatus.NewVersion
		if *cli.Channel == "beta" {
			newVersion = updateStatus.BetaVersion
		}
		fmt.Printf(
			"%s update available! (%s -> %s), updating...\n",
			prefix, updateStatus.OldVersion, newVersion,
		)

		updateStatus, err := triggerShellyUpdate(address)
		if err != nil {
			fmt.Printf("%s failed to start update: %s, aborting...\n", prefix, err)
			return
		}

		for updateStatus.Status == "updating" {
			time.Sleep(time.Second * 5)
			updateStatusCheck, err := checkShellyUpdateStatus(address)
			if err != nil {
				fmt.Printf("%s failed to query update status: %s, retrying...\n", prefix, err)
				continue
			}
			updateStatus = updateStatusCheck
		}

		fmt.Printf("%s device updated to %s!\n", prefix, updateStatus.OldVersion)
	} else {
		fmt.Printf("%s already up to date (%s)\n", prefix, updateStatus.OldVersion)
	}
}

func triggerShellyUpdate(hostname string) (*api.ShellyUpdateStatusResponse, error) {
	return doShellyUpdateRequest(hostname, true)
}

func checkShellyUpdateStatus(hostname string) (*api.ShellyUpdateStatusResponse, error) {
	return doShellyUpdateRequest(hostname, false)
}
