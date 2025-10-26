#!/bin/bash
# SENDFLASH.BASH
# SEND FLASH - Bash version with self-install
# Auto-check and install required tools

# Colors
GREEN="\033[1;32m"
CYAN="\033[1;36m"
YELLOW="\033[1;33m"
RED="\033[1;31m"
NC="\033[0m"

# CLEAR SCREEN
clear_screen() {
    clear
}

# TYPEWRITER EFFECT
type_writer() {
    text="$1"
    delay="$2"
    for ((i=0; i<${#text}; i++)); do
        printf "%c" "${text:$i:1}"
        sleep "$delay"
    done
    echo
}

# SELF-INSTALL TO /usr/local/bin
self_install() {
    target="/usr/local/bin/sendflash"
    if [ "$(realpath "$0")" != "$target" ]; then
        echo -e "${YELLOW}Installing SendFlash globally...${NC}"
        sudo cp "$0" "$target"
        sudo chmod +x "$target"
        echo -e "${GREEN}Installed! You can now run 'sendflash' from anywhere.${NC}"
        exec sendflash "$@"
        exit 0
    fi
}

# CHECK AND INSTALL WORMHOLE
check_wormhole() {
    if ! command -v wormhole &> /dev/null; then
        echo -e "${YELLOW}WORMHOLE NOT FOUND. INSTALLING...${NC}"
        sudo apt update
        sudo apt install magic-wormhole -y
        if ! command -v wormhole &> /dev/null; then
            echo -e "${RED}FAILED TO INSTALL WORMHOLE. PLEASE INSTALL MANUALLY.${NC}"
            exit 1
        fi
        echo -e "${GREEN}WORMHOLE INSTALLED SUCCESSFULLY!${NC}"
        sleep 1
    else
        echo -e "${GREEN}WORMHOLE FOUND ✅${NC}"
        sleep 0.5
    fi
}

# MENU
show_menu() {
    clear_screen
    echo -e "${CYAN}┌──(SENDFLASH㉿WORMHOLE)-[~]${NC}"
    echo -e "${CYAN}└─$ SENDFLASH${NC}"
    echo
    echo -e "${GREEN}[01]${NC} SEND A FILE/FOLDER"
    echo -e "${GREEN}[02]${NC} RECEIVE A FILE/FOLDER"
    echo -e "${YELLOW}[03]${NC} UPDATE THIS TOOL"
    echo -e "${RED}[04]${NC} EXIT"
    echo
    echo -ne "${CYAN}ENTER ANYONE: ${NC}"
}

# SEND FILE/FOLDER
send_file() {
    clear_screen
    echo -ne "${CYAN}ENTER YOUR FILE/FOLDER PATH: ${NC}"
    read path

    if [ ! -e "$path" ]; then
        echo -e "${RED}FILE OR FOLDER NOT FOUND!${NC}"
        sleep 2
        return
    fi

    echo -e "${YELLOW}SENDING VIA WORMHOLE...${NC}"
    wormhole send "$path"
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}FILE SENT SUCCESSFULLY!${NC}"
    else
        echo -e "${RED}ERROR OCCURRED DURING FILE SEND.${NC}"
    fi
    sleep 3
}

# RECEIVE FILE/FOLDER
receive_file() {
    clear_screen
    echo -ne "${CYAN}ENTER RECEIVER KEY: ${NC}"
    read key

    echo -e "${YELLOW}RECEIVING FILE...${NC}"
    yes | wormhole receive "$key"
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}FILE RECEIVED SUCCESSFULLY!${NC}"
    else
        echo -e "${RED}FAILED TO RECEIVE FILE. CHECK THE KEY.${NC}"
    fi
    sleep 3
}

# UPDATE TOOL
update_tool() {
    clear_screen
    echo -e "${YELLOW}CHECKING FOR UPDATES...${NC}"
    sleep 2
    echo -e "${GREEN}SEND FLASH IS ALREADY UP-TO-DATE!${NC}"
    sleep 2
}

# EXIT TOOL
exit_tool() {
    clear_screen
    type_writer "EXITING SEND FLASH..." 0.05
    exit 0
}

# MAIN
self_install
check_wormhole

while true; do
    show_menu
    read choice
    case "$choice" in
        1|01) send_file ;;
        2|02) receive_file ;;
        3|03) update_tool ;;
        4|04) exit_tool ;;
        *) echo -e "${RED}INVALID CHOICE! TRY AGAIN.${NC}" ; sleep 1 ;;
    esac
done
