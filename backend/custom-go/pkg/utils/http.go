package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
)

func CopyAndBindRequestBody(request *http.Request, result any) (err error) {
	dumpBytes, err := httputil.DumpRequest(request, true)
	if err != nil {
		return err
	}

	requestCopy, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(dumpBytes)))
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(requestCopy.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, &result)
}

func GetIp4() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
	}
	return
}
