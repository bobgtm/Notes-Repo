Week 2 Lecture Notes - ACLs & SELinux

## Notes

### Diff b/w engineering troubleshooting and admin troubleshooting

In admin, something worked at one point. Engineers do initial config for the first time. 

Likely some doc to go to from teh build or runtime. 
Compare one system to another working system for differences. 

### Common Administartion Tasks

- Powering on and off servers
```
shutdown -h now
``` 
```
shutdown -r now
```


- Maintaining inventory in a system
- - Adding Changing, and removing inventory

- Installing and updating software
- - - Sometimes giving sudo access to some teams
- - Configuring repositories
- - Configuring software to work within the cluster

### Networking - Tools of the trade

- ip addr, ip -br addr, ip route, ip neigh
- ping
- arp
- nslookup, host, dig
- ss or netstat
- nc r telnet

### The USE Method

- For every Resource Check
1. Utilization of the system resources
2. Saturation - when a queue length and time begins to form
3. Errors - errors reported in logs
