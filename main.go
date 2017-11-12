package main

import (
	"cloudgo/service"
	"os"

	flag "github.com/spf13/pflag"
)

const (
	//PORT default:8080
	PORT string = "8080"
)

func main() {
	//get port
	portnum := os.Getenv("PORT")
	if len(portnum) == 0 {
		portnum = PORT
	}
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		portnum = *pPort
	}
	service.NewServer(portnum)
}
