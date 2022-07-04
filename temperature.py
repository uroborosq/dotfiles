#!/bin/python3

import subprocess


result = subprocess.run(['sensors'], capture_output=True, text=True).stdout

indexes = [i for i in range(len(result)) if result.startswith('+', i)]
for i in [1, 0, 2, 7]:
    print(result[indexes[i] + 1:indexes[i] + 7], end=' ')
print()