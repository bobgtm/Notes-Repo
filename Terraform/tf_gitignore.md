# Important includes for Terraform gitignore

- Important to to commit certain files to version control in order to avoid sharing
  secrets, machine state, 

- terraform.tfstate / terraform.tfstate.* backup state files
-- These contain secrets and other sensitive information

- .terraform.tfstate.lock.info 
-- File is created and deleted automatically by TF when running ```terraform apply```

-- Contains info about state lock

- ```.terraform directory`` - where TF downloads providers and child modules

- saved plan files when running ```terraform plan -out```
