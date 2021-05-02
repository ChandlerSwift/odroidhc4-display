# odroidhc4-display

odroidhc4-display is a package to provide basic helpful output for the
[OLED display on the ODROID-HC4](https://wiki.odroid.com/odroid-hc4/application_note/oled).
It's intended to be a no-dependencies-required download-and-run alternative for
the `odroid-homecloud-display` (the base install requires pillow, which has some
native dependencies; the install of the examples requires pygame and all its
dependencies as well).

### Installing
```sh
# from source
git clone git@github.com:ChandlerSwift/odroidhc4-display.git
cd odroidhc4-display
go build # if not built elsewhere; you can also cross-compile with GOOS=linux GOARCH=arm64
sudo cp odroidhc4-display /usr/bin/
sudo cp odroidhc4-display.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now odroidhc4-display

# or prebuilt
sudo curl -Lo /usr/bin/odroidhc4-display $(curl -s https://api.github.com/repos/ChandlerSwift/odroidhc4-display/releases/latest | jq -r ".assets[0].browser_download_url") 
sudo chmod +x /usr/bin/odroidhc4-display
sudo curl -Lo /etc/systemd/system/odroidhc4-display.service https://raw.githubusercontent.com/ChandlerSwift/odroidhc4-display/main/odroidhc4-display.service
sudo systemctl daemon-reload
sudo systemctl enable --now odroidhc4-display
```
