package utils

import (
	"net"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func Ping(target string, timeout int) bool {
	log.Debugf("Pinging %s...", target)
	ip, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		log.Debug(err)
		return false
	}
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		log.Debugf("Error on ListenPacket %v", err)
		return false
	}
	defer conn.Close()

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(""),
		},
	}
	msg_bytes, err := msg.Marshal(nil)
	if err != nil {
		log.Debugf("Error on Marshal %s - %v", target, msg_bytes)
		return false
	}

	// Write the message to the listening connection
	if _, err := conn.WriteTo(msg_bytes, &net.UDPAddr{IP: net.ParseIP(ip.String())}); err != nil {
		log.Debugf("Error on WriteTo %s - %v", target, err)
		return false
	}

	err = conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(timeout)))
	if err != nil {
		log.Debugf("Error on SetReadDeadline %s - %v", target, err)
		return false
	}
	reply := make([]byte, 1500)
	n, _, err := conn.ReadFrom(reply)

	if err != nil {
		log.Debugf("Error on ReadFrom %s - %v", target, err)
		return false
	}
	parsed_reply, err := icmp.ParseMessage(1, reply[:n])

	if err != nil {
		log.Debugf("Error on ParseMessage %s - %v", target, err)
		return false
	}
	switch parsed_reply.Code {
	case 0:
		log.Debugf("Pinged %s -> true", target)
		return true
	default:
		log.Debugf("Pinged %s -> false", target)
		return false
	}
}
