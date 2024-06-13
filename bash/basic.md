#Basics

## Pipes
 send the output of one process to another as they were connected 
`cat loremipsum.txt | wc -l`

## Redirections 
Send streams to or from files 

`ls > list.txt`

Stream		Name			Content
0		Standard Inupt		Keyboard or other input

1		Standard Output		Regular Output

2		Standard Error		Output marked as error


Symbol		Function

`>`    	Output redirection

`>>`	    Append redirection

`<` 		Input redirection

`<<`    	Here Document - usefull for displaying options or long passages


## Commands and Built-ins

1. Many programs used in Bash are commands 
    A. These commands do not depend on Bash itself and can be used as a part of the command line interface more generally. 

2. Bash includes built-in commands which are part of Bash itself

3. Can specify which to use by prefacing with `command echo` or `builtin echo`

4. Command -V will show us if a command is a bash builtin or program


