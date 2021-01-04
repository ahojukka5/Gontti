package utils

import (
	"bufio"
	"log"
	"net"
)

/*TagImage tags image with `name` with `repo`.
  Example: TagImage("9e7244cf5587", "ahojukka5/anecdotes")
*/
func TagImage(name string, repo string) {
	addr := "/var/run/docker.sock"
	conn, err := net.Dial("unix", addr)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()

	uri := "http://localhost/v1.40/images/" + name + "/tag?repo=" + repo
	msg := "POST " + uri + " HTTP/1.1\r\nHost: *\r\n\r\n"

	println("## Send message to socket:")
	println("---")
	println(msg)
	println("---")

	// Send message to socket
	if _, err := conn.Write([]byte(msg)); err != nil {
		panic(err)
	}

	println("Response:\n")

	reader := bufio.NewReader(conn)

	println("Header:\n")
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
}
