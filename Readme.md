
```
sudo curl -L "https://github.com/skabashnyuk/cli/releases/download/$(curl -L -s -H 'Accept: application/json' https://github.com/skabashnyuk/cli/releases/latest |  sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')/$(echo che_`uname -s`_`uname -m`)" -o /usr/local/bin/che | sudo  chmod +x /usr/local/bin/che
```