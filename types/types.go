package types

type RequestPayload struct {
	Token     string   `json:"token"`
	Features  []string `json:"features"`
	Restart   []string `json:"restart"`
	Reference string   `json:"ref"`
}
