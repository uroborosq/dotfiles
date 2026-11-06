# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`~/.config/waybar` — the Waybar (v0.15.0) configuration for this machine. There is no build, test or lint step; the artifacts are JSON configs and a GTK CSS stylesheet consumed directly by the running bar.

## Layout and how the pieces are wired

Waybar is never started with the default `~/.config/waybar/config` name. Each window manager launches it with an explicit `--config` path, through a wrapper:

- Sway — `~/.config/sway/config.d/autostart:46` runs `~/.nimble/bin/uq_waybar ~/.config/waybar/sway.config.json`. `uq_waybar` is a small Nim supervisor (from the `uq_tools_nim` nimble package) that spawns `waybar --config <path>` and restarts it.
- Hyprland — `~/.config/hypr/hyprland.conf:33` runs `uq-exec-waybar --path ~/.config/waybar/hyprland.config.json`.

Because neither wrapper passes `-s`, Waybar falls back to auto-discovery for the stylesheet: **`style.css` in this directory styles both configs**.

`sway.config.json` pins `"height": 34`, so the bar no longer sizes itself to its content: the vertical `margin` on `.modules-*` in `style.css` is what controls the visible strip of bar background above and below each group. Changing one without the other changes the bar's height instead of the inset.

`sway.config.json` defines a full bar on every monitor except Wisecoco and the small YUE panel (currently on `HDMI-A-1`), plus a compact bar on YUE. It matches by make + model rather than port, so it survives replugging. Waybar builds that identifier as `make model serial` but drops a serial Sway reports as `Unknown` — the string `"YUE 2560x1600 CVT Unknown"` does **not** match. Check the exact identifiers with `waybar -l debug -c … 2>&1 | grep "Output detection"`.

`sway.config.json` is a top-level array with one object per bar. Both bars include `sway.modules.json`, a top-level object with their shared module settings. The YUE bar overrides the workspace format for its compact display.

### Workspaces (Sway)

`sway/workspaces` in `sway.modules.json` carries `persistent-workspaces` for 1–4 plus `window-rewrite` rules that turn each window's `app_id`/X11 class into an icon on the full bar (`format: "{name} {windows}"`).

This replaced `sworkstyle`, which used to rename workspaces to `N: <иконки>` from `~/.config/sway/config.d/autostart:52` (the line is now commented out). The two cannot coexist: Waybar matches `persistent-workspaces` keys against the workspace **name** as a plain string, so a persistent entry `"1"` does not match a workspace renamed to `1: ` and Waybar renders both buttons. Anything that renames workspaces away from bare numbers brings the duplicates back.

Add an icon for a new application under `window-rewrite` (`"class<app_id>": "<glyph>"`); unmatched windows fall back to `window-rewrite-default`. Get the id with `swaymsg -t get_tree | jq -r '.. | .app_id? // (.window_properties?.class) | select(.)'`.

## style.css is hand-written (wallust больше его не трогает)

`style.css` редактируется напрямую: grouped-sections design, `Mononoki Nerd Font Propo`.
Палитра (с 2026-10-07) — gruvbox-material в том же варианте, что в nvim (`background = "soft"`,
`foreground = "material"`, тёмный фон; см. `~/.config/nvim/lua/plugins/colorscheme.lua`).
Значения взяты из `~/.local/share/nvim/lazy/gruvbox-material/autoload/gruvbox_material.vim`,
у каждого `@define-color` в комментарии указано имя цвета из палитры плагина.

Раньше файл генерировался wallust'ом — `~/.config/wallust/wallust.toml` содержал строку
`waybar = {template = 'waybar.css', target = '~/.config/waybar/style.css', new_engine = true}`,
и любой `wallust run` перезаписывал стиль. 2026-09-05 по просьбе пользователя запись из
`wallust.toml` удалена, шаблон `~/.config/wallust/templates/waybar.css` и его локальная копия
`style.wallust.css` удалены. Копии лежат в `backup/`
(`wallust-template-waybar.css.*`, `style.wallust.css.*`, `wallust.toml.*`).
2026-10-07 пользователь удалил wallust целиком (бинарник и `~/.config/wallust`); последняя
копия `wallust.toml` — в `backup/`. `~/.config/wofi/style.css` теперь тоже ведётся вручную
в той же палитре gruvbox-material, что и `style.css` (префикс переменных `wf_`). Той же палитрой
заданы `~/.config/mako/config` и цвета окон в `~/.config/sway/config.d/windowrules` (там раньше был
`include ~/.cache/wal/colors-sway`, теперь значения записаны прямо в файле).

Keep the font family on the **`Propo`** variant (`Mononoki Nerd Font Propo`). The three Nerd Font builds differ only in icon metrics: `… Nerd Font Mono` squeezes icons into one cell so they render visibly undersized; the plain `… Nerd Font` draws them at full size but with a one-cell advance, so the glyph ink runs into the neighbouring character (icons visibly overlap the module dividers and adjacent text); `… Nerd Font Propo` gives each icon an advance equal to its ink — full size, no collisions. Verify a candidate with `pango-view --font="<family> 13px" -q file-with-the-icons.txt`.

Модуль `power-profiles-daemon` окрашивается по текущему профилю: Waybar сам выставляет класс `power-saver` / `balanced` / `performance`, правила и цвета `wb_pp_*` лежат в разделе «Состояния» `style.css`.

The stylesheet also references colours it never defines — `@warning_color`, `@error_color`, `@wm_icon_bg` (used by the `blink-warning` / `blink-critical` keyframes and the battery states). They resolve from the GTK theme, so a battery-warning colour change may need to be added as a `@define-color` rather than looked for in this file.

## Custom modules depend on external binaries

Every `custom/*` module shells out; a module rendering blank almost always means its `exec` is missing rather than a config error. Run the binary by hand to check.

The two configs point at different, non-interchangeable binary sets:

- `sway.config.json` uses the underscore names on `PATH` via `~/.nimble/bin`: `uq_cpuload`, `uq_ram`, `uq_cpufreq`, `uq_net`, `uq_iowait`, `uq_current_song`, `uq_temperature kvadra`.
- `hyprland.config.json` uses absolute hyphenated paths: `/usr/bin/uq-cpu-usage-once`, `/usr/bin/uq-ram`, `/usr/bin/uq-cpu-frequency`, `/usr/bin/uq-net`, `/usr/bin/uq-current-song`, `/usr/bin/uq-thinkpad-e14-gen-temperature`. **None of these exist on this machine** — that config targets a different host (ThinkPad E14 vs. the Kvadra sensor profile used by the Sway config), so do not assume it is runnable here.

The temperature module is hardware-specific by design (`uq_temperature <machine>`), so it is expected to differ between the two files.

Known drift in `hyprland.config.json`: `modules-right` lists `hyprland/language`, but the file only carries a `sway/language` settings block, so those settings are inert. `sway/window` and `custom/close` (`swaymsg kill`) are likewise leftovers from the Sway config.

## Commands

```bash
# Validate a config before applying it
jq . sway.config.json

# Reload the running bar (style + config) without restarting it
killall -SIGUSR2 waybar

# Restart under Sway, the way autostart does
killall uq_waybar; ~/.nimble/bin/uq_waybar ~/.config/waybar/sway.config.json

# See parse errors / module failures on stderr (Ctrl-C to stop)
waybar -l debug -c ~/.config/waybar/sway.config.json

# Check a custom module's output directly
uq_cpuload; uq_net; uq_temperature kvadra
```

When editing, change only the config for the compositor actually in use, and keep in mind the two files are maintained in parallel rather than shared — a module added to one does not appear in the other.
