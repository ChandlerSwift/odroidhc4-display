# odroidhc4-display

odroidhc4-display is a package to provide basic helpful output for the
[OLED display on the ODROID-HC4](https://wiki.odroid.com/odroid-hc4/application_note/oled).
It's intended to be a no-dependencies-required download-and-run alternative for
the `odroid-homecloud-display` (the base install requires pillow, which has some
native dependencies; the install of the examples requires pygame and all its
dependencies as well).

### Building
To cross-compile:
```sh
GOARCH=arm64 GOOS=linux go build
