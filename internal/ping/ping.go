package ping

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/mackerelio/go-osstat/uptime"
	"github.com/mudkipme/homelab-ping/config"
	probing "github.com/prometheus-community/pro-bing"
)

type HomelabPing struct {
	config      *config.Config
	currentFail int
}

func New(config *config.Config) *HomelabPing {
	return &HomelabPing{
		config: config,
	}
}

func (p *HomelabPing) Start() {
	ticker := time.NewTicker(time.Minute * time.Duration(p.config.PingInterval))
	defer ticker.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

loop:
	for {
		select {
		case <-ticker.C:
			if p.try() {
				break loop
			}
		case <-c:
			break loop
		}
	}
}

func (p *HomelabPing) try() bool {
	err := p.ping()
	if err != nil {
		log.Printf("ping failed: %v\n", err)
		p.currentFail++
		if p.currentFail >= p.config.FailTimes {
			return p.restart()
		}
	} else {
		p.currentFail = 0
	}
	return false
}

func (p *HomelabPing) restart() bool {
	p.currentFail = 0
	uptime, err := uptime.Get()
	if err != nil {
		log.Print(err.Error())
		return false
	}

	if uptime < time.Minute*time.Duration(p.config.RestartInterval) {
		return false
	}

	log.Print("restarting")
	reboot()
	return true
}

func (p *HomelabPing) ping() error {
	pinger, err := probing.NewPinger(p.config.Address)
	if err != nil {
		return err
	}

	// Unprivileged ping are not enabled by default on Proxmox
	pinger.SetPrivileged(true)
	pinger.Count = p.config.PingCount
	pinger.Timeout = time.Second * 20
	err = pinger.Run()
	if err != nil {
		return err
	}
	stat := pinger.Statistics()
	if stat.PacketsRecv > 0 {
		return nil
	}
	return errors.New("no packets received")
}
