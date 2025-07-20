package gen2plus

import (
	"encoding/json"
	"fmt"
	"github.com/fermayo/shelly-bulk-update/api"
	"github.com/fermayo/shelly-bulk-update/cli"
	"github.com/mongodb-forks/digest"
	"io"
	"net/http"
)

func doGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var resp *http.Response
	if cli.Password != nil {
		t := digest.NewTransport(*cli.Username, *cli.Password)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err = t.RoundTrip(req)
		if err != nil {
			return nil, err
		}
	} else {
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func doCheckForUpdateRequest(hostname string) (*api.CheckForUpdateResponse, error) {
	body, err := doGetRequest(fmt.Sprintf(api.CheckForUpdateUrl, hostname))
	if err != nil {
		return nil, err
	}

	var checkForUpdate *api.CheckForUpdateResponse
	err = json.Unmarshal(body, &checkForUpdate)
	if err != nil {
		return nil, err
	}

	return checkForUpdate, nil
}

func doUpdateRequest(hostname string, stage string) error {
	_, err := doGetRequest(fmt.Sprintf(api.UpdateUrl, hostname, stage))
	if err != nil {
		return err
	}

	return nil
}
