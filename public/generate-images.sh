#!/bin/bash

# Folder output
DIR="$(pwd)/dummy-images"

# Bikin folder kalau belum ada
mkdir -p "$DIR"

# Array resolusi (widthxheight)
resolutions=(
  "150x150"    # Thumbnail kecil
  "300x300"    # Thumbnail medium
  "500x500"    # Profil

  "640x360"    # SD (16:9)
  "800x600"    # 4:3
  "1024x768"   # XGA
  "1280x720"   # HD
  "1920x1080"  # Full HD
  "2560x1440"  # 2K QHD
  "3840x2160"  # 4K UHD

  "360x640"    # Mobile
  "720x1280"   # HD portrait
  "1080x1920"  # Full HD portrait

  "1200x300"   # Web banner
  "1920x600"   # Wide banner
  "2560x1080"  # Ultra-wide banner

  "1080x1080"  # Instagram post
  "1080x1920"  # Instagram story / TikTok
  "1200x628"   # Facebook share link
  "1500x500"   # Twitter header
  "820x312"    # Facebook cover
)

# Loop download gambar dummy
for res in "${resolutions[@]}"; do
  echo "Downloading $res ..."
  wget -q -O "$DIR/$res.png" "https://playground.fahrigunadi.dev/image/$res.png"
done

echo "✅ Semua dummy image tersimpan di $DIR"
