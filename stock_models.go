package alphavantage

/*
Metadata fields in the AV response can vary depending on the query.
Overload the interval, outputsize, and timezone struct tags to capture
the variations in the json response.
*/
type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval,omitempty"`
	OutputSize    string `json:"4. Output Size,5. Output Size,omitempty"`
	TimeZone      string `json:"4. Time Zone,5. Time Zone,6. Time Zone,omitempty"`
}

/*
Overload the struct tags to account for presence of adjusted close,
dividend amount, and split coefficient.
*/
type OHLCV struct {
	Open             string `json:"1. open"`
	High             string `json:"2. high"`
	Low              string `json:"3. low"`
	Close            string `json:"4. close"`
	AdjustedClose    string `json:"5. adjusted close,omitempty"`
	Volume           string `json:"5. volume,6. volume"`
	DividendAmt      string `json:"7. dividend amount,omitempty"`
	SplitCoefficient string `json:"8. split coefficient,omitempty"`
}

type IntraDayTimeSeries struct {
	MetaData        MetaData         `json:"Meta Data"`
	TimeSeries1Min  map[string]OHLCV `json:"Time Series (1min),omitempty"`
	TimeSeries5Min  map[string]OHLCV `json:"Time Series (5min),omitempty"`
	TimeSeries15Min map[string]OHLCV `json:"Time Series (15min),omitempty"`
	TimeSeries30Min map[string]OHLCV `json:"Time Series (30min),omitempty"`
	TimeSeries60Min map[string]OHLCV `json:"Time Series (60min),omitempty"`
}

type DailyTimeSeries struct {
	MetaData   MetaData         `json:"Meta Data"`
	TimeSeries map[string]OHLCV `json:"Time Series (Daily),omitempty"`
}

type DailyAdjustedTimeSeries struct {
	MetaData   MetaData         `json:"Meta Data"`
	TimeSeries map[string]OHLCV `json:"Time Series (Daily),omitempty"`
}

type WeeklyTimeSeries struct {
	MetaData   MetaData         `json:"Meta Data"`
	TimeSeries map[string]OHLCV `json:"Weekly Time Series,omitempty"`
}

type WeeklyAdjustedTimeSeries struct {
	MetaData   MetaData         `json:"Meta Data"`
	TimeSeries map[string]OHLCV `json:"Weekly Adjusted Time Series,omitempty"`
}

type MonthlyTimeSeries struct {
	MetaData   MetaData         `json:"Meta Data"`
	TimeSeries map[string]OHLCV `json:"Monthly Time Series,omitempty"`
}

type MonthlyAdjustedTimeSeries struct {
	MetaData   MetaData         `json:"Meta Data"`
	TimeSeries map[string]OHLCV `json:"Monthly Adjusted Time Series,omitempty"`
}

/*
 * Models for stock quotes.
 */
type Quote struct {
	Symbol         string `json:"01. symbol"`
	Open           string `json:"02. open"`
	High           string `json:"03. high"`
	Low            string `json:"04. low"`
	Price          string `json:"05. price"`
	Volume         string `json:"06. volume"`
	LatestTradeDay string `json:"07. latest trading day"`
	PreviousClose  string `json:"08. previous close"`
	Change         string `json:"09. change"`
	ChangePercent  string `json:"10. change percent"`
}

type GlobalQuote struct {
	Quote Quote `json:"Global Quote"`
}
