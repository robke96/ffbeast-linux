#
release:
    fyne package -os linux --release

appimage:
    chmod +x ./build/appimage/build_appimage.sh
    ./build/appimage/build_appimage.sh

