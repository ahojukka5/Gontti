package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"
)

func reader(r io.Reader) {
	buf := make([]byte, 32)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		print(string(buf[0:n]))
	}
}

func main2() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	println(resp)
	println(body)
}

func getVersion() {
	addr := "/var/run/docker.sock"
	conn, err := net.Dial("unix", addr)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()

	// go reader(conn)

	// cmd := "http://localhost/v1.40/build?remote=https://github.com/ahojukka5/anecdotes.git#master:/"
	// msg := "POST " + cmd + " HTTP/1.1\r\nHost: *\r\n\r\n"
	msg := "GET http://localhost/version HTTP/1.1\r\nHost: *\r\n\r\n"

	println("Send message to Docker Engine")
	println("---")
	println(msg)
	println("---")

	if _, err := conn.Write([]byte(msg)); err != nil {
		panic(err)
	}

	// buf := make([]byte, 32)
	println("Reading")
	println("---")
	reader := bufio.NewReader(conn)

	// logEntry := "Content-Encoding: gzip\r\nLast-Modified: Tue, 20 Aug 2013 15:45:41 GMT\r\nServer: nginx/0.8.54\r\nAge: 18884\r\nVary: Accept-Encoding\r\nContent-Type: text/html\r\nCache-Control: max-age=864000, public\r\nX-UA-Compatible: IE=Edge,chrome=1\r\nTiming-Allow-Origin: *\r\nContent-Length: 14888\r\nExpires: Mon, 31 Mar 2014 06:45:15 GMT\r\n"

	// we need to make sure to add a fake HTTP header here to make a valid request.
	// reader := bufio.NewReader(strings.NewReader("GET / HTTP/1.1\r\n" + logEntry + "\r\n"))

	// logReq, err := http.ReadRequest(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(logReq.Header)

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
	println("Header:")
	println(header)

	// Parse header
	headerReader := bufio.NewReader(strings.NewReader(header + "\r\n"))
	tp := textproto.NewReader(headerReader)

	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		log.Fatal(err)
	}

	// http.Header and textproto.MIMEHeader are both just a map[string][]string
	httpHeader := http.Header(mimeHeader)
	log.Println(httpHeader)

	contentLength, err := strconv.Atoi(httpHeader["Content-Length"][0])
	if err != nil {
		panic(err)
	}
	println("Content length:", contentLength)
	// Read content
	buf := make([]byte, contentLength)
	n, err := reader.Read(buf[:])
	if err != nil {
		panic(err)
	}
	print(string(buf[0:n]))

	// var content string
	// for {
	// 	line, err := reader.ReadBytes('\n')
	// 	if err != nil {
	// 		break
	// 	}
	// 	print(err)
	// 	print(string(line))
	// 	content += string(line)
	// }
	// println("Content:")
	// println(content)

	// if n, err := conn.Read(buf); err != nil {
	// 	println("Unable to read")
	// 	panic(err)
	// } else {
	// 	print(string(buf[0:n]))
	// }
	println("\n---")
}

func main() {
	getVersion()
}
