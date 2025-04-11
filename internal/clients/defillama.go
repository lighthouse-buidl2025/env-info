package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LlamaDappResponse struct {
	Description string `json:"description"`
}

type LlamaTokenResponse struct {
	Coins map[string]struct {
		Decimals   int     `json:"decimals"`
		Symbol     string  `json:"symbol"`
		Price      float64 `json:"price"`
		Timestamp  int64   `json:"timestamp"`
		Confidence float64 `json:"confidence"`
	} `json:"coins"`
}

func FetchDefillamaDappInfo(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}

	url := fmt.Sprintf("https://api.llama.fi/protocol/%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data LlamaDappResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data.Description, nil
}

func FetchDefillamaTokenInfo(address string) (string, error) {
	if address == "" {
		return "", errors.New("address is required")
	}

	address = strings.ToLower(address)
	url := fmt.Sprintf("https://coins.llama.fi/prices/current/ethereum:%s", address)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data LlamaTokenResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	key := fmt.Sprintf("ethereum:%s", address)
	coinInfo, exists := data.Coins[key]
	if !exists {
		return "", fmt.Errorf("symbol not found for address: %s", address)
	}

	return coinInfo.Symbol, nil
}
