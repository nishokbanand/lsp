package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"lsp/compiler"
	"lsp/lsp"
	"lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("./log.txt")
	logger.Println("Hey I have started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.SplitFunc)
	state := compiler.NewState()
	writer := os.Stdout
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.Decode(msg)
		if err != nil {
			logger.Fatalf("Got an error %s", err)
			continue
		}
		handlemsg(logger, writer, state, method, content)
	}
}

func writeResponse(write io.Writer, msg any) {
	writer := os.Stdout
	reply := rpc.Encode(msg)
	writer.Write([]byte(reply))
}

func handlemsg(logger *log.Logger, writer io.Writer, state compiler.State, method string, content []byte) {
	logger.Printf("Got the msg with Method %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Got an error while parsing this %s", err)
			return
		}
		logger.Printf("ClientName:%s , Version : %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)
		logger.Println("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println("textDocument/didOpen err: ", err)
			return
		}
		logger.Printf("URI:%s , Content : %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println("textDocument/didchange err", err)
			return
		}
		logger.Printf("URI:%s ", request.Params.TextDocument.URI)
		for _, changes := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, changes.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println("textDocument/didchange err", err)
			return
		}
		//write Response
		response := state.HoverResponse(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)
	case "textDocument/definition":
		var request lsp.DefintionRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println("textDocument/didchange err", err)
			return
		}
		//write Response
		response := state.DefintionResponse(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)
	}
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("The log-file is bad")
	}
	return log.New(logFile, "[lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
