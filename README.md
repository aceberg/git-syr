[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/git-syr)](https://goreportcard.com/report/github.com/aceberg/git-syr)
[![Binary-release](https://github.com/aceberg/git-syr/actions/workflows/release.yml/badge.svg)](https://github.com/aceberg/git-syr/actions/workflows/release.yml)
![GitHub](https://img.shields.io/github/license/aceberg/git-syr)

<h1><a href="https://github.com/aceberg/git-syr">
    <img src="https://raw.githubusercontent.com/aceberg/git-syr/main/assets/logo.png" width="20" />
</a>git-syr</h1>
<br/>

Sync Your Repos - pull or push your git repos regularly. For dotfiles backups or note taking in git repo

- [CLI and GUI](https://github.com/aceberg/git-syr#cli-and-gui) 
- [Installation](https://github.com/aceberg/git-syr#installation)   
- [Usage](https://github.com/aceberg/git-syr#usage)   
- [Config](https://github.com/aceberg/git-syr#config)   
- [Options](https://github.com/aceberg/git-syr#options)  
- [Thanks](https://github.com/aceberg/git-syr#thanks) 

![screenshot](https://raw.githubusercontent.com/aceberg/git-syr/main/assets/Screenshot%202022-12-22%20at%2012-51-48%20Sync%20Your%20Repos.png)


## CLI and GUI
There are two packages available: `git-syr` and `git-syr-cli`. For command line only installation you can use `git-syr-cli`, just replace the name in the instructions below. Though `git-syr` is capable of both CLI and GUI modes.

## Installation

### 1. From .deb repository (recommended)
```sh
curl -s --compressed "https://aceberg.github.io/ppa/KEY.gpg" | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/aceberg.gpg
```
```sh
sudo curl -s --compressed -o /etc/apt/sources.list.d/aceberg.list "https://aceberg.github.io/ppa/aceberg.list"
```
```sh
sudo apt update && sudo apt install git-syr
```
### 2. From .deb file
Download [latest](https://github.com/aceberg/git-syr/releases/latest) release, install with your package maneger

### 3. From .tar.gz
Download [latest](https://github.com/aceberg/git-syr/releases/latest) release, then
```sh
tar xvzf git-syr-*.tar.gz
cd git-syr
sudo ./install.sh
```



## Usage
### 1. Systemd as user (recommended)
Enable and start service, replace `MYUSER` with your username
```sh
systemctl enable git-syr@MYUSER.service
systemctl start git-syr@MYUSER.service
```
Web GUI will be available at [http://0.0.0.0:8844](http://0.0.0.0:8844)

### 2. Systemd as root
Enable and start service
```sh
systemctl enable git-syr.service
systemctl start git-syr.service
```
Web GUI will be available at [http://0.0.0.0:8844](http://0.0.0.0:8844)

### 3. From command line
Just run `git-syr`. Be mindful of the config files paths listed in [options](https://github.com/aceberg/git-syr#options) section.


## Config
### 1. With web GUI
You can do all configuration through web interface. Config files paths are listed in [options](https://github.com/aceberg/git-syr#options) section below.

### 2. CLI
`repos.yaml` example:
```yaml
- name: Dotfiles
  path: /home/data/repo/dotfiles
  timeout: 4h
  pull: "no"
  push: "yes"

- name: MyNotes
  path: /home/data/repo/MyNotes
  timeout: 1m
  pull: "yes"
  push: "yes"
```
`config.yaml` example:
```yaml
host: 0.0.0.0
port: "8844"
theme: cosmo
```



## Options
### 1. git-syr
| Key  | Description | Default | Systemd (user) | Systemd (root) |
| --------  | ----------- | ------- | --- | --- |
| -c | Path to GUI config file |./config.yaml| $HOME/.config/git-syr/config.yaml | /etc/git-syr/config.yaml |
| -l | Path to log file | ./git-syr.log | $HOME/.config/git-syr/git-syr.log | /var/log/git-syr.log |
| -r | Path to repos yaml file |./repos.yaml| $HOME/.config/git-syr/repos.yaml | /etc/git-syr/repos.yaml |
| -w | Launch with web gui | true | | |

### 2. git-syr-cli
| Key  | Description | Default | Systemd (user) | Systemd (root) |
| --------  | ----------- | ------- | --- | --- |
| -l | Path to log file | "" (No log file) | $HOME/.config/git-syr/git-syr.log | /var/log/git-syr.log |
| -r | Path to repos yaml file |./repos.yaml| $HOME/.config/git-syr/repos.yaml | /etc/git-syr/repos.yaml |

To run git-syr without log file pass empty string `-l ""`



## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/git-syr/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- <a href="https://www.flaticon.com/free-icons/cheese" title="cheese icons">Cheese icons created by Freepik - Flaticon</a>