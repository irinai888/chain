#!/bin/bash
# mkdir ~/faucet
# cd ~/faucet
# curl -O https://raw.githubusercontent.com/lcnem/jpyx/main/faucet/deploy.sh
docker-compose down
curl -O https://raw.githubusercontent.com/lcnem/jpyx/main/faucet/docker-compose.yml
curl -O https://raw.githubusercontent.com/lcnem/jpyx/main/faucet/nginx.conf
docker cp $(docker ps -qf "name=jpyxd"):/usr/bin/jpyxd ~/faucet/jpyxd
docker-compose up -d