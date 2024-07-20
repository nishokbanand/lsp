package lsp

type CodeCompletionRequest struct {
	Request
	Params TextDocumentPositionParams `json:"params"`
}

type CodeCompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label         string `json:"label"`
	Detail        string `json:"detail"`
	Documentation string `json:"documentation"`
}
