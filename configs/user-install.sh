#!/bin/sh

umask 0022

mkdir -p $HOME/.config/git-syr/

sudo cp git-syr /usr/bin/
sudo cp git-syr@.service /lib/systemd/system/