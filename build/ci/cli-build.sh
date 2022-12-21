#!/bin/bash

cd cmd/git-syr-cli/ && CGO_ENABLED=0 go build -o ../../git-syr-cli .
cd ../../

PKGDIR=/opt/git-syr-cli

umask 0022

mkdir -p $PKGDIR
cp git-syr-cli $PKGDIR/
cp configs/git-syr-cli.service $PKGDIR/
cp configs/git-syr-cli@.service $PKGDIR/
cp configs/cli-install.sh $PKGDIR/

cd /opt
tar cvzf git-syr-cli-$1.tar.gz git-syr-cli
cd -
cp /opt/git-syr-cli-$1.tar.gz .