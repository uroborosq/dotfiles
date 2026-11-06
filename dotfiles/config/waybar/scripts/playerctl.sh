#!/bin/sh

status=$(playerctl -s status 2>/dev/null) || status=Stopped

case $status in
    Playing)
        icon=
        tooltip=Воспроизведение
        ;;
    Paused)
        icon=
        tooltip=Пауза
        ;;
    *)
        icon=
        tooltip='Нет воспроизведения'
        ;;
esac

if [ "$status" = Playing ] || [ "$status" = Paused ]; then
    artist=$(playerctl -s metadata --format '{{markup_escape(artist)}}' 2>/dev/null) || artist=
    title=$(playerctl -s metadata --format '{{markup_escape(title)}}' 2>/dev/null) || title=

    if [ -n "$artist" ] && [ -n "$title" ]; then
        tooltip="$artist — $title"
    elif [ -n "$title" ]; then
        tooltip=$title
    elif [ -n "$artist" ]; then
        tooltip=$artist
    fi
fi

jq -cn --arg text "$icon" --arg tooltip "$tooltip" '{text: $text, tooltip: $tooltip}'
