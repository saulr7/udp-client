package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const ADDRS = ":2525"

func main() {

	c, err := net.Dial("tcp", ADDRS)

	if err != nil {
		fmt.Println(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		msg, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->", msg)

		if strings.TrimSpace(text) == "q" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

}
