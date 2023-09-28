package types

type Tx struct { // Let's keep it simple for now
	Message   string `json:"message"`
	Signature string `json:"signature"`
}
