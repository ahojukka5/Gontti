package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"strings"
)

// Stream struct
type Stream struct {
	Stream string `json:"stream"`
}

// ID struct
type ID struct {
	ID string `json:"ID"`
}

// Aux struct
type Aux struct {
	Aux ID `json:"aux"`
}

// BuildImage builds and tags Docker image. Tag must be given as input argument,
// for example 'ahojukka5/anecdotes'. Returns id of image built.
func BuildImage(tag string) string {
	addr := "/var/run/docker.sock"
	conn, err := net.Dial("unix", addr)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()

	cmd := "http://localhost/v1.40/build?remote=https://github.com/" + tag + ".git#master:/"
	msg := "POST " + cmd + " HTTP/1.1\r\nHost: *\r\n\r\n"

	println("## Send message to socket:")
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

	println("\n## Header:\n")
	println(header)

	var stream Stream
	var aux Aux

	println("## Build output:\n")
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
		// println("number of characters in chunk", contentLength)

		if contentLength == 0 {
			println("\n## End of build output\n")
			break
		}

		buf := make([]byte, contentLength)
		n, err := reader.Read(buf[:])
		if err != nil {
			panic(err)
		}
		isStream := bytes.HasPrefix(buf[0:n], []byte("{\"stream\""))
		isAux := bytes.HasPrefix(buf[0:n], []byte("{\"aux\""))

		if isStream {
			json.Unmarshal(buf[0:n], &stream)
			// str := strings.Trim(stream.Stream, "\r\n")
			print(stream.Stream)
		}

		if isAux {
			json.Unmarshal(buf[0:n], &aux)
		}

		// Discard two bytes
		_, err = reader.Discard(2)
		if err != nil {
			panic(err)
		}
	}
	id := aux.Aux.ID
	println("Image id", id)
	return id
}
