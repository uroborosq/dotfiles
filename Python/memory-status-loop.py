#!/bin/python

import time
file = open('/proc/meminfo', 'r')
log_file = open('~/Документы/log.txt', 'w+')
log_file.write('new launch')
while True:
    file.seek(0)
    memory_total = int(file.readline().split()[1])
    file.readline()
    memory_available = int(file.readline().split()[1])
    print(f"{round((memory_total - memory_available) / 1024 / 1024, 1)} GiB")
    time.sleep(5)


