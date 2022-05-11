# Host monitor
ICMP monitoring and notification tool (Gmail OAuth)


# Usage
./host-monitor -ip ip/range -mail example@gmail.com [-i interval] [-t timeout] [-debug] [-mailTest]
### Options
```
  -debug
        Debug logging
  -i int
        Interval in seconds between each check (default 60)
  -ip value
        IP range (CIDR or single IP) to monitor - e.g. 192.168.1.1 or 192.168.1.0/24 (can be repeated to specify multiple ip/ranges)
  -mail string
        Mail to notify
  -mailTest
        Do not send any mail, only try to connect and trigger notifications.
  -t int
        Timeout in milliseconds for ICMP (default 1000ms)
```
