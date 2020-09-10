package main

import (
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "hello world\n")
	})
	log.Println("starting ssh server on port 2224...")
	log.Fatal(ssh.ListenAndServe(":2224", nil))
}