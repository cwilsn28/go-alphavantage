package alphavantage

type StockAPI struct {
	Client *Client
}

/* Create a new CoreStock API */
func NewStockAPI(apiKey string) *StockAPI {
	client := NewClient(apiKey, "GET")
	return &StockAPI{Client: client}
}

/* TimeSeries methods */
func (api *StockAPI) IntraDay(queryParams QueryParams) (*IntraDayTimeSeries, error) {
	var err error
	var timeseries = IntraDayTimeSeries{}

	// Set the API function
	queryParams.Function = "TIME_SERIES_INTRADAY"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&timeseries)
	return &timeseries, err
}

func (api *StockAPI) Daily(queryParams QueryParams) (*DailyTimeSeries, error) {
	var err error
	var timeseries = DailyTimeSeries{}

	// Set the API function
	queryParams.Function = "TIME_SERIES_DAILY"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&timeseries)
	return &timeseries, err
}

func (api *StockAPI) DailyAdjusted(queryParams QueryParams) (*DailyAdjustedTimeSeries, error) {
	var err error
	var timeseries = DailyAdjustedTimeSeries{}

	// Set the API function
	queryParams.Function = "TIME_SERIES_DAILY_ADJUSTED"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&timeseries)
	return &timeseries, err
}

func (api *StockAPI) Weekly(queryParams QueryParams) (*WeeklyTimeSeries, error) {
	var err error
	var timeseries = WeeklyTimeSeries{}

	// Set the API function
	queryParams.Function = "TIME_SERIES_WEEKLY"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&timeseries)
	return &timeseries, err
}

func (api *StockAPI) WeeklyAdjusted(queryParams QueryParams) (*WeeklyAdjustedTimeSeries, error) {
	var err error
	var timeseries = WeeklyAdjustedTimeSeries{}

	// Set the API function
	queryParams.Function = "TIME_SERIES_WEEKLY_ADJUSTED"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&timeseries)
	return &timeseries, err
}

func (api *StockAPI) Monthly(queryParams QueryParams) (*MonthlyTimeSeries, error) {
	var err error
	var timeseries = MonthlyTimeSeries{}

	// Set the API function
	queryParams.Function = "TIME_SERIES_MONTHLY"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&timeseries)
	return &timeseries, err
}

func (api *StockAPI) MonthlyAdjusted(queryParams QueryParams) (*MonthlyAdjustedTimeSeries, error) {
	var err error
	var timeseries = MonthlyAdjustedTimeSeries{}

	// Set the API function
	queryParams.Function = "TIME_SERIES_MONTHLY_ADJUSTED"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&timeseries)
	return &timeseries, err
}

// func unmarshalDailyAdjusted(results map[string]interface{}) (*models.TimeSeries, error) {
// 	var err error
// 	var timeseries = models.TimeSeries{}

// 	// Parse the response metadata.
// 	meta := models.MetaData{
// 		Information:   results["Meta Data"].(map[string]interface{})["1. Information"].(string),
// 		Symbol:        results["Meta Data"].(map[string]interface{})["2. Symbol"].(string),
// 		LastRefreshed: results["Meta Data"].(map[string]interface{})["3. Last Refreshed"].(string),
// 		OutputSize:    results["Meta Data"].(map[string]interface{})["4. Output Size"].(string),
// 		TimeZone:      results["Meta Data"].(map[string]interface{})["5. Time Zone"].(string),
// 	}
// 	timeseries.MetaData = &meta

// 	// Parse and assemble the response records.
// 	records := results["Time Series (Daily)"]
// 	for k, v := range records.(map[string]interface{}) {
// 		record := models.TimeSeriesRecord{}
// 		record.TimeStamp, err = time.Parse("2006-01-02 15:04:05", k)
// 		data := v.(map[string]interface{})
// 		record.Open, err = strconv.ParseFloat(data["1. open"].(string), 64)
// 		record.High, err = strconv.ParseFloat(data["2. high"].(string), 64)
// 		record.Low, err = strconv.ParseFloat(data["3. low"].(string), 64)
// 		record.Close, err = strconv.ParseFloat(data["4. close"].(string), 64)
// 		record.AdjustedClose, err = strconv.ParseFloat(data["5. adjusted close"].(string), 64)
// 		record.Volume, err = strconv.ParseInt(data["6. volume"].(string), 10, 64)
// 		record.DividendAmount, err = strconv.ParseFloat(data["7. dividend amount"].(string), 64)
// 		record.SplitCoefficient, err = strconv.ParseFloat(data["8. split coefficient"].(string), 64)
// 		if err != nil {
// 			return &timeseries, err
// 		}
// 		timeseries.Records = append(timeseries.Records, record)
// 	}
// 	return &timeseries, err
// }

// func unmarshalWeekly(results map[string]interface{}) (*models.TimeSeries, error) {
// 	var err error
// 	var timeseries = models.TimeSeries{}

// 	// Parse the response metadata.
// 	meta := models.MetaData{
// 		Information:   results["Meta Data"].(map[string]interface{})["1. Information"].(string),
// 		Symbol:        results["Meta Data"].(map[string]interface{})["2. Symbol"].(string),
// 		LastRefreshed: results["Meta Data"].(map[string]interface{})["3. Last Refreshed"].(string),
// 		TimeZone:      results["Meta Data"].(map[string]interface{})["4. Time Zone"].(string),
// 	}
// 	timeseries.MetaData = &meta

// 	// Parse and assemble the reseponse records.
// 	records := results["Weekly Time Series"]
// 	for k, v := range records.(map[string]interface{}) {
// 		record := models.TimeSeriesRecord{}
// 		record.TimeStamp, err = time.Parse("2006-01-02 15:04:05", k)
// 		data := v.(map[string]interface{})
// 		record.Open, err = strconv.ParseFloat(data["1. open"].(string), 64)
// 		record.High, err = strconv.ParseFloat(data["2. high"].(string), 64)
// 		record.Low, err = strconv.ParseFloat(data["3. low"].(string), 64)
// 		record.Close, err = strconv.ParseFloat(data["4. close"].(string), 64)
// 		record.Volume, err = strconv.ParseInt(data["5. volume"].(string), 10, 64)
// 		if err != nil {
// 			return &timeseries, err
// 		}
// 		timeseries.Records = append(timeseries.Records, record)
// 	}
// 	return &timeseries, err
// }
