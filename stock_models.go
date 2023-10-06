package govantage

type IntraDayMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

type DailyMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type WeeklyMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

type DailyOHLCV struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type DailyAdjustedOHLCV struct {
	Open             string `json:"1. open"`
	High             string `json:"2. high"`
	Low              string `json:"3. low"`
	Close            string `json:"4. close"`
	AdjustedClose    string `json:"5. adjusted close"`
	Volume           string `json:"6. volume"`
	DividendAmt      string `json:"7. dividend amount"`
	SplitCoefficient string `json:"8. split coefficient,omitempty"`
}

type DailyAdjustedMetaData DailyMetaData
type WeeklyAdjustedMetaData WeeklyMetaData
type WeeklyOHLCV DailyOHLCV
type WeeklyAdjustedOHLCV DailyAdjustedOHLCV
type MonthlyMetaData WeeklyMetaData
type MonthlyOHLCV DailyOHLCV
type MonthlyAdjustedMetaData WeeklyMetaData
type MonthlyAdjustedOHLCV DailyAdjustedOHLCV

type IntraDayTimeSeries struct {
	MetaData        IntraDayMetaData      `json:"Meta Data"`
	TimeSeries1Min  map[string]DailyOHLCV `json:"Time Series (1min),omitempty"`
	TimeSeries5Min  map[string]DailyOHLCV `json:"Time Series (5min),omitempty"`
	TimeSeries15Min map[string]DailyOHLCV `json:"Time Series (15min),omitempty"`
	TimeSeries30Min map[string]DailyOHLCV `json:"Time Series (30min),omitempty"`
	TimeSeries60Min map[string]DailyOHLCV `json:"Time Series (60min),omitempty"`
}

type DailyTimeSeries struct {
	MetaData   DailyMetaData         `json:"Meta Data"`
	TimeSeries map[string]DailyOHLCV `json:"Time Series (Daily),omitempty"`
}

type DailyAdjustedTimeSeries struct {
	MetaData   DailyAdjustedMetaData         `json:"Meta Data"`
	TimeSeries map[string]DailyAdjustedOHLCV `json:"Time Series (Daily),omitempty"`
}

type WeeklyTimeSeries struct {
	MetaData   WeeklyMetaData                 `json:"Meta Data"`
	TimeSeries map[string]WeeklyAdjustedOHLCV `json:"Weekly Time Series,omitempty"`
}

type WeeklyAdjustedTimeSeries struct {
	MetaData   WeeklyAdjustedMetaData         `json:"Meta Data"`
	TimeSeries map[string]WeeklyAdjustedOHLCV `json:"Weekly Adjusted Time Series,omitempty"`
}

type MonthlyTimeSeries struct {
	MetaData   MonthlyMetaData         `json:"Meta Data"`
	TimeSeries map[string]MonthlyOHLCV `json:"Monthly Time Series,omitempty"`
}

type MonthlyAdjustedTimeSeries struct {
	MetaData   MonthlyAdjustedMetaData         `json:"Meta Data"`
	TimeSeries map[string]MonthlyAdjustedOHLCV `json:"Monthly Adjusted Time Series,omitempty"`
}
