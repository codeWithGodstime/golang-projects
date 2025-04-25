package core

type AppData struct {
	Collections []Collection `json:"collections"`
}

type Collection struct {
	Name     string    `json:"name"`
	Folders  []Folder  `json:"folders,omitempty"`
	Requests []Request `json:"requests,omitempty"` // top-level requests
}

type Folder struct {
	Name     string    `json:"name"`
	Requests []Request `json:"requests"`
}

type Request struct {
	Name    string            `json:"name"`
	Method  string            `json:"method"`  // GET, POST, etc.
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}
