#!/bin/bash
# SETUP.SH
# Auto-install SendFlash and dependencies

# Colors
GREEN="\033[1;32m"
YELLOW="\033[1;33m"
RED="\033[1;31m"
NC="\033[0m"

# Check if script exists in current folder
SCRIPT_NAME="sendflash.sh"
if [ ! -f "$SCRIPT_NAME" ]; then
    echo -e "${RED}Error: $SCRIPT_NAME not found in current directory!${NC}"
    exit 1
fi

# Install dependencies
echo -e "${YELLOW}Installing required tools...${NC}"
sudo apt update
sudo apt install -y magic-wormhole

# Make script executable
chmod +x "$SCRIPT_NAME"

# Move to /usr/local/bin
echo -e "${YELLOW}Moving $SCRIPT_NAME to /usr/local/bin/sendflash...${NC}"
sudo mv "$SCRIPT_NAME" /usr/local/bin/sendflash

# Ensure /usr/local/bin is in PATH
if ! echo "$PATH" | grep -q "/usr/local/bin"; then
    echo 'export PATH=$PATH:/usr/local/bin' >> "$HOME/.bashrc"
    export PATH=$PATH:/usr/local/bin
    echo -e "${GREEN}/usr/local/bin added to PATH${NC}"
fi

echo -e "${GREEN}Installation complete! You can now run 'sendflash' from anywhere.${NC}"
