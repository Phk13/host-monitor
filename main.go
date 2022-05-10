package main

import (
	"flag"
	"fmt"

	"github.com/phk13/host-monitor/pkg/app"
	"github.com/phk13/host-monitor/pkg/notification"

	log "github.com/sirupsen/logrus"
)

type ipFlag []string

func (i *ipFlag) String() string {
	return fmt.Sprintf("%s", *i)
}

func (i *ipFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var ipRangeCIDR ipFlag
	flag.Var(&ipRangeCIDR, "ip", "IP range (CIDR or single IP) to monitor - e.g. 192.168.1.1 or 192.168.1.0/24")

	interval := flag.Int("i", 60, "Time in seconds between each check")
	timeout := flag.Int("t", 1000, "timeout in milliseconds between each check")
	mailAddr := flag.String("mail", "", "Mail to notify")
	debug := flag.Bool("debug", false, "Debug logging")
	mailTest := flag.Bool("mailTest", false, "Do not send any mail, only try to connect and trigger notifications.")

	flag.Parse()

	// Parse IP range
	if len(ipRangeCIDR) == 0 {
		log.Fatal("Specify an IP range to check (-ip)")
	}

	app.SetupLogger(*debug)

	mailSrv := notification.NewMailNotificationService(*mailAddr, *mailTest)

	app.CheckHosts(ipRangeCIDR, *interval, *timeout, mailSrv)

}
