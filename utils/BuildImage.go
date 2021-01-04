package utils

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

// BuildImage builds and tags Docker image.
func BuildImage() {
	addr := "/var/run/docker.sock"
	conn, err := net.Dial("unix", addr)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()

	cmd := "http://localhost/v1.40/build?remote=https://github.com/ahojukka5/anecdotes.git#master:/"
	msg := "POST " + cmd + " HTTP/1.1\r\nHost: *\r\n\r\n"

	println("Send message to socket:")
	println("---")
	println(msg)
	println("---")

	// Send message to socket
	if _, err := conn.Write([]byte(msg)); err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)

	// Read header
	var header string
	var firstLine bool = true
	for {
		line, err := reader.ReadBytes('\n')
		if firstLine {
			firstLine = false
			continue
		}
		if err != nil {
			break
		}
		if string(line) == "\r\n" {
			break
		}
		header += string(line)
	}

	println("\nHeader:")
	println(header)

	println("Chunks:")
	for {
		content, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		contentLengthHex := strings.TrimSuffix(string(content), "\r\n")
		contentLength, err := strconv.ParseInt(contentLengthHex, 16, 64)
		if err != nil {
			panic(err)
		}
		println("number of characters in chunk", contentLength)

		if contentLength == 0 {
			println("End of stream")
			return
		}

		buf := make([]byte, contentLength)
		n, err := reader.Read(buf[:])
		if err != nil {
			panic(err)
		}
		print(string(buf[0:n]))

		// Discard two bytes
		_, err = reader.Discard(2)
		if err != nil {
			panic(err)
		}
	}
}
