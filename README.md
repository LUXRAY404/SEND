<h1 align="center">SEND FLASH</h1>

**SEND FLASH** is a terminal-based tool written in **Bash** that allows you to securely send and receive files and folders using the [magic-wormhole](https://github.com/warner/magic-wormhole) protocol. It features a clean, interactive CLI interface with color-coded menus, smooth animations, and **self-installation for global access**.

---

## ⚙️ Features

* Send and receive files/folders securely via Wormhole.
* Automatic Wormhole installation if missing (`sudo apt install magic-wormhole -y`).
* Self-installing Bash version (`/usr/local/bin/sendflash`) for global access.
* Interactive color-coded CLI menu.
* Typewriter animation when exiting.
* Error handling for missing files, invalid inputs, and installation issues.
* Default save location for received files: `~/Downloads/sendflash_received`.

---

## 🛠️ Installation

### Prerequisites

* Linux or WSL environment.
* **sudo** privileges (for installing Wormhole or Bash version globally).

### Install Bash Version

1. Download or clone the Bash script:

```bash
wget https://raw.githubusercontent.com/INTELEON404/sendflash/main/sendflash.sh
````

2. Make it executable:

```bash
chmod +x sendflash.sh
```

3. Run the script. The first run will automatically:

   * Move the script to `/usr/local/bin/sendflash`
     ```
     mv sendflash.sh /usr/local/bin/sendflash
     ```
   * Add `/usr/local/bin` to your PATH if needed
   * Make it globally executable

```bash
bash sendflash.sh
```

After installation, you can run it anywhere with:

```bash
sendflash
```

---

## 📂 Usage

When you run `sendflash`, the menu will appear:

```
┌──(SENDFLASH㉿WORMHOLE)-[~]
└─$ SENDFLASH

[01] SEND A FILE/FOLDER
[02] RECEIVE A FILE/FOLDER
[03] UPDATE THIS TOOL
[04] EXIT
```

* **Send a file/folder**: `[01]` → Enter the full path to your file or folder. You will receive a **Wormhole key** to share with the receiver.
* **Receive a file/folder**: `[02]` → Enter the key provided by the sender. Files are saved to `~/Downloads/sendflash_received`.
* **Update tool**: `[03]` → Checks for updates (currently displays "up-to-date").
* **Exit**: `[04]` → Exits the tool with a typewriter animation.

---

### Quick Demo Workflow

**Sender Terminal:**

```bash
sendflash
# 01
# /home/user/Documents/test.txt
# → Displays key: 7-crystal-cupcake
```

**Receiver Terminal:**

```bash
sendflash
# 02
# 7-crystal-cupcake
# → File saved to ~/Downloads/sendflash_received/test.txt
```

---

## 🧪 Development

1. **Fork** the repository.
2. **Clone** your fork:

```bash
git clone https://github.com/LUXRAY404/sendflash.git
cd sendflash
```

3. Make changes, commit, and push.
4. Open a **pull request** to merge your improvements.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


