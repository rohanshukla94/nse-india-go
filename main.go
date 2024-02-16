package main

import (
	"fmt"
	"net/http"
)

const (
	GLOSSARY                         string = "/api/cmsContent?url=/glossary"
	HOLIDAY_TRADING                         = "/api/holiday-master?type=trading"
	MARKET_DATA_PRE_OPEN                    = "/api/market-data-pre-open?key=ALL"
	MERGED_DAILY_REPORTS_CAPITAL            = "/api/merged-daily-reports?key=favCapital"
	MERGED_DAILY_REPORTS_DERIVATIVES        = "/api/merged-daily-reports?key=favDerivatives"
	MERGED_DAILY_REPORTS_DEBT               = "/api/merged-daily-reports?key=favDebt"
	HOLIDAY_CLEARING                        = "/api/holiday-master?type=clearing"
	MARKET_STATUS                           = "/api/marketStatus"
	MARKET_TURNOVER                         = "/api/market-turnover"
	ALL_INDICES                             = "/api/allIndices"
	INDEX_NAMES                             = "/api/index-names"
	CIRCULARS                               = "/api/circulars"
	LATEST_CIRCULARS                        = "/api/latest-circular"
	EQUITY_MASTER                           = "/api/equity-master"
)

func (api *NSEIndia) BuildUrl(uri string) string {
	return fmt.Sprintf("%s/%d", api.BaseURL, uri)
}

// GetSymbols retrieves a post by its ID
func (api *NSEIndia) GetSymbols() (string, error) {
	url := api.BuildUrl(MARKET_DATA_PRE_OPEN)
	resp, err := http.Get(url)
	if err != nil {
		return "sorrry", err
	}
	defer resp.Body.Close()

	//todo
	//very much drunk sir

	return drunk,
}

func main() {
	api := NewNSEIndia()

	symbols, err := api.GetSymbols()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("NSE data #%d\n", symbols)
}
