package objects

import "time"

// Base const type StockSerieType, StockCandlePeriod, StockSearch, StockSector, StockSymbolExchange
const (
	StockSerieTypeLine StockSerieType = "line"

	StockCandlePeriod1Min  StockCandlePeriod = "1min"
	StockCandlePeriod5Min  StockCandlePeriod = "5min"
	StockCandlePeriod15Min StockCandlePeriod = "15min"
	StockCandlePeriod30Min StockCandlePeriod = "30min"
	StockCandlePeriod1Hour StockCandlePeriod = "1hour"
	StockCandlePeriod4Hour StockCandlePeriod = "4hour"

	StockSearchNasdaq     StockSearch = "nasdaq"
	StockSearchNYSE       StockSearch = "nyse"
	StockSearchTSX        StockSearch = "tsx"
	StockSearchEuronext   StockSearch = "euronext"
	StockSearchMutualFund StockSearch = "mutual_fund"
	StockSearchETF        StockSearch = "etf"
	StockSearchAmex       StockSearch = "amex"
	StockSearchIndex      StockSearch = "index"
	StockSearchCommodity  StockSearch = "commodity"
	StockSearchForex      StockSearch = "forex"
	StockSearchCrypto     StockSearch = "crypto"

	IndexSP500     Index = "S&P 500"
	IndexDowJones  Index = "Dow Jones"
	IndexNasdaq100 Index = "Nasdaq 100"

	StockSectorBasicMaterials        StockSector = "Basic Materials"
	StockSectorBuilding              StockSector = "Building"
	StockSectorCommunicationServices StockSector = "Communication Services"
	StockSectorConglomerates         StockSector = "Conglomerates"
	StockSectorConsumerCyclical      StockSector = "Consumer Cyclical"
	StockSectorConsumerDefensive     StockSector = "Consumer Defensive"
	StockSectorEnergy                StockSector = "Energy"
	StockSectorFinancial             StockSector = "Financial"
	StockSectorFinancialServices     StockSector = "Financial Services"
	StockSectorHealthcare            StockSector = "Healthcare"
	StockSectorIndustrials           StockSector = "Industrials"
	StockSectorMedia                 StockSector = "Media"
	StockSectorPharmaceuticals       StockSector = "Pharmaceuticals"
	StockSectorRealEstate            StockSector = "Real Estate"
	StockSectorServices              StockSector = "Services"
	StockSectorTechnology            StockSector = "Technology"
	StockSectorUtilities             StockSector = "Utilities"

	StockSymbolNasdaq     StockSymbolExchange = "available-nasdaq"
	StockSymbolNYSE       StockSymbolExchange = "available-nyse"
	StockSymbolTSX        StockSymbolExchange = "available-tsx"
	StockSymbolEuronext   StockSymbolExchange = "available-euronext"
	StockSymbolMutualFund StockSymbolExchange = "available-mutual-funds"
	StockSymbolETF        StockSymbolExchange = "available-etfs"
	StockSymbolAmex       StockSymbolExchange = "available-amex"
	StockSymbolIndex      StockSymbolExchange = "available-indexes"
	StockSymbolCommodity  StockSymbolExchange = "available-commodities"
	StockSymbolForex      StockSymbolExchange = "available-forex"
	StockSymbolCrypto     StockSymbolExchange = "available-crypto"
	StockSymbolAll        StockSymbolExchange = "all"
)

// StockSymbolExchange ...
type StockSymbolExchange string

// StockSearch ...
type StockSearch string

// Index ...
type Index string

// StockSector ...
type StockSector string

// StockCandlePeriod ...
type StockCandlePeriod string

// StockSeriesType ...
type StockSeriesType string

// StockQuote ...
type StockQuote struct {
	Symbol               string   `json:"symbol"`
	Name                 string   `json:"name"`
	Price                float64  `json:"price"`
	ChangesPercentage    float64  `json:"changesPercentage"`
	Change               float64  `json:"change"`
	DayLow               float64  `json:"dayLow"`
	DayHigh              float64  `json:"dayHigh"`
	YearHigh             float64  `json:"yearHigh"`
	YearLow              float64  `json:"yearLow"`
	MarketCap            *float64 `json:"marketCap"`
	PriceAvg50           float64  `json:"priceAvg50"`
	PriceAvg200          float64  `json:"priceAvg200"`
	Volume               int64    `json:"volume"`
	AvgVolume            int64    `json:"avgVolume"`
	Exchange             string   `json:"exchange"`
	Open                 float64  `json:"open"`
	PreviousClose        float64  `json:"previousClose"`
	Eps                  *float64 `json:"eps"`
	Pe                   *float64 `json:"pe"`
	EarningsAnnouncement *string  `json:"earningsAnnouncement"` // 2020-07-22T16:09:24.000+0000
	SharesOutstanding    *int64   `json:"sharesOutstanding"`
	Timestamp            int64    `json:"timestamp"`
}

// StockQuoteShort ...
type StockQuoteShort struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Volume int64   `json:"volume"`
}

// RequestStockSearch ...
type RequestStockSearch struct {
	Exchange *StockSearch
	Query    string
	Limit    int64
}

// StockSymbol ...
type StockSymbol struct {
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Currency          *string `json:"currency"`
	StockExchange     string  `json:"stockExchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
}

// StockCompanyProfile ...
type StockCompanyProfile struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Beta              float64 `json:"beta"`
	VolAvg            int64   `json:"volAvg"`
	MktCap            int64   `json:"mktCap"`
	LastDiv           float64 `json:"lastDiv"`
	Range             string  `json:"range"`
	Changes           float64 `json:"changes"`
	CompanyName       string  `json:"companyName"`
	Currency          string  `json:"currency"`
	Cik               string  `json:"cik"`
	Isin              string  `json:"isin"`
	Cusip             string  `json:"cusip"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Industry          string  `json:"industry"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	Ceo               string  `json:"ceo"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	DcfDiff           float64 `json:"dcfDiff"`
	Dcf               float64 `json:"dcf"`
	Image             string  `json:"image"`
	IpoDate           string  `json:"ipoDate"` // 1980-12-12
	DefaultImage      bool    `json:"defaultImage"`
	IsEtf             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
}

// CompanyExecutive ...
type CompanyExecutive struct {
	Title       string  `json:"title"`
	Name        string  `json:"name"`
	Pay         *int    `json:"pay"`
	CurrencyPay string  `json:"currencyPay"`
	Gender      string  `json:"gender"`
	YearBorn    int     `json:"yearBorn"`
	TitleSince  *string `json:"titleSince"`
}

// StockSplit ...
type StockSplit struct {
	Symbol     string           `json:"symbol"`
	Historical []StockSplitInfo `json:"historical"`
}

// StockSplitInfo ...
type StockSplitInfo struct {
	Date        string  `json:"date"`  // "2014-06-09"
	Label       string  `json:"label"` // February 28, 05
	Numerator   float64 `json:"numerator"`
	Denominator float64 `json:"denominator"`
}

// StockDividends ...
type StockDividends struct {
	Symbol     string               `json:"symbol"`
	Historical []StockDividendsInfo `json:"historical"`
}

// StockDividendsInfo ...
type StockDividendsInfo struct {
	Date            string  `json:"date"`            // "2014-06-09"
	Label           string  `json:"label"`           // February 28, 05
	RecordDate      string  `json:"recordDate"`      // "2014-06-09"
	PaymentDate     string  `json:"paymentDate"`     // "2014-06-09"
	DeclarationDate string  `json:"declarationDate"` // "2014-06-09"
	AdjDividend     float64 `json:"adjDividend"`
	Dividend        float64 `json:"dividend"`
}

// RequestStockCandleList ...
type RequestStockCandleList struct {
	Period StockCandlePeriod
	Symbol string
	From   *time.Time
	To     *time.Time
}

// StockDailyLineList ...
type StockDailyLineList struct {
	Symbol     string           `json:"symbol"`
	Historical []StockDailyLine `json:"historical"`
}

// StockDailyLine ...
type StockDailyLine struct {
	Date  string  `json:"date"`
	Close float64 `json:"close"`
}

// StockDailyCandleList ...
type StockDailyCandleList struct {
	Symbol     string             `json:"symbol"`
	Historical []StockDailyCandle `json:"historical"`
}

// StockBatchDaily ...
type StockBatchDaily struct {
	Data []StockBatchData `json:"historicalStockList"`
}

// StockBatchData ...
type StockBatchData struct {
	Symbol     string             `json:"symbol"`
	Historical []StockDailyCandle `json:"historical"`
}

// StockDailyCandle ...
type StockDailyCandle struct {
	Date             string  `json:"date"` // 2019-03-11
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Close            float64 `json:"close"`
	AdjClose         float64 `json:"adjClose"`
	Volume           float64 `json:"volume"`
	UnadjustedVolume float64 `json:"unadjustedVolume"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
	Vwap             float64 `json:"vwap"`
	Label            string  `json:"label"` // March 12, 19
	ChangeOverTime   float64 `json:"changeOverTime"`
}

// StockCandle ...
type StockCandle struct {
	Date   string  `json:"date"` // 2020-09-14 07:27:00
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// StockSymbolList ...
type StockSymbolList struct {
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Exchange string  `json:"exchange"`
}

// IndexSymbol ...
type IndexSymbol struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Sector         string `json:"sector"`
	SubSector      string `json:"subSector"`
	HeadQuarter    string `json:"headQuarter"`
	DateFirstAdded string `json:"dateFirstAdded"` // (2016-09-06 || NaN)
	Cik            string `json:"cik"`
	Founded        string `json:"founded"`
}

// HistoryIndexSymbol ...
type HistoryIndexSymbol struct {
	DateAdded       string `json:"dateAdded"` // July 26, 2017
	AddedSecurity   string `json:"addedSecurity"`
	RemovedTicker   string `json:"removedTicker"`
	RemovedSecurity string `json:"removedSecurity"`
	Reason          string `json:"reason"`
	Symbol          string `json:"symbol"`
}

// StockEODCandle ...
type StockEODCandle struct {
	Symbol   string  `json:"symbol"`
	Date     string  `json:"date"`
	Open     float64 `json:"open"`
	Low      float64 `json:"low"`
	High     float64 `json:"high"`
	Close    float64 `json:"close"`
	AdjClose float64 `json:"adjClose"`
	Volume   int     `json:"volume"`
}

// Active ...
type Active struct {
	Ticker            string  `json:"ticker"`
	Changes           float64 `json:"changes"`
	Price             string  `json:"price"`
	ChangesPercentage string  `json:"changesPercentage"`
	CompanyName       string  `json:"companyName"`
}

// Loser ...
type Loser struct {
	Ticker            string  `json:"ticker"`
	Changes           float64 `json:"changes"`
	Price             string  `json:"price"`
	ChangesPercentage string  `json:"changesPercentage"`
	CompanyName       string  `json:"companyName"`
}

// Gainer ...
type Gainer struct {
	Ticker            string  `json:"ticker"`
	Changes           float64 `json:"changes"`
	Price             string  `json:"price"`
	ChangesPercentage string  `json:"changesPercentage"`
	CompanyName       string  `json:"companyName"`
}

// Sector ...
type Sector struct {
	Sector            StockSector `json:"sector"`
	ChangesPercentage string      `json:"changesPercentage"`
}

// HistorySector ...
type HistorySector struct {
	Date                                   string  `json:"date"` // 2020-09-14
	UtilitiesChangesPercentage             float64 `json:"utilitiesChangesPercentage"`
	BasicMaterialsChangesPercentage        float64 `json:"basicMaterialsChangesPercentage"`
	CommunicationServicesChangesPercentage float64 `json:"communicationServicesChangesPercentage"`
	ConglomeratesChangesPercentage         float64 `json:"conglomeratesChangesPercentage"`
	ConsumerCyclicalChangesPercentage      float64 `json:"consumerCyclicalChangesPercentage"`
	ConsumerDefensiveChangesPercentage     float64 `json:"consumerDefensiveChangesPercentage"`
	EnergyChangesPercentage                float64 `json:"energyChangesPercentage"`
	FinancialChangesPercentage             float64 `json:"financialChangesPercentage"`
	FinancialServicesChangesPercentage     float64 `json:"financialServicesChangesPercentage"`
	HealthcareChangesPercentage            float64 `json:"healthcareChangesPercentage"`
	IndustrialGoodsChangesPercentage       float64 `json:"industrialGoodsChangesPercentage"`
	IndustrialsChangesPercentage           float64 `json:"industrialsChangesPercentage"`
	RealEstateChangesPercentage            float64 `json:"realEstateChangesPercentage"`
	ServicesChangesPercentage              float64 `json:"servicesChangesPercentage"`
	TechnologyChangesPercentage            float64 `json:"technologyChangesPercentage"`
}

// SurvivorshipBiasFree ...
type SurvivorshipBiasFree struct {
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
	From   string  `json:"from"`
}

// StockSeriesType return string
func (s StockSeriesType) String() string {
	return string(s)
}

// StockSearch return string
func (s StockSearch) String() string {
	return string(s)
}

// StockSymbolExchange return string
func (s StockSymbolExchange) String() string {
	return string(s)
}
