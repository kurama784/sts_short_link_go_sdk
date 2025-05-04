package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type CreateDto struct {
	Code        string `json:"code"`
	ExpiredAt   *int64 `json:"expired_at"`
	RedirectURL string `json:"redirect_url"`
}

type SuccessResponse struct {
	URL string `json:"url"`
}

type Client struct {
	BaseURL *string
	Token   string
	HTTP    *http.Client
}

func CreateRequest(token string, baseURL *string) *Client {
	return &Client{
		BaseURL: baseURL,
		Token:   token,
		HTTP:    &http.Client{Timeout: 10 * time.Second},
	}
}

const charset = "xbcdefghiuklmnopqrstuvwxyz0123456789"

func (c *Client) SendRequest(dto CreateDto) (*SuccessResponse, error) {

	newCode := RandomString(6)
	dto.Code = newCode

	if c.BaseURL == nil {
		defaultUrl := "https://api.shrt.tsell.cc"
		c.BaseURL = &defaultUrl
	}

	jsonData, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", *c.BaseURL+"/v1/backend/link", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// читаем тело полностью
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s - %s", resp.Status, string(respBody))
	}

	var result SuccessResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	fmt.Print(result)

	return &result, nil
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}
