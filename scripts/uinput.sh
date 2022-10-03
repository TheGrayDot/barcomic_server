#!/bin/bash


# Check running as root (UID = 0)
if [ "$EUID" -ne 0 ]; then 
    echo "[*] Error: Please run as root. Exiting."
    exit 1
fi

# Make uinput group, and add user
GROUP_NAME="uinput"
if [ ! $(getent group $GROUP_NAME) ]; then
    echo "[*] Creating $GROUP_NAME..."
    sudo groupadd $GROUP_NAME
    usermod -aG $GROUP_NAME $USER
fi

echo "[*] Configuring uinput..."
sudo udevadm control --reload-rules
echo "SUBSYSTEM==\"misc\", KERNEL==\"uinput\", GROUP=\"uinput\", MODE=\"0660\"" | tee /etc/udev/rules.d/uinput.rules
echo "uinput" | tee /etc/modules-load.d/uinput.conf
chmod +0666 /dev/uinput

# Leverage switch user to reload group membership
exec su -l $USER
