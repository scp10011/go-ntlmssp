package ntlmssp

import (
	"encoding/base64"
	"strings"
)

type authheader string

func (h authheader) IsBasic() bool {
	return strings.Index(string(h), "Basic") != -1
}

func (h authheader) IsNegotiate() bool {
	return strings.Index(string(h), "Negotiate") != -1
}

func (h authheader) IsNTLM() bool {
	return strings.Index(string(h), "NTLM") != -1
}

func (h authheader) GetData() ([]byte, error) {
	p := strings.Split(string(h), " ")
	if len(p) < 2 {
		return nil, nil
	}
	return base64.StdEncoding.DecodeString(string(p[1]))
}

func (h authheader) GetBasicCreds() (username, password string, err error) {
	d, err := h.GetData()
	if err != nil {
		return "", "", err
	}
	parts := strings.SplitN(string(d), ":", 2)
	return parts[0], parts[1], nil
}
