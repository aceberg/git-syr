#!/bin/bash

PKGDIR=git-syr-$1-0_all

umask 0022

mkdir -p $PKGDIR/usr/bin
mkdir -p $PKGDIR/lib/systemd/system

cp configs/git-syr.service $PKGDIR/lib/systemd/system/
cp configs/git-syr@.service $PKGDIR/lib/systemd/system/

cp git-syr $PKGDIR/usr/bin/

mkdir -p $PKGDIR/DEBIAN

echo "Package: git-syr
Version: $1
Section: utils
Priority: optional
Architecture: all
Maintainer: aceberg <aceberg_a@proton.me>
Description: Sync Your Repos - pull or push your git repos regularly
" > $PKGDIR/DEBIAN/control

echo "
systemctl daemon-reload
" > $PKGDIR/DEBIAN/postinst

dpkg-deb --build --root-owner-group $PKGDIR

rm -rf $PKGDIR