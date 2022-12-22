[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/git-syr)](https://goreportcard.com/report/github.com/aceberg/git-syr)
[![Binary-release](https://github.com/aceberg/git-syr/actions/workflows/release.yml/badge.svg)](https://github.com/aceberg/git-syr/actions/workflows/release.yml)
![GitHub](https://img.shields.io/github/license/aceberg/git-syr)

# git-syr

Sync Your Repos - pull or push your git repos regularly. For dotfiles backups or note taking in git repo

[1. Installation](https://github.com/aceberg/git-syr#1-installation)   
[2. Usage](https://github.com/aceberg/git-syr#2-usage)   
[3. Config](https://github.com/aceberg/git-syr#3-config)   
[4. Options](https://github.com/aceberg/git-syr#4-options)  

SCREENSHOT




## 1. Installation

### From .deb repository (recommended)
```sh
curl -s --compressed "https://aceberg.github.io/ppa/KEY.gpg" | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/aceberg.gpg
```
```sh
sudo curl -s --compressed -o /etc/apt/sources.list.d/aceberg.list "https://aceberg.github.io/ppa/aceberg.list"
```
```sh
sudo apt update && sudo apt install git-syr
```
### From .deb file
Download [latest](https://github.com/aceberg/git-syr/releases/latest) release, install with your package maneger

### From .tar.gz
Download [latest](https://github.com/aceberg/git-syr/releases/latest) release, then
```sh
tar xvzf git-syr-*.tar.gz
cd git-syr
sudo ./install.sh
```



## 2. Usage
### As user (recommended)
#### 1. With systemd
Enable and start service, replace `MYUSER` with your username
```sh
systemctl enable git-syr@MYUSER.service
systemctl start git-syr@MYUSER.service
```
Web GUI will be available at [http://0.0.0.0:8844](http://0.0.0.0:8844)





## 3. Config





## 4. Options
| Key  | Description | Default |
| --------  | ----------- | ------- |
| -c | Path to config yaml file |config.yaml|







## 5. Thanks




<a href="https://www.flaticon.com/free-icons/cheese" title="cheese icons">Cheese icons created by Freepik - Flaticon</a>