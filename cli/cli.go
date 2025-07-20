package cli

import (
	flag "github.com/spf13/pflag"
	"log"
	"time"
)

var (
	ScanTimeout = time.Second * 60
	ShowVersion = flag.BoolP("version", "v", false, "Show version information")
	Username    = flag.String("Username", "admin", "Username to use for the authentication. Keep it to 'admin' (default)")
	Password    = flag.String("Password", "", "Password to use for authentication (optional)")
	Channel     = flag.String("channel", "stable", "Stage to update to: stable, beta")
	GenToUpdate = flag.String("gen", "all", "Generation to update. One of this values: all,gen1,gen2")
	Hosts       = flag.StringSlice("host", []string{}, "Use host/IP address(es) instead of device discovery (can be specified multiple times or be comma-separated)")
)

func ParseArgs() {
	flag.Parse()

	if *Channel != "stable" && *Channel != "beta" {
		flag.Usage()
		log.Fatal("Unknown update stage")
	}

	switch *GenToUpdate {
	case "all", "gen1", "gen2", "gen3":
	default:
		flag.Usage()
		log.Fatal("Unknown generation")
	}
}
