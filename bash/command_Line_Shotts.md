# Linux Command Line

Terms/Questions: 
- Command History
- Is there a difference between Superuser/root? 
- Who is Steve Bourne? 
- How is the file system organized? How does the gui display file system differntly than the terminal? 
- How are storage devices displayed differently on windows v linux? 
- Directories
    - / 
    - /boot - bootloader files
    - /etc - contains conf files for the system
        - /passwd - user info
        - /fstab - table of devices mounted on system boot
        - /hosts - lists network host names and IP address
        - /init.d - scripts that start system services

    - /sbin | /usr/sbin - programs for system admin for use by superuser
    - /tmp - progams write temporary files
    - /dev - contains devices available to the system (Devices treated like files in Linux/Unix
    - /proc - processes running on the system 

- Symbolic Links - points to another file: when system is given a file name that is a symbolic link, it transparently maps it to the file it is point to. 
- Check out Filesystem Hierarch Standard


- Wildcards
    - * Matches any characters
    - ? Matches any single character
    - 
