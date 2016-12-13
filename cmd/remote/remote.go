package main

import (
	"github.com/pkg/term"
	"os"
	"strconv"
	"net"
	"io"
	"sync"
	"log"

	"bufio"
)

var clients map[net.Conn]bool
var clientsLock *sync.Mutex
var writeLock *sync.Mutex

func main() {
	log.SetOutput(os.Stdout)

	clients = make(map[net.Conn]bool)
	clientsLock = new(sync.Mutex)
	writeLock = new(sync.Mutex)

	if len(os.Args) < 4 {
		log.Printf("usage: %s <serial port> <baud> <addr>\n", os.Args[0])
		os.Exit(1)
	}

	serialPortName := os.Args[1]
	addr := os.Args[3]
	baud, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println("[ERROR] Error parsing baud")
		os.Exit(1)
	}

	serialPort, err := term.Open(serialPortName)
	if err != nil {
		log.Printf("[ERROR] Unable to open Serial Port %+v\n", err)
		os.Exit(1)
	}
	serialPort.SetSpeed(baud)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("[ERROR] Unable to open Listen on Port %s %+v\n", addr, err)
		os.Exit(1)
	}
	defer l.Close()

	go serialPortReader(serialPort)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[ERROR] accepting: ", err.Error())
			continue
		}
		go handleRequest(conn, serialPort)
	}
}

func serialPortReader(serialPort io.Reader) {
	scanner := bufio.NewScanner(serialPort)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		log.Println("[ELK]", scanner.Text())
		writeLock.Lock()
		for conn := range clients {
			_, err := conn.Write(append(scanner.Bytes(), byte('\r'), byte('\n')))
			if err != nil {
				log.Println("[ERROR] writeing to", conn.RemoteAddr(), err)
				continue
			}
		}
		writeLock.Unlock()
		if err := scanner.Err(); err != nil {
			log.Println("[ERROR] reading from elk", err)
			return
		}
	}
}

func handleRequest(conn net.Conn, serialPort io.Writer) {
	defer func() {
		conn.Close()
		clientsLock.Lock()
		delete(clients, conn)
		clientsLock.Unlock()
		log.Printf("[DISCONNECT] %s Total: %d\n", conn.RemoteAddr(), len(clients))
	}()

	clientsLock.Lock()
	clients[conn] = true
	clientsLock.Unlock()

	log.Printf("[CONNECT] %s Total: %d\n", conn.RemoteAddr(), len(clients))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		writeLock.Lock()
		_, err := serialPort.Write(append(scanner.Bytes(), byte('\r'), byte('\n')))
		if err != nil {
			log.Printf("[ERROR] read from conn %s %+v\n", conn.RemoteAddr(), err)
			writeLock.Unlock()
			return
		}
		writeLock.Unlock()
		log.Printf("[%s] %s\n", conn.RemoteAddr(), scanner.Text())
	}
}