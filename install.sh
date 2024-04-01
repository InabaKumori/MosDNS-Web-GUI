#!/bin/bash

# Check for root privileges
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root" 
   exit 1
fi

# Set installation directory
install_dir="/usr/local/mosdns-web-gui"

# Create installation directory
mkdir -p "$install_dir"

# Copy binary executable
cp mosdns-web-gui "$install_dir"

# Create configuration directory
mkdir -p "$install_dir/config"

# Copy configuration file (if needed)
# cp mosdns.yaml "$install_dir/config"

# Set permissions
chown -R root:root "$install_dir"
chmod +x "$install_dir/mosdns-web-gui"

echo "MosDNS Web GUI installed successfully!"
echo "Please edit the configuration file if needed and then run:"
echo "$install_dir/mosdns-web-gui"
