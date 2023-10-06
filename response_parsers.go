package govantage

import (
	"fmt"
	"strconv"
	"time"
)

/* Unmarshal JSON responses */

func UnmarshalIntradayResults(results map[string]interface{}) (*TimeSeries, error) {
	var err error
	var timeseries = TimeSeries{}

	// Parse the response metadata.
	meta := MetaData{
		Information:   results["Meta Data"].(map[string]interface{})["1. Information"].(string),
		Symbol:        results["Meta Data"].(map[string]interface{})["2. Symbol"].(string),
		LastRefreshed: results["Meta Data"].(map[string]interface{})["3. Last Refreshed"].(string),
		Interval:      results["Meta Data"].(map[string]interface{})["4. Interval"].(string),
		OutputSize:    results["Meta Data"].(map[string]interface{})["5. Output Size"].(string),
		TimeZone:      results["Meta Data"].(map[string]interface{})["6. Time Zone"].(string),
	}
	timeseries.MetaData = &meta

	// Parse and assemble the reseponse records.
	records := results[fmt.Sprintf("Time Series (%s)", meta.Interval)]
	for k, v := range records.(map[string]interface{}) {
		record := TimeSeriesRecord{}
		record.TimeStamp, err = time.Parse("2006-01-02 15:04:05", k)
		data := v.(map[string]interface{})
		record.Open, err = strconv.ParseFloat(data["1. open"].(string), 64)
		record.High, err = strconv.ParseFloat(data["2. high"].(string), 64)
		record.Low, err = strconv.ParseFloat(data["3. low"].(string), 64)
		record.Close, err = strconv.ParseFloat(data["4. close"].(string), 64)
		record.Volume, err = strconv.ParseInt(data["5. volume"].(string), 10, 64)
		if err != nil {
			return &timeseries, err
		}
		timeseries.Records = append(timeseries.Records, record)
	}
	return &timeseries, err
}

func UnmarshalDailyResults(results map[string]interface{}) (*TimeSeries, error) {
	var err error
	var timeseries = TimeSeries{}

	// Parse the response metadata.
	meta := MetaData{
		Information:   results["Meta Data"].(map[string]interface{})["1. Information"].(string),
		Symbol:        results["Meta Data"].(map[string]interface{})["2. Symbol"].(string),
		LastRefreshed: results["Meta Data"].(map[string]interface{})["3. Last Refreshed"].(string),
		OutputSize:    results["Meta Data"].(map[string]interface{})["4. Output Size"].(string),
		TimeZone:      results["Meta Data"].(map[string]interface{})["5. Time Zone"].(string),
	}

	timeseries.MetaData = &meta

	// Parse and assemble the response records.
	records := results["Time Series (Daily)"]
	for k, v := range records.(map[string]interface{}) {
		record := TimeSeriesRecord{}
		record.TimeStamp, err = time.Parse("2006-01-02 15:04:05", k)
		data := v.(map[string]interface{})
		record.Open, err = strconv.ParseFloat(data["1. open"].(string), 64)
		record.High, err = strconv.ParseFloat(data["2. high"].(string), 64)
		record.Low, err = strconv.ParseFloat(data["3. low"].(string), 64)
		record.Close, err = strconv.ParseFloat(data["4. close"].(string), 64)
		record.Volume, err = strconv.ParseInt(data["5. volume"].(string), 10, 64)
		if err != nil {
			return &timeseries, err
		}
		timeseries.Records = append(timeseries.Records, record)
	}
	return &timeseries, err
}
