#!/bin/sh
docker login -u $DOCKER_ACCOUNT -p $DOCKER_PASSWORD
if [ "$TRAVIS_BRANCH" = "master" ]; then
    TAG="latest"
else
    TAG="$TRAVIS_BRANCH"
fi
docker push $IMAGE_NAME:latest
