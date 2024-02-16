package main

import (
	"fmt"
	"net/http"
)

const (
	GLOSSARY                         string = "cmsContent?url=/glossary"
	HOLIDAY_TRADING                         = "holiday-master?type=trading"
	MARKET_DATA_PRE_OPEN                    = "market-data-pre-open?key=ALL"
	MERGED_DAILY_REPORTS_CAPITAL            = "merged-daily-reports?key=favCapital"
	MERGED_DAILY_REPORTS_DERIVATIVES        = "merged-daily-reports?key=favDerivatives"
	MERGED_DAILY_REPORTS_DEBT               = "merged-daily-reports?key=favDebt"
	HOLIDAY_CLEARING                        = "holiday-master?type=clearing"
	MARKET_STATUS                           = "marketStatus"
	MARKET_TURNOVER                         = "market-turnover"
	ALL_INDICES                             = "allIndices"
	INDEX_NAMES                             = "index-names"
	CIRCULARS                               = "circulars"
	LATEST_CIRCULARS                        = "latest-circular"
	EQUITY_MASTER                           = "equity-master"
)

func (api *NSEIndia) BuildUrl(uri string) string {
	return fmt.Sprintf("%s/api/%s", api.BaseURL, uri)
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

	return "drunk", nil
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
