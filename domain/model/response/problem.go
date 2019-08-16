package response

// Problem is model for error response conforming to RFC 7807
// see https://tools.ietf.org/html/rfc7807
type Problem struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
}
