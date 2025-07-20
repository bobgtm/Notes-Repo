# Notes: 

## Definitions/Terminology

1. Kernel - software that sits between the hardware and userspace (such as the apps we'd run on our machine). It is responsible for: 
- Process Management
- Memory management
- Networking
- Filesystems
- Devices
    1. uname -r displays the Kernel Version
    2. uname -a

2. Kernel Params: used to overwrite default values and set specific hardware settings
    - Defined in the boot entry configuration file for each boot entry
    - What is a boot entry? 
        - Collection of options stored in config file and tied to a particular kernel version
3. OS Version: 
	- uname -o
4. Modules: 
	- modprobe: looks through the kernel command line and collects module paramenters when it loads a module
	- modinfo
5. Mount Points: 
	- lsblk
6. Text Editor:
