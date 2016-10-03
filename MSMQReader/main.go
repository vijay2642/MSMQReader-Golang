package main

import (
	"MSMQReader/lib"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	var journal bool
	var server string
	var queuename string
	var search string
	var write string

	flag.StringVar(&server, "server", "", "server name to be connected")
	flag.StringVar(&queuename, "queuename", "", "queue name to found on the server")
	flag.StringVar(&search, "search", "", "search the msmq for messages which contains the string")
	flag.StringVar(&write, "write", "", "Specifies the directory where the file has to be written")
	flag.BoolVar(&journal, "journal", false, "Reads the messages from journal")

	flag.Parse()
	fmt.Println(server, queuename)

	var list []lib.QueueMsg
	if journal {
		list = lib.ReadQueue(server, "journal", queuename, search)
	} else {
		list = lib.ReadQueue(server, "queue", queuename, search)
	}

	// fileHandle, _ := os.Create("E:\\output.txt")
	if len(write) > 1 {
		fileHandle, _ := os.Create(write)
		writer := bufio.NewWriter(fileHandle)
		defer fileHandle.Close()
		for _, val := range list {
			fmt.Println(val.Body)
			fmt.Fprintln(writer, val.Body+"\n")
		}
		writer.Flush()
	} else {
		for _, val := range list {
			fmt.Println(val.Body)
		}
	}

}
