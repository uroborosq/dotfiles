#!/bin/bash

while [[ $(playerctl status) == 'Playing' ]]; do
    echo sleep
    sleep 5
done
