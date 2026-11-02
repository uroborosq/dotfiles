import re

from kittens.tui.handler import result_handler
from kitty.key_encoding import KeyEvent, parse_shortcut


def is_window_vim(window, vim_id):
    fp = window.child.foreground_processes
    return any(re.search(vim_id, p['cmdline'][0] if len(p['cmdline']) else '', re.I) for p in fp)


def encode_key_mapping(window, key_mapping):
    mods, key = parse_shortcut(key_mapping)
    event = KeyEvent(
        mods=mods,
        key=key,
        shift=bool(mods & 1),
        alt=bool(mods & 2),
        ctrl=bool(mods & 4),
        super=bool(mods & 8),
        hyper=bool(mods & 16),
        meta=bool(mods & 32),
    ).as_window_system_event()

    return window.encoded_key(event)


def send_to_child(window, key_mapping):
    for keymap in key_mapping.split(">"):
        window.write_to_child(encode_key_mapping(window, keymap))


def main():
    pass


@result_handler(no_ui=True)
def handle_result(args, result, target_window_id, boss):
    direction = args[1]
    key_mapping = args[2]
    vim_id = args[3] if len(args) > 3 else "n?vim"

    window = boss.window_id_map.get(target_window_id)

    if window is None:
        return

    # В nvim клавиша всегда отдаётся редактору: он сам управляет своими сплитами.
    if is_window_vim(window, vim_id):
        send_to_child(window, key_mapping)
        return

    # Иначе сначала пробуем перейти в соседнее окно kitty. Если соседа в этом
    # направлении нет, клавиша уходит в программу — так её получает tmux
    # (в том числе удалённый, через ssh) и переключает уже свои панели.
    tab = boss.active_tab
    if tab is not None and tab.neighboring_group_id(direction):
        tab.neighboring_window(direction)
    else:
        send_to_child(window, key_mapping)
