package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
)

func main() {

	url := "https://api.dev.Aexample.com/core/v1/domains/{webapp4.com}"

	payload := strings.NewReader("{\n  \"autorenewEnabled\": true,\n  \"privacyEnabled\": true,\n  \"locked\": true\n}")

	req, _ := http.NewRequest("PATCH", url, payload)

	req.Header.Add("Authorization", "Basic QWV4YW1wbGUuY29tOk1BbmtpbmQ=")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
