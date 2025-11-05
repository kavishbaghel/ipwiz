#!/bin/bash

echo "Build ipez binary"

go build -o binaries/ipez

if [ $? -eq 0 ]; then
    echo "Build successful"
    cp binaries/ipez /usr/local/bin
    echo -e "\n\n"
else
    echo "Build failed"
fi

ipez --help

