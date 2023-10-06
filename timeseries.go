package govantage

import (
	"fmt"
	"time"
)

type TimeSeries struct {
	MetaData *MetaData
	Records  []TimeSeriesRecord
}

type MetaData struct {
	Information   string
	Symbol        string
	LastRefreshed string
	Interval      string
	OutputSize    string
	TimeZone      string
}

type TimeSeriesRecord struct {
	TimeStamp        time.Time `json: "timestamp"`
	Open             float64   `json: "open"`
	High             float64   `json: "high"`
	Low              float64   `json: "low"`
	Close            float64   `json: "close"`
	Volume           int64     `json: "volume"`
	AdjustedClose    float64   `json: "adjusted_close"`
	DividendAmount   float64   `json: "dividend_close"`
	SplitCoefficient float64   `json: "split_coefficient"`
}

func (ts *TimeSeries) PrintMeta() {
	fmt.Printf("\n---------------\n")
	fmt.Printf("Symbol: %s\n", ts.MetaData.Symbol)
	fmt.Printf("Interval: %s\n", ts.MetaData.Interval)
	fmt.Printf("Output Size: %s\n", ts.MetaData.OutputSize)
	fmt.Printf("Timezone: %s\n", ts.MetaData.TimeZone)
	fmt.Printf("Last Refreshed: %s\n", ts.MetaData.LastRefreshed)
	fmt.Printf("---------------\n")
}

func (ts *TimeSeries) Open() []float64 {
	var open []float64
	for _, record := range ts.Records {
		open = append(open, record.Open)
	}
	return open
}

func (ts *TimeSeries) Close() []float64 {
	var close []float64
	for _, record := range ts.Records {
		close = append(close, record.Close)
	}
	return close
}

func (ts *TimeSeries) Low() []float64 {
	var low []float64
	for _, record := range ts.Records {
		low = append(low, record.Low)
	}
	return low
}

func (ts *TimeSeries) High() []float64 {
	var high []float64
	for _, record := range ts.Records {
		high = append(high, record.High)
	}
	return high
}

func (ts *TimeSeries) Volume() []int64 {
	var volume []int64
	for _, record := range ts.Records {
		volume = append(volume, record.Volume)
	}
	return volume
}

func (ts *TimeSeries) AdjustedClose() []float64 {
	var adjusted []float64
	for _, record := range ts.Records {
		adjusted = append(adjusted, record.AdjustedClose)
	}
	return adjusted
}

func (ts *TimeSeries) DividendAmount() []float64 {
	var dividends []float64
	for _, record := range ts.Records {
		dividends = append(dividends, record.DividendAmount)
	}
	return dividends
}
