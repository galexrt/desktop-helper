# Building
This section shows how to build the desktop-helper yourself.

## Requirements
* Golang (tested with 1.8.x and higher)
    * With `GOPATH` set.
* `libnotify`
    * Debian Jessie: `libnotify-dev`
    * Archlinux: `libnotify`

## Instructions
1. Run `go get github.com/galexrt/desktop-helper`.
2. Run `CGO go build github.com/galexrt/desktop-helper/cmd/desktophelper -o "$GOPATH/bin/desktop-helper"`

This results in a binary named `desktop-helper` located at `$GOPATH/bin/`. Now you just need to copy the binary somewhere in your `$PATH`, example locations `/usr/local/bin`.
