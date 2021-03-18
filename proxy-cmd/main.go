package main

// cc: https://github.com/smallnest/1m-go-tcp-server/blob/master/1_simple_tcp_server/server.go

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

const cmd = "stdbuf"

var addr string
var cmdArgs = []string{"-i0", "-o0", "-e0"}
var stdinAsArg1 = 0 != len(os.Getenv("PROXY_CMD_STDIN_AS_ARG1"))

const maxSTDINSize = 50 * 1024 // 50 KB
const stdinTimeout = 5 * time.Second

func main() {
	// setrLimit()
	if len(os.Args) < 4 {
		log.Println(`Usage:`, os.Args[0], " working_dir listen_address command args...\r\n\tEg:", os.Args[0], "/opt :9999 pwd\r\n\r\nSet non-empty environ variable PROXY_CMD_STDIN_AS_ARG1 to pass stdin as first arg to prog.")
		os.Exit(1)
	}

	err := os.Chdir(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	addr = os.Args[2]

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Listening", addr)

	cmdArgs = append(cmdArgs, os.Args[3:]...)

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Println("Accept temp err:", ne)
				continue
			}
			log.Println("Accept err:", e)
			return
		}
		if stdinAsArg1 {
			go handleConn1(conn)
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn1(conn net.Conn) {
	defer func() {
		time.Sleep(time.Second)
		conn.Close()
	}()
	conn.Write([]byte("We take your input to pass to the program as an argument. We accept binary formats like \xff, except null-character (\\x00). We will wait 5 seconds for you to finish typing.\r\n"))
	var data []byte
	lenData := 0
	buf := make([]byte, 1024)
	for {
		conn.SetReadDeadline(time.Now().Add(stdinTimeout))
		n, err := conn.Read(buf)
		if n != 0 {
			if n >= maxSTDINSize-lenData {
				buf = buf[:maxSTDINSize-lenData]
				data = append(data, buf...)
				break
			}
			buf = buf[:n]
			lenData += n
			data = append(data, buf...)
		}
		if err != nil {
			// log.Println("Read remote:", err)
			break
		}
	}
	arg1 := string(data)
	newArgs := append(cmdArgs, arg1)
	cmd := exec.Command(cmd, newArgs...)
	cmd.Stdout = conn
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Run cmd:", err)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		time.Sleep(time.Second)
		conn.Close()
	}()

	var setDone sync.Once
	done := make(chan struct{})

	cmd := exec.Command(cmd, cmdArgs...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Stdout:", err)
		return
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println("Stdin:", err)
		return
	}
	cmd.Stderr = os.Stderr

	go func() {
		_, err := io.Copy(conn, stdout)
		if err != nil {
			// log.Println("Write remote:", err)
		}
		setDone.Do(func() {
			close(done)
		})
	}()
	go func() {
		_, err := io.Copy(stdin, conn)
		if err != nil {
			// log.Println("Read remote:", err)
		}
		setDone.Do(func() {
			close(done)
		})
	}()

	err = cmd.Start()
	if err != nil {
		log.Println("Run cmd:", err)
		return
	}
	defer cmd.Process.Kill()

	<-done
}

func setrLimit() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Panicln(err)
	}
	rLimit.Cur = rLimit.Max
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Panicln(err)
	}
}
