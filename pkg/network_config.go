package pkg

type NetworkConfig struct {
	RelayerCount            int32   `json:"relayer_count"`
	ProviderUrl             string  `json:"provider_url"`
	FundAmountInEther       float32 `json:"fund_amount_in_ether"`
	MinBalanceToFundInEther float32 `json:"min_balance_to_fund_in_ether"`
	FundingWalletPrivateKey string  `json:"funding_wallet_private_key"`
}
