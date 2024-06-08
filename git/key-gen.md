Source - https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent


1. ```ssh-keygen -t ed25519 -c "email"```

2. Enter a file in which to save the key (/home/YOU/.ssh/id_ALGORITHM):[Press enter]

3. Start the ssh-agent in the background.

$ eval "$(ssh-agent -s)"
> Agent pid 59566

4. Add your SSH private key to the ssh-agent.

If you created your key with a different name, or if you are adding an existing key that has a different name, replace id_ed25519 in the command with the name of your private key file.

ssh-add ~/.ssh/id_ed25519



Click New SSH key or Add SSH key.

In the "Title" field, add a descriptive label for the new key. For example, if you're using a personal laptop, you might call this key "Personal laptop".

Select the type of key, either authentication or signing. For more information about commit signing, see "About commit signature verification."

In the "Key" field, paste your public key.

Click Add SSH key.

If prompted, confirm access to your account on GitHub.
