#!/bin/sh

# enviroment
PROJECT='escapade'
REPO='wavepark'

call() 
{   # invoke the command sucessfuly or terminate
    "$@" && return || echo " -- command: '$@' failed" && exit 1
}

push()
{   # tag the images and push to the docker hub
    NAME="$1"
    IMAGE="${PROJECT}_${NAME}"
    BUILD="${REPO}/${PROJECT}-${NAME}:latest"
    echo " >> assigning an image '${IMAGE}' to '${BUILD}'..."
    call docker tag ${IMAGE} ${BUILD}
    echo " >> pushing an image '${BUILD}'..."
    call docker push ${BUILD}
    echo " >> image '${BUILD}' was processed!"
}

# push the images asyncronously
for NAME in         \
    auth            \
    api             \
    game            \
    history         \
    prometheus      \
; do 
    push ${NAME}
done
