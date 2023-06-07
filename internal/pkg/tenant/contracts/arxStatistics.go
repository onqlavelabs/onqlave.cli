package contracts

type ArxInsights struct {
	TotalCluster int `json:"total_clusters"`
	TotalActive  int `json:"total_active"`
	TotalSealed  int `json:"total_sealed"`
}
