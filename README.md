homelab-ping
============

## Deprecation

This project is no longer needed as this is a known issue with Intel i129-V NIC, which can be fixed by disabling segment offloading.

```bash
ethtool -K <interface> tso off gso off
```

---

This is a small project targeting certain issues related to my Homelab computer.

Specifically, the Intel i219-V Ethernet controller has an issue that causes the PC to randomly lose the network card or connectivity after a few weeks. I haven't found a reliable solution and am unable to restart it when I'm not at home.

This program can periodically ping the router and will restart the PC if it loses connection after multiple attempts.

**This program may unexpectedly restart your computer and may cause data loss.**

## Installation

This program only supports Linux and requires the `root` user.

```bash
go install github.com/mudkipme/homelab-ping
```

## Usage

```bash
homelab-ping [flags]

Flags:
      --address string         the router address to ping (default "192.168.1.1")
      --fail-times int         the number of attampts to fail before restarting (default 5)
      --ping-count int         the number of pings to send (default 5)
      --ping-interval int      the interval between pings in minutes (default 1)
      --restart-interval int   the interval between restarts in minutes (default 60)
```

### Use with systemd

Put the `homelab-ping` binary in `/usr/local/bin` and create `homelab-ping.service` in `/etc/systemd/system` folder.

```ini
[Unit]
Description=homelab-ping
After=network.target

[Service]
ExecStart=/usr/local/bin/homelab-ping
Restart=on-abort
RemainAfterExit=yes
RestartSec=300s
TimeoutSec=300s

[Install]
WantedBy=multi-user.target
```

Enable unit for automatic start:

```bash
systemctl enable --now homelab-ping.service
```

## License

[MIT License](LICENSE)