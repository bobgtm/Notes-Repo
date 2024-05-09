# Basics

## Pipes
 send the output of one process to another as they were connected 

`cat loremipsum.txt | wc -l`

## Redirections 
send streams to or from files 

`ls > list.txt`

Stream		Name			Content
0		Standard Inupt		Keyboard or other input

1		Standard Output		Regular Output

2		Standard Error		Output marked as error


Symbol		Function

> 		Output redirection

>>		Append redirection

< 		Input redirection

<<		Here Document - usefull for displaying options or long passages


## Commands and Built-ins

Many programs used in Bash are commands 
- Thes do not depend on Bash

Bash includes built-in commands which are part of Bash itself

Can specify which to use by prefacing with `command echo` or `builtin echo`

command -V will show us if a command is builtin or wrogram


