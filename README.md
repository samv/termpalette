
# README for termpalette

This library is a small and stupid way to set the terminal to have a
fixed palette.

This all started with
[this issue](https://github.com/aybabtme/humanlog/issues/3) and
descended into the silliness within.

## Supported Terminals

* xterm (Xorg & similar)
* iTerm2

## Unsupported Terminals

* MacOS Terminal.app does not support the necessary escape sequences
  to control the themek, or if it does they are not documented
  anywhere.

