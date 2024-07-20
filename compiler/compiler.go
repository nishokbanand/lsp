package compiler

import (
	"fmt"
	"lsp/lsp"
	"strings"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func getDocumentDiagnostics(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "VS Code") {
			idx := strings.Index(line, "VS Code")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("VS Code")),
				Severity: 1,
				Source:   "My lsp",
				Message:  "Use Neovim",
			})
		}
		if strings.Contains(line, "Neovim") {
			idx := strings.Index(line, "Neovim")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("Neovim")),
				Severity: 3,
				Source:   "My lsp",
				Message:  "Noice",
			})
		}
	}
	return diagnostics
}

func (s *State) OpenDocument(uri string, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	diagnostics := getDocumentDiagnostics(text)
	return diagnostics
}

func (s *State) UpdateDocument(uri string, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDocumentDiagnostics(text)
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

func (s *State) CodeActionResponse(id int, uri string) lsp.CodeActionResponse {
	text := s.Documents[uri]
	codeActions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		index := strings.Index(line, "VS Code")
		if index >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, index, index+len("VS Code")),
					NewText: "Vs C*ode",
				},
			}
			codeActions = append(codeActions, lsp.CodeAction{
				Title: "Replace VS Code with VS C*ode",
				Edit: &lsp.WorkSpaceEdit{
					Changes: replaceChange,
				},
			})

			updateChange := map[string][]lsp.TextEdit{}
			updateChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, index, index+len("VS Code")),
					NewText: "Neovim",
				},
			}
			codeActions = append(codeActions, lsp.CodeAction{
				Title: "Update with cool editor",
				Edit: &lsp.WorkSpaceEdit{
					Changes: updateChange,
				},
			})
		}
	}
	return lsp.CodeActionResponse{
		Response: lsp.Response{RPC: "jsonrpc", ID: &id},
		Result:   codeActions,
	}
}

func (s *State) CodeCompletionResponse(id int, uri string) lsp.CodeCompletionResponse {
	items := []lsp.CompletionItem{}
	items = append(items, lsp.CompletionItem{
		Label:         "fmt.Printf",
		Detail:        "Prints the statement to Stdout",
		Documentation: "Documentation",
	},
	)
	return lsp.CodeCompletionResponse{
		Response: lsp.Response{RPC: "jsonrpc", ID: &id},
		Result:   items,
	}
}

func LineRange(line int, start int, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}

}
