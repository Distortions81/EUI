#!/bin/sh
set -e
sudo apt-get update
sudo apt-get install -y libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libxxf86vm-dev libgl1-mesa-dev brotli
sudo apt-get install -y build-essential
sudo apt-get install -y libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config
