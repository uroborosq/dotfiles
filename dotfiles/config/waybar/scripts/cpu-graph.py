#!/usr/bin/env python3
"""Stream a 24-second CPU usage sparkline to Waybar."""

from collections import deque
import json
import time


def read_cpu():
    with open('/proc/stat', encoding='ascii') as stat:
        # Guest time is already included in user/nice; count only through steal.
        values = [int(value) for value in stat.readline().split()[1:9]]
    return sum(values), values[3] + values[4]


def main():
    bars = '▁▂▃▄▅▆▇█'
    history = deque(maxlen=24)
    previous_total, previous_idle = read_cpu()
    while True:
        time.sleep(1)
        total, idle = read_cpu()
        elapsed = total - previous_total
        usage = 100 * (1 - (idle - previous_idle) / elapsed) if elapsed > 0 else 0
        usage = max(0, min(100, usage))
        previous_total, previous_idle = total, idle
        history.append(bars[min(7, int(usage * 8 / 100))])
        print(json.dumps({
            'text': ''.join(history).rjust(24),
            'tooltip': f'CPU: {usage:.1f}%\nПоследние 24 секунды · шкала 0–100%\nНовые значения справа',
            'percentage': round(usage),
        }, ensure_ascii=False), flush=True)


if __name__ == '__main__':
    try:
        main()
    except (BrokenPipeError, KeyboardInterrupt):
        pass
