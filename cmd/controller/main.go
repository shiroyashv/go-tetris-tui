//go:build windows

package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

const (
	VK_LEFT  = 0x25
	VK_UP    = 0x26
	VK_RIGHT = 0x27
	VK_DOWN  = 0x28
	VK_SPACE = 0x20

	DasDelay = 170 * time.Millisecond
	DasSpeed = 33 * time.Millisecond
)

func isPressed(vk int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vk))
	return ret&0x8000 != 0
}

// getWSLIP запрашивает IP адрес внутри WSL, выполняя команду wsl.exe
func getWSLIP() string {
	cmd := exec.Command("wsl", "hostname", "-I")
	out, err := cmd.Output()
	if err != nil {
		return "127.0.0.1"
	}
	ip := strings.Fields(string(out))[0]
	return ip
}

func main() {
	wslIP := getWSLIP()
	address := fmt.Sprintf("%s:9999", wslIP)
	fmt.Printf("Connecting to WSL at: %s\n", address)

	serverAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Error resolving:", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Controller running! Press arrow keys.")

	dasTimers := make(map[string]time.Time)
	lastMoveTime := make(map[string]time.Time)
	ticker := time.NewTicker(8 * time.Millisecond)

	for range ticker.C {
		handleDasKey(conn, VK_LEFT, "left", dasTimers, lastMoveTime)
		handleDasKey(conn, VK_RIGHT, "right", dasTimers, lastMoveTime)
		handleDasKey(conn, VK_DOWN, "down", dasTimers, lastMoveTime)
		handleSingleKey(conn, VK_UP, "up", dasTimers)
		handleSingleKey(conn, VK_SPACE, "space", dasTimers)
	}
}

func handleDasKey(conn *net.UDPConn, vk int, name string, timers, lastMoves map[string]time.Time) {
	if isPressed(vk) {
		now := time.Now()
		startTime, active := timers[name]
		if !active {
			send(conn, name)
			timers[name] = now
			lastMoves[name] = now
		} else {
			if now.Sub(startTime) > DasDelay {
				if now.Sub(lastMoves[name]) > DasSpeed {
					send(conn, name)
					lastMoves[name] = now
				}
			}
		}
	} else {
		delete(timers, name)
		delete(lastMoves, name)
	}
}

func handleSingleKey(conn *net.UDPConn, vk int, name string, timers map[string]time.Time) {
	if isPressed(vk) {
		if _, active := timers[name]; !active {
			send(conn, name)
			timers[name] = time.Now()
		}
	} else {
		delete(timers, name)
	}
}

func send(conn *net.UDPConn, cmd string) {
	data := struct{ C string }{C: cmd}
	bytes, _ := json.Marshal(data)
	conn.Write(bytes)
}
