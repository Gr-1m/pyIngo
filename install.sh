#!/bin/bash

if [ $(id -u) -ne 0 ]
then
	echo 'You are not root, your permissions may not be sufficient'
	exit 1
fi

rm -f "$(go env GOROOT)/src/goInpy"
rm -f "$(go env GOROOT)/src/pyIngo"
ln -s $(pwd) "$(go env GOROOT)/src/pyIngo";

if [ $? -eq 0 ]
then
    echo -e "\033[32;1m[+]\033[0m OK "
else
    echo -e "\033[31;1m[-]\033[0m Error"
fi

