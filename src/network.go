// network.go

package main

import (
	"fmt"
	"io"
	"os"
	"net"
	"net/http"
	"bytes"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	servers := os.Args[1]

	conn, err := net.Dial("tcp", servers)
	checkError(err)

	_, err = conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	conn.Close()

	fmt.Println(string(result))

//	resp, err := http.Get("http://www.baidu.com/")
//	checkError(err)
//	defer resp.Body.Close()
//	io.Copy(os.Stdout, resp.Body)

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal %s\n", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return result.Bytes(), nil
}

