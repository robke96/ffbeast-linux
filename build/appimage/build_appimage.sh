#!/usr/bin/env bash
set -euo pipefail

# url
FYNEMETA_URL="https://github.com/ErikKalkoken/fynemeta/releases/download/v0.1.1/fynemeta-0.1.1-linux-amd64.tar.gz"
LINUXDEPLOY_URL="https://github.com/linuxdeploy/linuxdeploy/releases/download/continuous/linuxdeploy-x86_64.AppImage"
# temp dir
SOURCEDIR="temp.Source"
APPDIR="temp.Appdir"

# clean old folders
mkdir -p "$SOURCEDIR" "$APPDIR"

# Download fynemeta
if [ ! -f fynemeta ]; then
    wget -q "$FYNEMETA_URL" -O fynemeta.tar.gz
    tar xf fynemeta.tar.gz
    rm fynemeta.tar.gz
fi

# Download linuxdeploy
if [ ! -f linuxdeploy ]; then
    wget -q "$LINUXDEPLOY_URL" -O linuxdeploy
    chmod +x linuxdeploy
fi

# Get app info
appname=$(./fynemeta lookup -k Details.Name -s .)
appid=$(./fynemeta lookup -k Details.ID -s .)
buildname=$(./fynemeta lookup -k Release.BuildName -s .)

# Extract fyne release package
tar xvfJ "$appname.tar.xz" -C "$SOURCEDIR"
cp -r "$SOURCEDIR"/* "$APPDIR/"

cp ./build/appimage/AppRun "$APPDIR/AppRun"
chmod +x "$APPDIR/AppRun"

NO_STRIP=true ./linuxdeploy \
    --appdir "$APPDIR" \
    -o appimage \
    -e "$APPDIR/usr/local/bin/$buildname" \
    -d "$APPDIR/usr/local/share/applications/$appid.desktop" \
    -i "$APPDIR/usr/local/share/pixmaps/$appid.png"

# cleanup
rm -rf "$SOURCEDIR" "$APPDIR" fynemeta linuxdeploy "$appname.tar.xz"
