#!/bin/bash

git clone https://github.com/tomaspinho/rtl8821ce.git
cd rtl8821ce

./dkms-install.sh

cd ..

rm -rf rtl8821ce
