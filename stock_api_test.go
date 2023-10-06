package govantage

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

const (
	STOCK_SYMBOL = "IBM"
	INTERVAL     = "5min"
)

/* --------------------------------------
 * Only testing the non-premium endpoints
 * ----------------------------------- */

func TestTimeSeriesDaily(t *testing.T) {
	token, err := GetAPIToken()
	if err != nil {
		t.Fatalf("Token Error: %s", err)
	}

	stockAPI := NewStockAPI(token)

	/* Test building of query using parameters from Alphavantage API documentation. */
	queryParams := QueryParams{
		Symbol:   STOCK_SYMBOL,
		Interval: INTERVAL,
	}

	/* Fetch daily timeseries */
	timeseries, err := stockAPI.Daily(queryParams)
	if err != nil {
		t.Fatalf("StockAPI Error: %s", err)
	}
	assert.Equal(t, timeseries.MetaData.Symbol, STOCK_SYMBOL)
}

func TestTimeSeriesWeekly(t *testing.T) {
	token, err := GetAPIToken()
	if err != nil {
		t.Fatalf("Token Error: %s", err)
	}

	stockAPI := NewStockAPI(token)

	/* Test building of query using parameters from Alphavantage API documentation. */
	queryParams := QueryParams{
		Symbol:   STOCK_SYMBOL,
		Interval: INTERVAL,
	}

	/* Fetch daily timeseries */
	timeseries, err := stockAPI.Weekly(queryParams)
	if err != nil {
		t.Fatalf("StockAPI Error: %s", err)
	}
	assert.Equal(t, timeseries.MetaData.Symbol, STOCK_SYMBOL)
}

func TestTimeSeriesMonthly(t *testing.T) {
	token, err := GetAPIToken()
	if err != nil {
		t.Fatalf("Token Error: %s", err)
	}

	stockAPI := NewStockAPI(token)

	/* Test building of query using parameters from Alphavantage API documentation. */
	queryParams := QueryParams{
		Symbol:   STOCK_SYMBOL,
		Interval: INTERVAL,
	}

	/* Fetch daily timeseries */
	timeseries, err := stockAPI.Monthly(queryParams)
	if err != nil {
		t.Fatalf("StockAPI Error: %s", err)
	}
	assert.Equal(t, timeseries.MetaData.Symbol, STOCK_SYMBOL)
}
