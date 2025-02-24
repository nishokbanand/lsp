package lsp

type CodeActionRequest struct {
	Request
	Params CodeActionParams `json:"params"`
}
type CodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
	Context      CodeActionContext      `json:"context"`
}

type CodeActionContext struct{}

type CodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

type CodeAction struct {
	Title   string         `json:"title"`
	Edit    *WorkSpaceEdit `json:"edit"`
	Command *Command       `json:"command"`
}

type Command struct {
	Title     string      `json:"title"`
	Command   string      `json:"command"`
	Arguments interface{} `json:"arguments"`
}
