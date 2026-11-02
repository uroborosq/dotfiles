#!/bin/sh
# Stream CAVA frames as Unicode bars for Waybar's custom module.
config_dir=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd) || exit 1
cava -p "$config_dir/cava.config" |
    sed -u '/^[0-7;][0-7;]*$/!d; s/;//g; y/01234567/▁▂▃▄▅▆▇█/'
