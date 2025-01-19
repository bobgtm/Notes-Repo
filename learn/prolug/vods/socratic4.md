# Questions and Commands

1. For all users, can you list a count of what default shells they have? 

Can you verify an indidvidual user's limits of open files? Know where to
change this? 

```
ulimit -a -u scott
```

```
limits.conf or limits.d 
``` 

Can you tell how system was booted by grub? 

```
cat /proc/cmdline
```
Want to be able to look at certain things to make sure no changes are
unexpected. 

Can you show everything connected to PCI bus? 

```
lspci 

dmesg | grep -i pci
```

Tell me the running kernel version? 

```
uname -r
```

Show that the sshd service is running? 

``` 
ps -aux | grep -i service
```
``` 
ps -ef | grep 
```
``` systemctl | grep -i prom 
```
``` ss -ntulp | grep service 
```


Tell me version of Linux you're using:

``` 
lsb_release -a
cat /etc/os-release | grep -i version
cat /etc/*release
```

Set a process to run every 5 minutes on the server? 

```
cat /etc/crontab
```

