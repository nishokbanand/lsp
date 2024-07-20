package lsp

type DefintionRequest struct {
	Request
	Params TextDocumentPositionParams `json:"params"`
}

type DefintionResponse struct {
	Response
	Result Location `json:"result"`
}
