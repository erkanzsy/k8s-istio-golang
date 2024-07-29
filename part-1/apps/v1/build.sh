#!/bin/bash
VERSION="v1.0.1"
USERNAME=erkanozsoy

docker build --tag app-v1:$VERSION .

docker tag app-v1:$VERSION $USERNAME/app-v1:$VERSION

docker push $USERNAME/app-v1:$VERSION

docker run -d -p 3000:3000 $USERNAME/app-v1:$VERSION

open http://localhost:3000

