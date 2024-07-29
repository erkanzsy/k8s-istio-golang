#!/bin/bash
VERSION="v1.0.1"
USERNAME=erkanozsoy

docker build --tag app-v2:$VERSION .

docker tag app-v2:$VERSION $USERNAME/app-v2:$VERSION

docker push $USERNAME/app-v2:$VERSION

docker run -d -p 3001:3001 $USERNAME/app-v2:$VERSION

open http://localhost:3001