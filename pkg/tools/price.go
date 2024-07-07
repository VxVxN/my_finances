package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type MultiTokensResponse struct {
	Data []struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Address           string `json:"address"`
			Name              string `json:"name"`
			Symbol            string `json:"symbol"`
			CoingeckoCoinId   string `json:"coingecko_coin_id"`
			Decimals          int    `json:"decimals"`
			TotalSupply       string `json:"total_supply"`
			PriceUsd          string `json:"price_usd"`
			FdvUsd            string `json:"fdv_usd"`
			TotalReserveInUsd string `json:"total_reserve_in_usd"`
			VolumeUsd         struct {
				H24 string `json:"h24"`
			} `json:"volume_usd"`
			MarketCapUsd interface{} `json:"market_cap_usd"`
		} `json:"attributes"`
		Relationships struct {
			TopPools struct {
				Data []struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"top_pools"`
		} `json:"relationships"`
	} `json:"data"`
}

func FetchPrices(tokens []string) (map[string]float64, error) {
	joinedTokens := strings.Join(tokens, ",")
	geckoUrl := fmt.Sprintf("https://api.geckoterminal.com/api/v2/networks/eth/tokens/multi/%s", url.QueryEscape(joinedTokens))
	var body []byte
	resp, err := http.Get(geckoUrl)
	if err != nil {
		return nil, fmt.Errorf("%v, url: %s", err, geckoUrl)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &MultiTokensResponse{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	if len(result.Data) == 0 {
		return nil, fmt.Errorf("cannot get gecko tokens, body: %s", string(body))
	}
	prices := make(map[string]float64)
	for _, data := range result.Data {
		priceUsd, err := strconv.ParseFloat(data.Attributes.PriceUsd, 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse price: %v", err)
		}
		prices[data.Attributes.Symbol] = priceUsd
	}

	return prices, nil
}
