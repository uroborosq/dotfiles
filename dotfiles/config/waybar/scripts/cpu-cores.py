#!/usr/bin/env python3
"""Stream one CPU usage bar per logical CPU to Waybar."""

import json
import time


BARS = '▁▂▃▄▅▆▇█'


def colored_bar(load):
    # Match the gruvbox-material palette in style.css.
    if load >= 90:
        color = '#ea6962'
    elif load >= 70:
        color = '#e78a4e'
    elif load >= 40:
        color = '#d8a657'
    else:
        color = '#a9b665'
    bar = BARS[min(7, int(load * 8 / 100))]
    return f'<span foreground="{color}">{bar}</span>'


def read_cpus():
    cpus = {}
    with open('/proc/stat', encoding='ascii') as stat:
        for line in stat:
            fields = line.split()
            if not fields or not fields[0].startswith('cpu'):
                continue
            index = fields[0][3:]
            if not index.isdigit():
                continue
            # Guest time is already included in user/nice.
            values = [int(value) for value in fields[1:9]]
            cpus[int(index)] = (sum(values), values[3] + values[4])
    return cpus


def usage(previous, current):
    total = current[0] - previous[0]
    idle = current[1] - previous[1]
    if total <= 0:
        return 0.0
    return max(0.0, min(100.0, 100 * (total - idle) / total))


def frame(previous, current):
    loads = [(index, usage(previous.get(index, counters), counters))
             for index, counters in sorted(current.items())]
    return {
        'text': ''.join(colored_bar(load) for _, load in loads),
        'tooltip': 'Загрузка логических ядер · шкала 0–100%\n'
                   + '\n'.join(f'CPU {index}: {load:.1f}%'
                               for index, load in loads),
    }


def main():
    previous = read_cpus()
    while True:
        time.sleep(1)
        current = read_cpus()
        print(json.dumps(frame(previous, current), ensure_ascii=False), flush=True)
        previous = current


if __name__ == '__main__':
    try:
        main()
    except (BrokenPipeError, KeyboardInterrupt):
        pass
