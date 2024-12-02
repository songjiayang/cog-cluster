package cog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/songjiayang/cog-cluster/pkg/logger"
	"github.com/songjiayang/cog-cluster/pkg/util"
	"go.uber.org/zap"
)

var (
	client *Client
)

type Client struct {
	cogServerAddr string
	apiServerAddr string
	httpClient    *http.Client
}

func NewClient() *Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	return &Client{
		cogServerAddr: util.GetEnvOr("COG_SERVER_ADDR", "http://localhost:5000"),
		apiServerAddr: util.GetEnvOr("API_SERVER_ADDR", "http://localhost:8000"),
		httpClient:    &http.Client{Transport: tr},
	}
}

func GetClient() *Client {
	sync.OnceFunc(func() {
		client = NewClient()
	})()
	return client
}

func (c *Client) Predict(taskID string, payload []byte) error {
	var input Input
	if err := json.Unmarshal(payload, &input); err != nil {
		logger.Log().Error("unmarshal predict payload with error", zap.Error(err))
		return err
	}

	input.Webhook = fmt.Sprintf("%s/predictions/%s/callback", c.apiServerAddr, taskID)
	input.WebhookEventsFilter = []string{
		"start",
		"completed",
	}

	predictUrl := fmt.Sprintf("%s/predictions/%s", c.cogServerAddr, taskID)
	data, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPut, predictUrl, bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Prefer", "respond-async")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		logger.Log().Error("cog response with error", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	logger.Log().Info("cog response with data", zap.String("body", string(body)), zap.Int("httpCode", resp.StatusCode))

	return nil
}
