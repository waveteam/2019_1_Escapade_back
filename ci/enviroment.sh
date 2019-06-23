#!/bin/sh

# remove a previous version
echo " >> removeing a previous docker-compose version..."
sudo rm /usr/local/bin/docker-compose
pip3 uninstall docker-compose
echo " >> removeing was completted!"

# install a docker-compose @see https://docs.docker.com/compose/install/
echo " >> installing a new docker-compose version..."
sudo curl -L https://github.com/docker/compose/releases/download/1.24.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
pip3 install docker-compose
echo " >> installation was completted!"
