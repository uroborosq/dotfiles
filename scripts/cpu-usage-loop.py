#!/bin/python

import time

def get_stats(s: str) -> [int]:
    values_of_first_time = s.split()[1:]
    total = 0
    work = 0
    values_count = len(values_of_first_time)
    i = 0
    while i < values_count:
        parsed_value = int(values_of_first_time[i])
        total += parsed_value
        if i < 3:
            work += parsed_value
        i+=1
    return total, work

if __name__ == '__main__':
    while True:
        file = open('/proc/stat')
        first_total, first_work = get_stats(file.readline())
        time.sleep(1)
        file.seek(0)
        second_total, second_work = get_stats(file.readline())
        print(100 * ((second_work - first_work) / (second_total - first_total)))
        file.close()
