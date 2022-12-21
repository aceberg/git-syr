#!/bin/bash

PKGDIR=git-syr-cli-$1-0_all

umask 0022

mkdir -p $PKGDIR/usr/bin
mkdir -p $PKGDIR/lib/systemd/system

cp configs/git-syr-cli.service $PKGDIR/lib/systemd/system/
cp configs/git-syr-cli@.service $PKGDIR/lib/systemd/system/

cp git-syr-cli $PKGDIR/usr/bin/

mkdir -p $PKGDIR/DEBIAN

echo "Package: git-syr-cli
Version: $1
Section: utils
Priority: optional
Architecture: all
Maintainer: aceberg <aceberg_a@proton.me>
Description: Sync Your Repos - pull or push your git repos regularly
" > $PKGDIR/DEBIAN/control

dpkg-deb --build --root-owner-group $PKGDIR

rm -rf $PKGDIR