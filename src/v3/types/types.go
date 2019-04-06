package types

// Ping https://api.coingecko.com/api/v3/ping
type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

// SimpleSinglePrice https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd
type SimpleSinglePrice struct {
	ID          string
	Currency    string
	MarketPrice float32
}

// SimpleSupportedVSCurrencies https://api.coingecko.com/api/v3/simple/supported_vs_currencies
type SimpleSupportedVSCurrencies []string

// CoinList https://api.coingecko.com/api/v3/coins/list
type CoinList []CoinsListItem

// CoinsMarket

// CoinsId

// CoinsIDTickers

// CoinsIDHistory

// CoinsIDMarketChart

// CoinsIDStatusUpdates

// CoinsIDContractAddress https://api.coingecko.com/api/v3/coins/{id}/contract/{contract_address}
// type CoinsIDContractAddress struct {
// 	ID                  string           `json:"id"`
// 	Symbol              string           `json:"symbol"`
// 	Name                string           `json:"name"`
// 	BlockTimeInMin      uint16           `json:"block_time_in_minutes"`
// 	Categories          []string         `json:"categories"`
// 	Localization        LocalizationItem `json:"localization"`
// 	Description         DescriptionItem  `json:"description"`
// 	Links               LinksItem        `json:"links"`
// 	Image               ImageItem        `json:"image"`
// 	CountryOrigin       string           `json:"country_origin"`
// 	GenesisDate         string           `json:"genesis_date"`
// 	ContractAddress     string           `json:"contract_address"`
// 	MarketCapRank       uint16           `json:"market_cap_rank"`
// 	CoinGeckoRank       uint16           `json:"coingecko_rank"`
// 	CoinGeckoScore      float32          `json:"coingecko_score"`
// 	DeveloperScore      float32          `json:"developer_score"`
// 	CommunityScore      float32          `json:"community_score"`
// 	LiquidityScore      float32          `json:"liquidity_score"`
// 	PublicInterestScore float32          `json:"public_interest_score"`
// 	MarketData          `json:"market_data"`
// }

// EventsCountries https://api.coingecko.com/api/v3/events/countries
type EventsCountries struct {
	Data []EventCountryItem `json:"data"`
}

// EventsTypes https://api.coingecko.com/api/v3/events/types
type EventsTypes struct {
	Data  []string `json:"data"`
	Count uint16   `json:"count"`
}

// ExchangeRatesResponse https://api.coingecko.com/api/v3/exchange_rates
type ExchangeRatesResponse struct {
	Rates ExchangeRatesItem `json:"rates"`
}

// GlobalResponse https://api.coingecko.com/api/v3/global
type GlobalResponse struct {
	Data Global `json:"data"`
}
