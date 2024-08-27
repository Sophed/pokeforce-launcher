#!/bin/bash

if [ "$EUID" -ne 0 ]
  then echo "please run this script with sudo"
  exit 1
fi

curl -o pokeforce https://github.com/Sophed/pokeforce-launcher/releases/download/v0.1/launcher-x86
chmod +x pokeforce
sudo mv pokeforce /usr/bin/
echo "successfully installed! run 'pokeforce' to start."