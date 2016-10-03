package main

import (
	"MSMQReader/lib"
	"flag"
	"fmt"
)

func main() {

	var journal bool
	var server string
	var queuename string

	flag.StringVar(&server, "server", "", "server name to be connected")
	flag.StringVar(&queuename, "queuename", "", "queue name to found on the server")
	flag.BoolVar(&journal, "journal", false, "Reads the messages from journal")
	flag.Parse()
	fmt.Println(server, queuename)

	var list []lib.QueueMsg
	if journal {
		list = lib.ReadQueue(server, "journal", queuename)
	} else {
		list = lib.ReadQueue(server, "queue", queuename)
	}

	for _, val := range list {
		fmt.Println(val.Body)
	}

}
