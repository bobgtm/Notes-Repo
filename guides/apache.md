# Apache Installation

```
sudo apt update
```
k
```
sudo apt install apache2
```

```sudo ufw app list
```

```sudo ufw allow `\Apache`\
```
Check the server
```http://server_ip
```
## Creating the Web Root Directory

```
sudo mkdir /var/www/example.com
```

```
sudo chown -R $USER:$USER /var/www/example.com
```

```
sudo chmod -R 755 /var/www/example.com
```

```
sudo vim /var/www/example.com/index.html
```

- Create some html text

## Creating the HTTP Virtual Host
