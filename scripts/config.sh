#!/bin/env bash

cd ../config

for i in $(find .); do
    path=$(echo $i | cut -c2-)
    if [[ -d $i ]]; then
        mkdir -p $path
    elif [[ -f $i ]]; then
        ln -f $i $path
    fi
done