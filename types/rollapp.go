package types

type Tx struct { // Let's keep it simple for now
	Message   Message `json:"message"`
	Signature string  `json:"signature"`
}

type Message struct {
	Action  string `json:"action"`
	Content string `json:"content"`
	Index   int    `json:"index"`
}
