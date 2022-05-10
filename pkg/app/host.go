package app

import (
	"net/netip"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/phk13/host-monitor/pkg/domain"
)

func CheckHosts(ipRangeCIDR []string, interval, timeout int, notificationSrv domain.Notification) {
	var ipRange []netip.Addr
	var err error
	ipRange, err = ParseRanges(ipRangeCIDR)
	if err != nil {
		log.Fatal("Invalid IP range -> ", err)
	}

	log.Infof("Monitoring %d host(s) in range %s", len(ipRange), ipRangeCIDR)
	log.Info("Launching initial scan...")
	var hostList []*domain.Host
	for _, ip := range ipRange {
		host := domain.NewHost(ip.String(), timeout)
		if host.Status {
			hostList = append(hostList, host)
		}
	}
	wg := sync.WaitGroup{}
	pingLock := sync.Mutex{}
	for _, host := range hostList {
		wg.Add(1)
		go Monitor(host, &pingLock, &wg, timeout, interval, notificationSrv)
	}
	log.Infof("%d monitors launched", len(hostList))
	wg.Wait()
	log.Info("No IPs being monitored. Exiting...")
}
