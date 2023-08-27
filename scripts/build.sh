#!/bin/bash 

cd ../cmd

for i in $(ls); do
    if [[ $i != "bin" ]] && [[ -d $i ]]; then
        cd $i
        echo "Building $i"
        go mod download 2>/dev/null && echo "Dependencies checked" && go build -o ../bin/$i main.go && echo "Build successful"
        cd ..
    fi
done