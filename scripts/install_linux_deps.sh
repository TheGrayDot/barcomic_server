#!/bin/bash


# Packages for compiling robotgo on linux
package_list_robotgo=(
    libx11-dev
    xorg-dev
    libxtst-dev
    xcb
    libxcb-xkb-dev
    x11-xkb-utils
    libx11-xcb-dev
    libxkbcommon-x11-dev
    libxkbcommon-dev
    # Windows CC packages
    gcc-multilib
    gcc-mingw-w64
    libz-mingw-w64-dev
)

sudo apt-get update
sudo apt-get install -y "${package_list_robotgo[@]}"
