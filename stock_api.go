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

func (api *StockAPI) GlobalQuote(queryParams QueryParams) (*GlobalQuote, error) {
	var err error
	var globalQuote = GlobalQuote{}

	// Set the API function
	queryParams.Function = "GLOBAL_QUOTE"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&globalQuote)
	return &globalQuote, err
}

func (api *StockAPI) TickerSearch(queryParams QueryParams) (*TickerSearchList, error) {
	var err error
	var searchList = TickerSearchList{}

	// Set the API function
	queryParams.Function = "SYMBOL_SEARCH"

	// Create and execute a new query.
	api.Client.NewRequest("GET")
	api.Client.NewQuery(queryParams)

	err = api.Client.ExecuteQuery(&searchList)
	return &searchList, err
}
