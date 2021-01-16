package utils

import (
	"bufio"
	"log"
	"net"
)

/*PushImage pushes image to container registry.
  `repo` is the name of the image, e.g. (ahojukka5/anecdotes)
  `XRA` is base64 encoded X-Registry-Auth
*/
func PushImage(name string, XRA string) {
	addr := "/var/run/docker.sock"
	conn, err := net.Dial("unix", addr)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()

	l1 := "POST /v1.40/images/" + name + "/push?tag=latest HTTP/1.1"
	l2 := "Host: localhost"
	l3 := "User-Agent: curl/7.68.0"
	l4 := "Accept: */*"
	l5 := "Content-Type: application/json"
	l6 := "X-Registry-Auth: " + XRA
	l62 := "X-Registry-Auth: ***"
	msg := l1 + "\r\n" + l2 + "\r\n" + l3 + "\r\n" + l4 + "\r\n" + l5 + "\r\n" + l6 + "\r\n\r\n"
	msg2 := l1 + "\r\n" + l2 + "\r\n" + l3 + "\r\n" + l4 + "\r\n" + l5 + "\r\n" + l62 + "\r\n\r\n"

	println("## Send message to socket:")
	println("---")
	println(msg2)
	println("---")

	// Send message to socket
	if _, err := conn.Write([]byte(msg)); err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)

	println("Header:\n")
	// Read header
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		print(string(line))
		if string(line) == "\r\n" {
			break
		}
	}

	println("Push log:\n")
	// Read push log
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		print(string(line))
		if string(line) == "0\r\n" {
			break
		}
	}

	println("Image deployed to docker hub")

}
