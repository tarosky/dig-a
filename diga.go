package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Error: Illegal number of arguments")
	}

	domain := os.Args[1]
	addrs, err := net.LookupHost(domain)

	if err != nil {
		return
	}

	s := ""
	for _, addr := range addrs {
		s += fmt.Sprintf("%s 0 IN A %s\n", domain, addr)
	}

	fmt.Print(s)
}
