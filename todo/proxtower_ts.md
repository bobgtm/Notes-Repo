# Troubleshooting Proxmox Tower Network Connection

## Issue

When powering on, local proxmox host is able to connect across the home network.
However, after a period of brief connectivity, the host becomes unreachable despite
remaining powered on with link light activity showing on ports on multiple home
network switches. 

### Troubleshooting Thus Far

The issue persists after attempting the following hardware based troubleshooting:
- Changing network cables
- Changing switch ports

The issue persists after attempting the following software based troubleshooting:
- Installing a different OS on the hardware (issue does not arise when using machine
  as a stadalone system instead). 

### Further Troubleshooting

#### CMD Results

- ifconfig
    1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group
    default qlen 1000
        link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00 
        inet 127.0.0.1/8 scope host lo
            valid_lft forever preferred_lft forever
        inet6 ::1/128 scope host noprefixroute
            valid_lft forever preferred_lft forever
    2: eno1 <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_code1 master vmbr0 state UP group default qlen 1000
        inet 127.0.0.1/8 scope host lo
            valid_lft forever preferred_lft forever
        inet6 ::1/128 scope host noprefixroute
            valid_lft forever preferred_lft forever
            altname enp0s25
            altname enx6451062af551
