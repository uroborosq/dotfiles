#!/bin/bash

cat /proc/cpuinfo | grep MHz | awk '{print $4}' | sort | tail -n 1
