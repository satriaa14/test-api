package logging

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"

	"net/http"
	"strings"
)

const (
	Token = "USER-TOKEN: [%s]"

	PathURL = "PATH: [%s]"

	Method = "METHOD: [%s]"

	IP = "IP: [%s]"

	Headers = "HEADERS: [%+v]"

	Content = "CONTENT: [%s]"
)

type LogUserRequest struct {
	Token, PathURL, Method, IPAddress, Headers, Content, Info string
}

func getIPClient(w http.ResponseWriter, r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return "127.0.0.1"
}

func LogRequestClient(w http.ResponseWriter, r *http.Request) LogUserRequest {
	if r.URL.Path != "/" {
		var body []byte
		if r.Body != nil {
			body, _ = ioutil.ReadAll(r.Body)
			r.Body = ioutil.NopCloser(bytes.NewReader(body))
		}
		token := r.Header.Get("Authorization")
		logResp := LogUserRequest{
			Token:     token,
			PathURL:   r.URL.Path,
			Method:    r.Method,
			IPAddress: getIPClient(w, r),
			Headers:   fmt.Sprintf("%+v", r.Header),
		}
		return logResp
	}
	return LogUserRequest{}
}
