#!/bin/python3

import sh

priv = sh.playerctl('metadata').split('\n')

for i in priv:
    if i.find('title') != -1:
        print()
