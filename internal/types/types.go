package types

type DappInfo struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type TokenInfo struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}
