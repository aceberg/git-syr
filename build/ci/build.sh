#!/bin/bash

cd cmd/git-syr/ && CGO_ENABLED=0 go build -o ../../git-syr .
cd ../../

PKGDIR=/opt/git-syr

umask 0022

mkdir -p $PKGDIR
cp git-syr $PKGDIR/
cp configs/git-syr.service $PKGDIR/
cp configs/git-syr@.service $PKGDIR/
cp configs/user-install.sh $PKGDIR/

cd /opt
tar cvzf git-syr-$1.tar.gz git-syr
cd -
cp /opt/git-syr-$1.tar.gz .