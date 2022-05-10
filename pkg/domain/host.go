package domain

import "github.com/phk13/host-monitor/pkg/utils"

type Host struct {
	IP     string
	Status bool
}

func (h *Host) CheckStatus(timeout int) bool {
	h.Status = utils.Ping(h.IP, timeout)
	return h.Status
}

func NewHost(ip string, timeout int) *Host {
	host := &Host{
		IP: ip,
	}
	host.Status = host.CheckStatus(timeout)
	return host
}
