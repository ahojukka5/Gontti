package main

import (
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 32)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("Client got:", string(buf[0:n]))
	}
}

func main() {
	addr := "/var/run/docker.sock"
	conn, err := net.Dial("unix", addr)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()

	go reader(conn)

	msg := "GET http://localhost/version HTTP/1.1\r\nHost: *\r\n\r\n"
	println("Send message to Docker Engine")
	println("---")
	println(msg)
	println("---")

	if _, err := conn.Write([]byte(msg)); err != nil {
		panic(err)
	}

	// buf := make([]byte, 32)
	println("Read from socket:")
	println("---")

	// var buf bytes.Buffer
	// io.Copy(&buf, conn)
	// fmt.Println("total size:", buf.Len())

	// for {
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	println()
	// 	println("n", n)
	// 	println("err", err)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	print("\nClient got:", string(buf[0:n]))
	// }

	// if n, err := conn.Read(buf); err != nil {
	// 	println("Unable to read")
	// 	panic(err)
	// } else {
	// 	print(string(buf[0:n]))
	// }
	time.Sleep(1e9)
	println("\n---")
}
