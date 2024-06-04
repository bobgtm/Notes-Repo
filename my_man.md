# My Man Page
List of commands I commonly use with descriptions given succinctly and in my own words. 

## Files
Stat - show file or file system status

setfacl -m user:root:rwx aclfile.txt - set acl attributes on file
getfacl -t aclfile - get acl atts on file

sudo chattr +i aclfile.txt - set immutable attr on file
lsattr file - get att on file
