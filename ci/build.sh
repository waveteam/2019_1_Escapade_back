#!/bin/sh

# enviroment
PROJECT='escapade'
REPO='wavepark'

call() 
{   # invoke the command sucessfuly or terminate
    "$@" && return || echo " -- command: '$@' failed" && exit 1
}

# invoke docker-compose build; -> ${PROJECT}_*
echo " >> starting docker-compose build..."
call sudo docker-compose -p $PROJECT --skip-hostname-check build --parallel
echo " >> docker-compose build was completed!"
