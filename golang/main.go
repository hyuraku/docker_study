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

	log.Fatal(ssh.ListenAndServe(":2224", nil))
}
