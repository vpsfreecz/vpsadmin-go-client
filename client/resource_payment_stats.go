package client

// Type for resource Payment_stats
type ResourcePaymentStats struct {
	// Pointer to client
	Client *Client

	// Action Payment_stats#Estimate_income
	EstimateIncome *ActionPaymentStatsEstimateIncome
}

func NewResourcePaymentStats(client *Client) *ResourcePaymentStats {
	actionEstimateIncome := NewActionPaymentStatsEstimateIncome(client)

	return &ResourcePaymentStats{
		Client: client,
		EstimateIncome: actionEstimateIncome,
	}
}
