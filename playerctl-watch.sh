#!/bin/bash

installed_suspend_time_ac=$(gsettings get org.gnome.settings-daemon.plugins.power sleep-inactive-ac-timeout)
installed_suspend_time_battery=$(gsettings get org.gnome.settings-daemon.plugins.power sleep-inactive-battery-timeout)

function restore()
{
    echo "Interrupt signal received, restoring default settings..."
    gsettings set org.gnome.settings-daemon.plugins.power sleep-inactive-ac-timeout $installed_suspend_time_ac
    gsettings set org.gnome.settings-daemon.plugins.power sleep-inactive-battery-timeout $installed_suspend_time_battery
    echo "Restored, exiting..."
}

trap restore exit
suspend_enabled=1
while true; do
    playerctl_status=$(playerctl status)
    if [[ "$playerctl_status" == 'Playing' ]] && [[ $suspend_enabled -eq 1 ]]
    then
        echo "Audio playing was detected, suspend disabled"
        suspend_enabled=0
        gsettings set org.gnome.settings-daemon.plugins.power sleep-inactive-ac-timeout 0
        gsettings set org.gnome.settings-daemon.plugins.power sleep-inactive-battery-timeout 0
    elif [[ "$playerctl_status" != "Playing" ]] && [[ $suspend_enabled -eq 0 ]]
    then
        echo "Suspend enabled"
        suspend_enabled=1
        gsettings set org.gnome.settings-daemon.plugins.power sleep-inactive-ac-timeout $installed_suspend_time_ac
        gsettings set org.gnome.settings-daemon.plugins.power sleep-inactive-battery-timeout $installed_suspend_time_battery
    fi
    sleep 240
done
