package alphavantage

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

/* ---------------------------------------------------------------------
* Read timeseries data from a local .json file.
* This function assumes the filename includes the full path to the
* input file.
* -------------------------------------------------------------------- */
func TimeSeriesFromJSON(filename string, v interface{}) error {
	rawJSON, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	timeseries := v
	err = json.Unmarshal(rawJSON, &timeseries)
	return err
}

/* ---------------------------------------------------------------------
* Write timeseries data to a local .csv file.
* This function assumes the filename includes the full path to the
* output file.
* ------------------------------------------------------------------- */
func TimeSeriesToCSV(timeseries map[string]OHLCV, filename string) error {
	var err error

	// Start records off with the header row.
	records := [][]string{
		{
			"date",
			"open",
			"high",
			"low",
			"close",
			"adjusted_close",
			"volume",
			"divident_amount",
			"split_coefficient",
		},
	}

	// Append all timeseries records
	for k, v := range timeseries {
		r := []string{
			k,
			v.Open,
			v.High,
			v.Low,
			v.Close,
			v.AdjustedClose,
			v.Volume,
			v.DividendAmt,
			v.SplitCoefficient,
		}
		records = append(records, r)
	}

	// Create/Open the output file
	outfile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer outfile.Close()

	w := csv.NewWriter(outfile)
	return w.WriteAll(records)

}
