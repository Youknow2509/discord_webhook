package executewebhook

import (
	"io"
	"net/http"
	"time"
)

// excute webhook to discord
func ExcuteWebhook(urlWebhook string, body io.Reader) (res *http.Response, err error) {
	req, err := http.NewRequest(
		http.MethodPost,
		urlWebhook,
		body,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return r, nil
}
