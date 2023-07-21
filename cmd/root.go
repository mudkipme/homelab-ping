package cmd

import (
	"os"

	"github.com/mudkipme/homelab-ping/config"
	"github.com/mudkipme/homelab-ping/internal/ping"
	"github.com/spf13/cobra"
)

var c config.Config

var rootCmd = &cobra.Command{
	Use:   "homelab-ping",
	Short: "Ping an address periodically and restart the computer when ping fails",
	Run: func(cmd *cobra.Command, args []string) {
		hp := ping.New(&c)
		hp.Start()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&c.Address, "address", "192.168.1.1", "the router address to ping")
	rootCmd.PersistentFlags().IntVar(&c.PingCount, "ping-count", 5, "the number of pings to send")
	rootCmd.PersistentFlags().IntVar(&c.PingInterval, "ping-interval", 1, "the interval between pings in minutes")
	rootCmd.PersistentFlags().IntVar(&c.RestartInterval, "restart-interval", 60, "the interval between restarts in minutes")
	rootCmd.PersistentFlags().IntVar(&c.FailTimes, "fail-times", 5, "the number of attampts to fail before restarting")
}
