package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"
)

func DoRequest(httpClient *http.Client, req *http.Request, resultTemplate interface{}) error {
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Cannot read %d error response", res.StatusCode)
		}
		return fmt.Errorf("Status: %d, Error body: %s", res.StatusCode, bodyBytes)
	}

	if err := render.DecodeJSON(res.Body, &resultTemplate); err != nil {
		return fmt.Errorf("Couldn't decode message %w", err)
	}
	return nil
}
