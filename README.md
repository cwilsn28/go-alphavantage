# go-alphavantage

A Go wrapper for stock data and stock indicators from the Alpha Vantage API.

**Please Note:** This project is a work in progress.

## Introduction

Alpha Vantage provides a free API for real time financial data and financial
indicators. go-alphavantage provides a wrapper to the Alpha Vantage
API(http://www.alphavantage.co/). All methods require and API key which can be
requested at http://www.alphavantage.co/support/#api-key. Some API calls will
require a premium API key. Full Alpha Vantage API documentation is available at
http://www.alphavantage.co/documentation.

## Implemented Methods

### Time Series Stock Data APIs (WIP)
- [x] Intraday
- [ ] Intraday Extended History
- [x] Daily
- [x] Daily Adjusted
- [x] Weekly
- [x] Weekly Adjusted
- [x] Monthly
- [x] Monthly Adjusted
- [ ] Quote
- [ ] Search

### Alpha Intelligence (Not Implemented)

### Fundamental Data (Not Implemented)

### Forex(FX) (Not Implemented)

### Cryptocurrencies (Not Implemented)

### Economic Indicators (Not Implemented)

### Technical Indicators (Not Implemented)

## Usage

```bash
go get github.com/cwilsn28/go-alphavantage
```

Add your API token to a .env file in the root of your project directory
```bash
ALPHAVANTAGE_TOKEN=<your_api_token>
```

```go
token, err := alphavantage.GetAPIToken()
if err != nil {
    panic(err)
}

stockAPI := alphavantage.NewStockAPI(token)

/* Test building of query using parameters from Alphavantage API documentation. */
queryParams := alphavantage.QueryParams{
    Symbol:   "IBM",
    Interval: "5min",
}

/* Fetch daily timeseries */
results, err := stockAPI.IntraDay(queryParams)
if err != nil {
    panic(err)
}

...
```