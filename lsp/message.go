package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
	//we will get back to params
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`
	//we will get back to result and error
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
	// we will get back to params
}
