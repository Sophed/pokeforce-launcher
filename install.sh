#!/bin/bash

if [ "$EUID" -ne 0 ]
  then echo "please run this script with sudo"
  exit 1
fi

wget https://github.com/Sophed/pokeforce-launcher/releases/download/v0.1/launcher-x86
chmod +x launcher-x86
sudo mv launcher-x86 /usr/bin/pokeforce
curl -o pokeforce.desktop https://raw.githubusercontent.com/Sophed/pokeforce-launcher/main/pokeforce.desktop
sudo mv pokeforce.desktop /usr/share/applications/
echo "successfully installed! run 'pokeforce' to start."