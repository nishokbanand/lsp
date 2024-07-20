package compiler

import (
	"fmt"
	"lsp/lsp"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri string, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri string, text string) {
	s.Documents[uri] = text
}

func (s *State) HoverResponse(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]
	return lsp.HoverResponse{
		Response: lsp.Response{RPC: "jsonrpc", ID: &id},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File %s, Characters %d", uri, len(document)),
		},
	}
}

func (s *State) DefintionResponse(id int, uri string, position lsp.Position) lsp.DefintionResponse {
	return lsp.DefintionResponse{
		Response: lsp.Response{RPC: "jsonrpc", ID: &id},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
