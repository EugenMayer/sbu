package update

import (
	"context"
	"fmt"
	"github.com/fermayo/shelly-bulk-update/cli"
	"github.com/grandcat/zeroconf"
	"log"
	"strings"
	"sync"
)

func AutoDiscoverUsingAndUpdate() {
	// Listen only for IPv4 addresses. Otherwise, it may happen that ServiceEntry has an empty
	// AddrIPv4 slice. It happens when the IPv6 arrives first and ServiceEntries are not updated
	// when more data arrives.
	// See https://github.com/grandcat/zeroconf/issues/27
	resolver, err := zeroconf.NewResolver(zeroconf.SelectIPTraffic(zeroconf.IPv4))
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	var wg sync.WaitGroup
	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		fmt.Printf("[scanner] looking for Shelly devices using mDNS (%ds timeout)...\n", int(cli.ScanTimeout.Seconds()))
		for entry := range results {
			entry := entry
			if strings.HasPrefix(strings.ToLower(entry.Instance), "shelly") {
				wg.Add(1)
				go func() {
					address := entry.HostName
					if len(entry.AddrIPv4) > 0 {
						address = entry.AddrIPv4[0].String()
						// IPv6 support is still very limited
						// See https://shelly-api-docs.shelly.cloud/gen2/General/IPv6
						//} else if len(entry.AddrIPv6) > 0 {
						//	address = fmt.Sprintf("[%s]", entry.AddrIPv6[0].String())
					}
					UpdateShelly(entry.Instance, address, entry.Text, *cli.GenToUpdate)
					wg.Done()
				}()
			}
		}
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), cli.ScanTimeout)
	defer cancel()
	err = resolver.Browse(ctx, "_http._tcp", "local.", entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	fmt.Println("[scanner] scanning process finished")
	wg.Wait()
}
