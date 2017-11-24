# Flashy LEDs for the Pi Hut 3D Xmas Tree

## …what?

I've got [one of these](https://thepihut.com/products/3d-xmas-tree-for-raspberry-pi), and I wanted it to live atop my Raspberry Pi running [OSMC](https://osmc.tv/). The software tab on the product page says:

> The easiest way to control your XMAS board is with GPIO Zero. This is pre-installed in Raspbian Stretch

…and sadly for me, whatever version of OSMC I have is based on Debian Jessie. `apt` didn't have the [GPIO Zero libraries](https://gpiozero.readthedocs.io/en/stable/installing.html), and I don't know enough about `pip` / the-Python-ecosystem to fix that. I don't know Python anyway, so I wrote this, in Go.

## …ok.

Download the binary from the [releases tab](https://github.com/nickgrim/xmas-pi/releases). Or build it yourself if you've got a working Go environment; remember to add `GOARCH=arm` if you're not building on a Raspberry Pi (or other ARM device).
