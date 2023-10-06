package govantage

type QueryParams struct {
	// Required by all timeseries calls.
	Function string
	// Required by all timeseries calls EXCEPT search endpoint.
	Symbol string
	// Required by search endpoint.
	Keywords string
	// Required by intraday endpoint.
	Interval string
	// Optional for all daily timeseries calls.
	OutputSize string
	// Optional for all timeseries calls
	Datatype string
}

func (q QueryParams) HasInterval() bool {
	if q.Interval != "" {
		return true
	}
	return false
}

func (q QueryParams) HasOutputSize() bool {
	if q.OutputSize != "" {
		return true
	}
	return false
}

func (q QueryParams) HasDatatype() bool {
	if q.Datatype != "" {
		return true
	}
	return false
}
