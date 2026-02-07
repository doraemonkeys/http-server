package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/doraemonkeys/http-server/internal/echo"
	"github.com/doraemonkeys/http-server/internal/fileserver"
)

const usage = `Usage: http-server <command> [options]

Commands:
  echo    Start an HTTP echo server that returns request details as JSON
  file    Start an HTTP file server that serves a directory

Run 'http-server <command> -h' for more information on a command.`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "echo":
		echoCmd := flag.NewFlagSet("echo", flag.ExitOnError)
		port := echoCmd.Int("port", 6688, "Port to listen on")
		ip := echoCmd.String("ip", "0.0.0.0", "IP address to listen on")
		echoCmd.Parse(os.Args[2:])
		echo.Start(*ip, *port)

	case "file":
		fileCmd := flag.NewFlagSet("file", flag.ExitOnError)
		port := fileCmd.Int("port", 8080, "Port to listen on")
		ip := fileCmd.String("ip", "0.0.0.0", "IP address to listen on")
		dir := fileCmd.String("d", ".", "Directory to serve")
		fileCmd.Parse(os.Args[2:])
		fileserver.Start(*ip, *port, *dir)

	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		fmt.Println(usage)
		os.Exit(1)
	}
}
