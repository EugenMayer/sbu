package api

const (
	// https://shelly-api-docs.shelly.cloud/gen1/#ota
	OtaUrl = "http://%s/ota"

	// https://shelly-api-docs.shelly.cloud/gen1/#ota-check
	OtaCheckUrl = "http://%s/ota/check"

	// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#shellycheckforupdate
	CheckForUpdateUrl = "http://%s/rpc/Shelly.CheckForUpdate"

	// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#shellyupdate
	UpdateUrl = "http://%s/rpc/Shelly.Update?stage=%s"

	// https://shelly-api-docs.shelly.cloud/gen1/#shelly
	ShellyInfoUrl = "http://%s/shelly"
)

type (
	versionInfo struct {
		Version string `json:"version"`
		BuildId string `json:"build_id"`
	}

	CheckForUpdateResponse struct {
		Stable versionInfo `json:"stable"`
		Beta   versionInfo `json:"beta"`
	}

	ShellyUpdateStatusResponse struct {
		Status      string `json:"status"`
		HasUpdate   bool   `json:"has_update"`
		NewVersion  string `json:"new_version"`
		OldVersion  string `json:"old_version"`
		BetaVersion string `json:"beta_version"`
	}

	ShellyUpdateCheckResponse struct {
		Status string `json:"status"`
	}

	ShellyInfo struct {
		Id              string `json:"id"`
		Name            string `json:"name"`
		Model           string `json:"model"`
		Generation      int    `json:"gen"`
		FirmwareVersion string `json:"ver"`
		AuthDomain      string `json:"auth_domain"`
	}
)
