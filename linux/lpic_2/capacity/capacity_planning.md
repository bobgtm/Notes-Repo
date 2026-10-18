# Capacity planning

Determining whether a server has enough power to meet certain requirements

Big 4 of performance:

CPU - Brain doing much of the computational work
Memory (RAM) - Store and fetch data as quickly as possible
Disk Storage - Disk I/O 
Network I/O - amount of data that can be sent into and out of the system

DB are ram and disk i/o intensive

Web Server - active on ram and cpu 

Useful tools 
- Top
- sysstat
- sar -u = cpu perf
- sar -r = mem perf
- sar -b = disk i/o
- sar -n = networking
- uptime
