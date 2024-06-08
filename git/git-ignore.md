
### Excluding local files without creating a .gitignore file

If you don't want to create a .gitignore file to share with others, you can create rules that are not committed with the repository. You can use this technique for locally-generated files that you don't expect other users to generate, such as files created by your editor.

Use your favorite text editor to open the file called .git/info/exclude within the root of your Git repository. Any rule you add here will not be checked in, and will only ignore files for your local repository.

* Open Git Bash.
* Navigate to the location of your Git repository.
* Using your favorite text editor, open the file .git/info/exclude.

