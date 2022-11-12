#!/bin/bash


source scripts/export.sh

echo "[*] Current version: $PROJECT_VERSION"

read -p "[*] Tag and Release? (y/N) " yn
case $yn in
	y ) git tag $PROJECT_VERSION;
        git push --tags;
        exit 0;;
	n ) echo "[*] Exiting...";
		exit 0;;
	* ) echo "[*] Invalid response... Exiting";
        exit 1;;
esac
