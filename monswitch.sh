#/bin/bash

if [[ $(xrandr --listactivemonitors | wc -l) == "2" ]]; then
	gnome-randr --output eDP-1 --left-of HDMI-1 --auto --persistent
	gnome-randr --output HDMI-1 --mode 1920x1080 --rate 72 --primary
else
	gnome-randr --output eDP-1 --off --persistent
	gnome-randr --output HDMI-1 --mode 1920x1080 --rate 72 --primary
fi
