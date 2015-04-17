package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"strconv"
)

func main() {
	remoteHost := flag.String("host", "", "Remote host to connect. Mandatory.")
	remotePort := flag.Int("port", 0, "Remote port to connect. Mandatory.")
	localInteface := flag.String("interface", "", "Local interface to bind to. Defaults to every interface available.")
	localPort := flag.Int("local-port", 8888, "Remote port to bind to. Defaults to port-1.")
	InsecureSkipVerify := flag.Bool("insecure-skip-verify", false, "Skip verification of server's certificate chain and host name")
	flag.Parse()

	if *remoteHost == "" {
		log.Fatal("-host is a mandatory parameter. Run again with -h for a list of arguments")
	}

	if *remotePort == 0 {
		log.Fatal("-port is a mandatory parameter. Run again with -h for a list of arguments")
	}

	ln, err := net.Listen("tcp", *localInteface+":"+strconv.Itoa(*localPort))
	if err != nil {
		log.Fatal(err)
	}

	for {
		TCPConn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		tlsConfig := tls.Config{}
		tlsConfig.InsecureSkipVerify = *InsecureSkipVerify

		TLSConn, err := tls.Dial("tcp", *remoteHost+":"+strconv.Itoa(*remotePort), &tlsConfig)
		if err != nil {
			log.Fatal(err)
		}

		go (func() {
			_, err = io.Copy(TLSConn, TCPConn)
			if err != nil {
				log.Fatal(err)
			}
		})()

		go (func() {
			_, err = io.Copy(TCPConn, TLSConn)
			if err != nil {
				log.Fatal(err)
			}
		})()
	}
}
