#!/bin/bash
while true
    do
        clear
        sensors | grep "Tctl\|edge\|Sensor 1\|Sensor 2" | awk -F+ '{print $2}' | cut -c 1-7 | tr '\n' ' '
        echo
        sleep 1
    done