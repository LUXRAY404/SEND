# SEND FLASH

**SEND FLASH** is a terminal-based tool written in Go that allows you to securely send and receive files and folders using the [magic-wormhole](https://github.com/warner/magic-wormhole) protocol. It features a clean, interactive CLI interface with color-coded menus and smooth animations for a polished user experience.

---

## ⚙️ Features

* Send and receive files/folders securely with Wormhole.
* Automatic Wormhole installation via `sudo apt install magic-wormhole -y`.
* Interactive and color-coded CLI interface.
* Smooth typewriter animation when exiting.
* Error handling for missing files, installation issues, and invalid inputs.

---

## 🛠️ Installation

### Prerequisites

* **Go** version 1.18 or higher.
* **sudo** privileges for Wormhole installation.

### Install via Go

Simply run:

```bash
go install github.com/LUXRAY404/sendflash@latest
```

This will download, build, and install the `sendflash` binary in your Go bin directory (usually `$HOME/go/bin`).

Make sure `$HOME/go/bin` is in your `PATH`:

```bash
export PATH=$PATH:$HOME/go/bin
```

Then run:

```bash
sendflash
```

The tool will check for Wormhole and install it automatically if needed.

---

## 📂 Usage

When you run `sendflash`, you’ll see:

```
┌──(SENDFLASH㉿WORMHOLE)-[~]
└─$ SENDFLASH

[01] SEND A FILE/FOLDER
[02] RECEIVE A FILE/FOLDER
[03] UPDATE THIS TOOL
[04] EXIT
```

* **Send a file/folder**: `[01]` → Enter file/folder path.
* **Receive a file/folder**: `[02]` → Enter the receiver key from the sender.
* **Update tool**: `[03]` → Checks for updates (currently displays "up-to-date").
* **Exit**: `[04]`.

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


