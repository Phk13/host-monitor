package app

import (
	"sync"
	"time"

	"github.com/phk13/host-monitor/pkg/domain"
	log "github.com/sirupsen/logrus"
)

func Monitor(host *domain.Host, pingLock *sync.Mutex, wg *sync.WaitGroup, timeout, monitorInterval int, notificationSrv domain.Notification) {
	defer wg.Done()
	// If host is down in first try, stop monitoring
	if !host.Status {
		log.Infof("IP %s is unreachable", host.IP)
		return
	}
	log.Infof("Host %s is up", host.IP)

	var sentMailTime time.Time

	for {
		prevStatus := host.Status
		// Acquire lock to ping (1 simultaneous ping)
		pingLock.Lock()
		host.CheckStatus(timeout)
		pingLock.Unlock()

		if !host.Status && prevStatus {
			log.Infof("Host %s is down, sending notification...", host.IP)
			sentMailTime = time.Now()
			notificationSrv.SendNotification(host.IP, sentMailTime)

		} else {
			log.Infof("Host %s is up", host.IP)
		}
		time.Sleep(time.Second * time.Duration(monitorInterval))
	}
}
