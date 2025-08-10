#!/bin/bash

workdir="/home/$USER/.local/fast-share.app"
mkdir -p "$workdir"
cd "$workdir"

cat > /home/$USER/.local/fast-share.app/fast-share.desktop <<EOL
[Desktop Entry]
Type=Application
Name=Fast Share
Terminal=true
NoDisplay=true
Icon=/home/$USER/.local/fast-share.app/icon.png
Exec=/home/$USER/.local/fast-share.app/main "%f"
EOL
