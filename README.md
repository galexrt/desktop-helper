# desktop-helper
A program to automatically switch your screen configurations and/or execute commands on location (network, screen, docked) changes.

The documentation can be found on the ReadTheDocs here: http://desktop-helper.readthedocs.io/en/latest/

## Installation
To install the binary run the following command as root (sudo) to put into the `/usr/local/bin` directory:
```
# curl https://github.com/galexrt/desktop-helper/releases/download/v0.1.0/desktop-helper_linux_amd64 > /usr/local/bin/desktop-helper
```

### Building
This section shows how to build the desktop-helper yourself.

#### Requirements
* Golang (tested with 1.8.x and higher)
    * With `GOPATH` set.
* `libnotify`
    * Debian Jessie: `libnotify-dev`
    * Archlinux: `libnotify`

#### Instructions
1. Run `go get github.com/galexrt/desktop-helper`.
2. Run `CGO go build github.com/galexrt/desktop-helper/cmd/desktophelper -o "$GOPATH/bin/desktop-helper"`

This results in a binary named `desktop-helper` located at `$GOPATH/bin/`. Now you just need to copy the binary somewhere in your `$PATH`, example locations `/usr/local/bin`.

## Configuration
An example configuration can be found here `config.example.yaml`.

### Basic two profile example
```yaml
```
