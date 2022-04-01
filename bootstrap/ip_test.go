package bootstrap

import (
	"io"
	"net/http"
	"testing"
)

func TestIp(t *testing.T) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.ip.sb/ip", nil)

	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36 Edg/99.0.1150.55")

	resp, err := client.Do(request)
	if err != nil {
		return
	}
	bts, _ := io.ReadAll(resp.Body)
	t.Log(string(bts))
}
