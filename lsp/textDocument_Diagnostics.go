package lsp

type PublishDiagnosticsNotification struct {
	Notification
	Params PublishDiagnosticParams `json:"params"`
}
type PublishDiagnosticParams struct {
	URI         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Range    Range  `json:"range"`
	Severity int    `json:"severity"`
	Source   string `json:"source"`
	Message  string `json:"message"`
}
