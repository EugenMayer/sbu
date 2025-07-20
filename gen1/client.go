package gen1

import (
	"encoding/json"
	"fmt"
	"github.com/fermayo/shelly-bulk-update/api"
	"github.com/fermayo/shelly-bulk-update/cli"
	"io"
	"net/http"
)

func doGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if cli.Password != nil {
		req.SetBasicAuth(*cli.Username, *cli.Password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
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

func doShellyUpdateRequest(hostname string, update bool) (*api.ShellyUpdateStatusResponse, error) {
	url := api.OtaUrl
	if update {
		if *cli.Channel == "beta" {
			url += "?beta=1"
		} else {
			url += "?update=1"
		}
	}

	body, err := doGetRequest(fmt.Sprintf(url, hostname))
	if err != nil {
		return nil, err
	}

	var updateStatus *api.ShellyUpdateStatusResponse
	err = json.Unmarshal(body, &updateStatus)
	if err != nil {
		return nil, err
	}

	return updateStatus, nil
}

func triggerShellyUpdateCheck(hostname string) (*api.ShellyUpdateCheckResponse, error) {
	body, err := doGetRequest(fmt.Sprintf(api.OtaCheckUrl, hostname))
	if err != nil {
		return nil, err
	}

	var checkStatus *api.ShellyUpdateCheckResponse
	err = json.Unmarshal(body, &checkStatus)
	if err != nil {
		return nil, err
	}

	return checkStatus, nil
}

func GetShellyInfo(hostname string) (*api.ShellyInfo, error) {
	url := api.ShellyInfoUrl
	body, err := doGetRequest(fmt.Sprintf(url, hostname))
	if err != nil {
		return nil, err
	}

	var shellyInfo *api.ShellyInfo
	err = json.Unmarshal(body, &shellyInfo)
	if err != nil {
		return nil, err
	}

	return shellyInfo, nil
}
