package alphavantage

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	STOCK_SYMBOL = "MSFT"
	INTERVAL     = "5min"
)

/* --------------------------------------
 * Only testing the non-premium endpoints
 * ----------------------------------- */
func TestTimeSeriesIntraDay(t *testing.T) {
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
	timeseries, err := stockAPI.IntraDay(queryParams)
	if err != nil {
		t.Fatalf("StockAPI Error: %s", err)
	}
	assert.Equal(t, timeseries.MetaData.Symbol, STOCK_SYMBOL)
}

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

func TestGlobalQuote(t *testing.T) {
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
	quote, err := stockAPI.GlobalQuote(queryParams)
	if err != nil {
		t.Fatalf("StockAPI Error: %s", err)
	}
	assert.Equal(t, quote.Quote.Symbol, STOCK_SYMBOL)
}

func TestSymbolSearch(t *testing.T) {
	token, err := GetAPIToken()
	if err != nil {
		t.Fatalf("Token Error: %s", err)
	}

	stockAPI := NewStockAPI(token)

	/* Test building of query using parameters from Alphavantage API documentation. */
	queryParams := QueryParams{
		Keywords: "IBM",
	}

	/* Fetch daily timeseries */
	results, err := stockAPI.TickerSearch(queryParams)
	if err != nil {
		t.Fatalf("StockAPI Error: %s", err)
	}
	assert.GreaterOrEqual(t, len(results.BestMatches), 0)
}

/* -----------------------------------------------------------------------------
 * Import/Export test functions use data stored in local .json/.csv files.
 * This saves on alphavantage.co rate/usage limits.
 * -------------------------------------------------------------------------- */
func TestExportToCSV(t *testing.T) {
	testFilename := "export_to_csv_test.csv"
	timeseries := DailyTimeSeries{}

	// Read timeseries data from .json
	fmt.Println("Importing daily timeseries data from json")
	err := TimeSeriesFromJSON("test_data/ibm_daily.json", &timeseries)
	if err != nil {
		t.Fatalf("TimeSeriesFromJSON Error: %s", err)
	}

	// Write timeseries data to .csv
	fmt.Println("Exporting daily timeseries data to csv")
	err = TimeSeriesToCSV(timeseries.TimeSeries, testFilename)
	if err != nil {
		t.Fatalf("CSV Export Error: %s", err)
	}

	// Make sure the .csv exists in the system.
	_, err = os.Open(testFilename)
	if os.IsNotExist(err) {
		t.Fatalf("CSV Export Error: %s", err)
	}

	os.Remove(testFilename)
}
