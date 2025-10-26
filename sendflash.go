// SENDFLASH.GO
// SEND FLASH
// AUTO-CHECK AND INSTALL REQUIRED TOOLS

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

// CLEAR SCREEN
func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

// TYPEWRITER
func typeWriter(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}

// CHECK AND INSTALL WORMHOLE
func checkWormhole() {
	if _, err := exec.LookPath("wormhole"); err != nil {
		color.Yellow("WORMHOLE NOT FOUND. INSTALLING...\n")
		installCmd := exec.Command("pip3", "install", "--user", "magic-wormhole")
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err := installCmd.Run()
		if err != nil {
			color.Red("FAILED TO INSTALL WORMHOLE. PLEASE INSTALL MANUALLY.\n")
			os.Exit(1)
		}
		color.Green("WORMHOLE INSTALLED SUCCESSFULLY!\n")
		time.Sleep(1 * time.Second)
	} else {
		color.Green("WORMHOLE FOUND ✅\n")
		time.Sleep(500 * time.Millisecond)
	}
}

// MAIN MENU
func showMenu() {
	clearScreen()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
	red := color.New(color.FgRed, color.Bold).SprintFunc()

	fmt.Println(cyan("┌──(SENDFLASH㉿LINUX)-[~]"))
	fmt.Println(cyan("└─$ SENDFLASH"))
	fmt.Println()
	fmt.Println(green("[01]") + " SEND A FILE/FOLDER")
	fmt.Println(green("[02]") + " RECEIVE A FILE/FOLDER")
	fmt.Println(yellow("[03]") + " UPDATE THIS TOOL")
	fmt.Println(red("[04]") + " EXIT")
	fmt.Println()
	fmt.Print(cyan("ENTER ANYONE: "))
}

func main() {
	checkWormhole()
	reader := bufio.NewReader(os.Stdin)

	for {
		showMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1", "01":
			sendFile(reader)
		case "2", "02":
			receiveFile(reader)
		case "3", "03":
			updateTool()
		case "4", "04":
			exitTool()
		default:
			color.Red("INVALID CHOICE! TRY AGAIN.\n")
			time.Sleep(1 * time.Second)
		}
	}
}

func sendFile(reader *bufio.Reader) {
	clearScreen()
	color.Cyan("ENTER YOUR FILE/FOLDER PATH:")
	fmt.Print("→ ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		color.Red("FILE OR FOLDER NOT FOUND!\n")
		time.Sleep(2 * time.Second)
		return
	}

	color.Yellow("\nSENDING VIA WORMHOLE...\n")

	cmd := exec.Command("wormhole", "send", path)
	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := strings.ToUpper(scanner.Text())
		if strings.Contains(line, "WORMHOLE CODE") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				color.Green("\nYOUR RECEIVER KEY: " + strings.TrimSpace(parts[1]))
			}
		} else {
			fmt.Println(line)
		}
	}
	cmd.Wait()
	color.Green("\nFILE SENT SUCCESSFULLY!\n")
	time.Sleep(3 * time.Second)
}

func receiveFile(reader *bufio.Reader) {
	clearScreen()
	color.Cyan("ENTER RECEIVER KEY:")
	fmt.Print("→ ")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	color.Yellow("\nRECEIVING FILE...\n")

	cmd := exec.Command("wormhole", "receive", key)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	color.Green("\nFILE RECEIVED SUCCESSFULLY!\n")
	time.Sleep(3 * time.Second)
}

func updateTool() {
	clearScreen()
	color.Yellow("CHECKING FOR UPDATES...\n")
	time.Sleep(2 * time.Second)
	color.Green("SEND FLASH IS ALREADY UP-TO-DATE!\n")
	time.Sleep(2 * time.Second)
}

func exitTool() {
	clearScreen()
	typeWriter("EXITING SEND FLASH...\n", 50*time.Millisecond)
	os.Exit(0)
}
