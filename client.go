package alphavantage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type errorResponse struct {
	ErrorMessage string `json:"Error Message,omitempty"`
	Information  string `json:"Information,omitempty"`
}

type Client struct {
	apiKey     string
	baseURL    string
	HTTPClient *http.Client
	Request    *http.Request
	Query      *url.Values
}

/* Return an instance of an API client object to the caller */
func NewClient(apiKey, method string) *Client {
	return &Client{
		apiKey:     apiKey,
		baseURL:    AVBaseURL,
		HTTPClient: newHTTPClient(),
	}
}

/* Initialize a new HTTP request */
func (c *Client) NewRequest(reqMethod string) error {
	var err error
	c.Request, err = http.NewRequest(reqMethod, c.baseURL, nil)
	if err != nil {
		return err
	}
	c.Request.Header.Set("Content-Type", "application/json; charset=utf-8")
	return nil
}

/* Assemble a new API query. */
func (c *Client) NewQuery(queryParams QueryParams) {
	query := c.Request.URL.Query()

	/* Set the Alphavantage function, ticker symbol, interval and APIKey */
	query.Add("function", queryParams.Function)
	query.Add("apikey", c.apiKey)

	// Check for optional args.
	if queryParams.HasSymbol() {
		query.Add("symbol", queryParams.Symbol)
	}
	if queryParams.HasInterval() {
		query.Add("interval", queryParams.Interval)
	}
	if queryParams.HasOutputSize() {
		query.Add("outputsize", queryParams.OutputSize)
	}
	if queryParams.HasDatatype() {
		query.Add("datatype", queryParams.Datatype)
	}
	if queryParams.HasKeywords() {
		query.Add("keywords", queryParams.Keywords)
	}

	/* Store the encoded query back in the client for later execution. */
	c.Request.URL.RawQuery = query.Encode()
}

/* Execute an existing API query. */
func (c *Client) ExecuteQuery(v interface{}) error {
	var err error

	// Execute the request, handle the response.
	response, err := c.HTTPClient.Do(c.Request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	results := v
	err = json.Unmarshal(body, &results)
	return err
}

/* -----------------------------------------------------------------------------
 * Local helpers
 * -------------------------------------------------------------------------- */
func minRedirects(req *http.Request, via []*http.Request) error {
	// Print the URL of the redirected request
	fmt.Println("Redirected to:", req.URL.String())
	// Allow a maximum of 5 redirects
	if len(via) >= 5 {
		return errors.New("too many redirects")
	}
	return nil
}

func newHTTPClient() *http.Client {
	return &http.Client{CheckRedirect: minRedirects, Timeout: DefaultHTTPTimeout}
}
