package update

import (
	"github.com/fermayo/shelly-bulk-update/gen1"
	"github.com/fermayo/shelly-bulk-update/gen2plus"
	"slices"
)

func UpdateShelly(name, address string, txtRecords []string, genToUpdate string) {
	if slices.Contains(txtRecords, "gen=2") || slices.Contains(txtRecords, "gen=3") {
		if genToUpdate == "gen2" || genToUpdate == "gen3" || genToUpdate == "all" {
			gen2plus.UpdateShelly(name, address)
		}
		return
	}

	if genToUpdate == "gen1" || genToUpdate == "all" {
		gen1.UpdateShelly(name, address)
	}
}
