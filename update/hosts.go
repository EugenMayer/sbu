package update

import (
	"fmt"
	"github.com/fermayo/shelly-bulk-update/cli"
	"github.com/fermayo/shelly-bulk-update/gen1"
	"log"
)

func SpecificHosts(hosts []string) {
	fmt.Println("[host-updates] Updating only specific hosts")

	for _, host := range hosts {
		fmt.Printf("Updating host %s\n", host)

		info, err := gen1.GetShellyInfo(host)
		if err != nil {
			log.Fatalln("Failed to initialize resolver:", err.Error())
		}
		txtRecords := []string{fmt.Sprintf("gen=%d", info.Generation)}
		UpdateShelly(info.Name, host, txtRecords, *cli.GenToUpdate)
	}
}
