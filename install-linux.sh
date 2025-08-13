#!/bin/bash

workdir="/home/$USER/.local/fast-share.app"
mkdir -p "$workdir"
cd "$workdir"

echo "Cleaning up previous installation ..."
rm -f /home/$USER/.local/bin/fshare
rm -f /home/$USER/.local/share/applications/fast-share.desktop
rm -f fhsare
rm -f icon.png

echo "Downloading binary ..."
curl -L -o fshare https://github.com/taufiqmaulana/fast-share/releases/download/v1.0.0-beta/fshare-linux-amd64
curl -L -o icon.png https://raw.githubusercontent.com/taufiqmaulana/fast-share/refs/heads/master/icon.png

ln -s /home/$USER/.local/fast-share.app/fshare /home/$USER/.local/bin/fshare

cat > /home/$USER/.local/share/applications/fast-share.desktop <<EOL
[Desktop Entry]
Type=Application
Name=Fast Share
Terminal=true
NoDisplay=true
Icon=/home/$USER/.local/fast-share.app/icon.png
Exec=/home/$USER/.local/fast-share.app/fshare "%f"
EOL

chmod +x /home/$USER/.local/fast-share.app/fshare
update-desktop-database /home/$USER/.local/share/applications
echo "Installation completed"
