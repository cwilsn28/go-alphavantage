package alphavantage

import "time"

/* -----------------------------------------------------------------------------
 * Public constants.
 * -------------------------------------------------------------------------- */

const AVBaseURL = "https://alphavantage.co/query"
const DefaultHTTPTimeout = 80 * time.Second

/* ---
 * Timeseries API functions
 * --- */
var TSFUNCTIONS = map[string]string{
	"intraday":         "TIME_SERIES_INTRADAY",
	"daily":            "TIME_SERIES_DAILY",
	"daily_adjusted":   "TIME_SERIES_DAILY_ADJUSTED",
	"weekly":           "TIME_SERIES_WEEKLY",
	"weekly_adjusted":  "TIME_SERIES_WEEKLY_ADJUSTED",
	"monthly":          "TIME_SERIES_MONTHLY",
	"monthly_adjusted": "TIME_SERIES_MONTHLY_ADJUSTED",
	"quote":            "GLOBAL_QUOTES",
	"search":           "SYMBOL_SEARCH",
}

/* ---
 * Timeseries intervals
 * --- */
var TSINTERVALS = map[int]string{
	1:  "1min",
	5:  "5min",
	15: "15min",
	30: "30min",
	60: "60min",
}
