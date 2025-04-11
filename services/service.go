package services

import (
	"env-info/internal/clients"
	"env-info/internal/types"
	"fmt"
	"log"
)

func GetContractInfo(address string) (interface{}, error) {

	if address == "" {
		return nil, fmt.Errorf("address cannot be empty")
	}

	name, isToken, err := clients.FetchContractTags(address)
	if err != nil {
		log.Printf("Error fetching contract info: %v", err)
		return nil, fmt.Errorf("failed to fetch contract info: %v", err)
	}

	if isToken {
		symbol, err := FetchTokenInfo(address)
		if err != nil {
			log.Printf("Error fetching token info: %v", err)
			return nil, fmt.Errorf("failed to fetch token info: %v", err)
		}

		tokenInfo := types.TokenInfo{
			Address: address,
			Name:    name,
			Symbol:  symbol,
		}
		return tokenInfo, nil
	} else {
		category, ok := DappCategoryMap[name]
		if !ok {
			log.Printf("unknown dapp category for: %s", name)
			category = "unknown"
		}

		description, err := FetchDappInfo(name)
		if err != nil {
			log.Printf("Error fetching dapp info: %v", err)
			return nil, fmt.Errorf("failed to fetch dapp info: %v", err)
		}

		dappInfo := types.DappInfo{
			Address:     address,
			Name:        name,
			Category:    category,
			Description: description,
		}

		return dappInfo, nil
	}
}

func FetchTokenInfo(address string) (string, error) {
	symbol, err := clients.FetchDefillamaTokenInfo(address)
	if err != nil {
		log.Printf("Error fetching token info: %v", err)
		return "", fmt.Errorf("failed to fetch token info: %v", err)
	}

	return symbol, nil
}

func FetchDappInfo(name string) (string, error) {
	description, err := clients.FetchDefillamaDappInfo(name)
	if err != nil {
		log.Printf("Error fetching dapp info: %v", err)
		return "", fmt.Errorf("failed to fetch dapp info: %v", err)
	}

	return description, nil
}
