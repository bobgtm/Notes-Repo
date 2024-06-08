# My Man Page
List of commands I commonly use with descriptions given succinctly and in my own words. 

## Files
Stat - show file or file system status

setfacl -m user:root:rwx aclfile.txt - set acl attributes on file
getfacl -t aclfile - get acl atts on file

sudo chattr +i aclfile.txt - set immutable attr on file
lsattr file - get att on file

* Globbing - expanding a nonspecific file name containing wildcard characters in specific filenames 
``` ls file* ```
returns 

``` file ```
``` file.txt ```
``` file.jpg ```
``` file.png ```


``` ls file?.txt ```
* - Matches any character
? - Matches one of any character

Character Classes
[:digit:] - Whole Numbers  = 1, 2, 3 = [0-9]
[:upper:] - Uppercase Letters = A, B, C = [A-Z]
[:lower:] - Lowercase Letters
[:alpha:] - Upper and Loewrcase
[:alnum:] - Upper, lower, and numbers
[:space:]
[:punct:] - Punctuation


Extended Glob Syntax: 

Function            | Extended Glob             | Regular Expression
--------------------------------------------------------------------
                    |                           |
Group Characters    | (abc)                     | (abc)
--------------------|---------------------------|-------------------
                    |                           |                   
Mutliple Patterns   | (abc|def)                 | (abc|def)
--------------------|---------------------------|------------------
                    |                           |
Invert Match        | !(abc)                    |
                    | ^(abc)                    |
--------------------|---------------------------|------------------



Specifying Occurrences: 

Occurrences         | Extended Glob             | Regular Expression
--------------------------------------------------------------------
                    |                           |
0 or 1              | ?(abc)                    | (abc)?
--------------------|---------------------------|-------------------
                    |                           |                   
Exactly 1           | @(abc)                    | (abc)
--------------------|---------------------------|------------------
                    |                           |
1 or more           | +(abc)                    | (abc)+
--------------------|---------------------------|------------------
                    |                           |
0 or more           | *(abc)                    | (abc)*
--------------------|---------------------------|------------------


Permissions:
fff

Special Bits: 
* Set user ID (SUID) - Run as user owner of the file
--> Has no affect on directoris
* Set group ID (SGID) - Run as group owner of the file
--> provides group inheritance
* Sticky - Remain in swap but not functional in linux on files as linux caches stuff. 


## Gathering System Information



