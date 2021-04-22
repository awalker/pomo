# pomo
Pomodoro timer for use with rofi

## Currently

This is more of a test of intergrating with rofi.

## Planned Features

* CLI interface to timers
* Daemon for the timers
* Rofi interface
* GUI interface

## Rofi Notes

Rofi can support two interface methods: dmenu and modi.

### Dmenu

Dmenu mode takes stdin and converts that to a menu. The selection is printed (with a newline) to stdout and rofi exits.

Our app start rofi. Our app is in control.

Multi-level menus mean rofi flickers on and off. Free-form entry is possible.

### Modi

In this mode, rofi starts our app and formats stdout into a menu. Selections are feed to the next instance of our app
as a parameter. Prompts and messages can be controlled with special characters in the output.

Rofi is in control. Our app runs multiple times. Rofi doesn't flicker in multi-level menus. Any context from a multi-level
menu needs to be part of the selection. Free-form entry is more difficult.

## Setup
...

## CLI Usage
...

