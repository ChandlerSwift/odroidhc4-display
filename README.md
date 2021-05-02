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
```

### Installing
```sh
git clone git@github.com:ChandlerSwift/odroidhc4-display.git
cd odroidhc4-display
go build # if not built elsewhere; you can also cross-compile as shown above
sudo cp odroidhc4-display /usr/bin/
sudo cp odroidhc4-display.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now odroidhc4-display
```

