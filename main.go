package main

import (
	"bufio"
	"encoding/json"
	"log"
	"lsp/lsp"
	"lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("./log.txt")
	logger.Println("Hey I have started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.SplitFunc)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.Decode(msg)
		if err != nil {
			logger.Fatalf("Got an error %s", err)
			continue
		}
		handlemsg(logger, method, content)
	}
}

func handlemsg(logger *log.Logger, method string, content []byte) {
	logger.Printf("Got the msg with Method %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Got an error while parsing this %s", err)
		}
		logger.Printf("ClientName:%s , Version : %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.Encode(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Println("Sent the reply")
	}
}

	func getLogger(filename string) *log.Logger {
		logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			panic("The log-file is bad")
		}
		return log.New(logFile, "[lsp]", log.Ldate|log.Ltime|log.Lshortfile)
	}
