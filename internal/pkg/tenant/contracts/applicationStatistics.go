package contracts

type ApplicationStatistics struct {
	Total    int16 `json:"total_applications"`
	Sealed   int16 `json:"sealed_applications"`
	Archived int16 `json:"archived_applications"`
}
