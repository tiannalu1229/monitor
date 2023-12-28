package data

import (
	"time"
)

type GoPlus struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Result  map[string]GoPlusResult `json:"result"`
}

type GoPlusResult struct {
	AntiWhaleModifiable        string           `json:"anti_whale_modifiable"`
	BuyTax                     string           `json:"buy_tax"`
	CanTakeBackOwnership       string           `json:"can_take_back_ownership"`
	CannotBuy                  string           `json:"cannot_buy"`
	CannotSellAll              string           `json:"cannot_sell_all"`
	CreatorAddress             string           `json:"creator_address"`
	CreatorBalance             string           `json:"creator_balance"`
	CreatorPercent             string           `json:"creator_percent"`
	Dex                        []GoPlusDex      `json:"dex"`
	ExternalCall               string           `json:"external_call"`
	HiddenOwner                string           `json:"hidden_owner"`
	HolderCount                string           `json:"holder_count"`
	Holders                    []GoPlusHolder   `json:"holders"`
	HoneypotWithSameCreator    string           `json:"honeypot_with_same_creator"`
	IsAntiWhale                string           `json:"is_anti_whale"`
	IsBlacklisted              string           `json:"is_blacklisted"`
	IsHoneypot                 string           `json:"is_honeypot"`
	IsInDex                    string           `json:"is_in_dex"`
	IsMintable                 string           `json:"is_mintable"`
	IsOpenSource               string           `json:"is_open_source"`
	IsProxy                    string           `json:"is_proxy"`
	IsWhitelisted              string           `json:"is_whitelisted"`
	LpHolderCount              string           `json:"lp_holder_count"`
	LpHolders                  []GoPlusLpHolder `json:"lp_holders"`
	LpTotalSupply              string           `json:"lp_total_supply"`
	OwnerAddress               string           `json:"owner_address"`
	OwnerBalance               string           `json:"owner_balance"`
	OwnerChangeBalance         string           `json:"owner_change_balance"`
	OwnerPercent               string           `json:"owner_percent"`
	PersonalSlippageModifiable string           `json:"personal_slippage_modifiable"`
	Selfdestruct               string           `json:"selfdestruct"`
	SellTax                    string           `json:"sell_tax"`
	SlippageModifiable         string           `json:"slippage_modifiable"`
	TokenName                  string           `json:"token_name"`
	TokenSymbol                string           `json:"token_symbol"`
	TotalSupply                string           `json:"total_supply"`
	TradingCooldown            string           `json:"trading_cooldown"`
	TransferPausable           string           `json:"transfer_pausable"`
}

type GoPlusDex struct {
	LiquidityType string `json:"liquidity_type"`
	Name          string `json:"name"`
	Liquidity     string `json:"liquidity"`
	Pair          string `json:"pair"`
}

type GoPlusHolder struct {
	Address    string `json:"address"`
	Tag        string `json:"tag"`
	IsContract int    `json:"is_contract"`
	Balance    string `json:"balance"`
	Percent    string `json:"percent"`
	IsLocked   int    `json:"is_locked"`
}

type GoPlusLpHolder struct {
	Address      string                       `json:"address"`
	Tag          string                       `json:"tag"`
	Value        interface{}                  `json:"value"`
	IsContract   int                          `json:"is_contract"`
	Balance      string                       `json:"balance"`
	Percent      string                       `json:"percent"`
	NFTList      interface{}                  `json:"NFT_list"`
	IsLocked     int                          `json:"is_locked"`
	LockedDetail []GoPlusLpHolderLockedDetail `json:"locked_detail,omitempty"`
}

type GoPlusLpHolderLockedDetail struct {
	Amount  string    `json:"amount"`
	EndTime time.Time `json:"end_time"`
	OptTime time.Time `json:"opt_time"`
}

type TokenSniffer struct {
	Message      string        `json:"message"`
	Status       string        `json:"status"`
	ChainId      string        `json:"chainId"`
	Address      string        `json:"address"`
	RefreshedAt  int64         `json:"refreshed_at"`
	Name         string        `json:"name"`
	Symbol       string        `json:"symbol"`
	TotalSupply  int           `json:"total_supply"`
	Decimals     int           `json:"decimals"`
	CreatedAt    int64         `json:"created_at"`
	DeployerAddr string        `json:"deployer_addr"`
	IsFlagged    bool          `json:"is_flagged"`
	Exploits     []interface{} `json:"exploits"`
	Contract     struct {
		IsSourceVerified        bool `json:"is_source_verified"`
		HasMint                 bool `json:"has_mint"`
		HasFeeModifier          bool `json:"has_fee_modifier"`
		HasMaxTransactionAmount bool `json:"has_max_transaction_amount"`
		HasBlocklist            bool `json:"has_blocklist"`
		HasProxy                bool `json:"has_proxy"`
		HasPausable             bool `json:"has_pausable"`
	} `json:"contract"`
	Score     int    `json:"score"`
	RiskLevel string `json:"riskLevel"`
	Tests     []struct {
		Id          string        `json:"id"`
		Description string        `json:"description"`
		Result      bool          `json:"result"`
		Value       float64       `json:"value,omitempty"`
		ValuePct    float64       `json:"valuePct,omitempty"`
		Data        []interface{} `json:"data,omitempty"`
		Currency    string        `json:"currency,omitempty"`
	} `json:"tests"`
	Permissions struct {
		OwnerAddress         string `json:"owner_address"`
		IsOwnershipRenounced bool   `json:"is_ownership_renounced"`
	} `json:"permissions"`
	SwapSimulation struct {
		IsSellable bool    `json:"is_sellable"`
		BuyFee     float64 `json:"buy_fee"`
		SellFee    float64 `json:"sell_fee"`
	} `json:"swap_simulation"`
	Balances struct {
		BurnBalance     int `json:"burn_balance"`
		LockBalance     int `json:"lock_balance"`
		DeployerBalance int `json:"deployer_balance"`
		OwnerBalance    int `json:"owner_balance"`
		TopHolders      []struct {
			Address    string  `json:"address"`
			Balance    float64 `json:"balance"`
			IsContract bool    `json:"is_contract"`
		} `json:"top_holders"`
	} `json:"balances"`
	Pools []struct {
		Address            string  `json:"address"`
		Name               string  `json:"name"`
		Version            string  `json:"version"`
		BaseSymbol         string  `json:"base_symbol"`
		BaseAddress        string  `json:"base_address"`
		TotalSupply        int64   `json:"total_supply"`
		Decimals           int     `json:"decimals"`
		BaseReserve        float64 `json:"base_reserve"`
		InitialBaseReserve int     `json:"initial_base_reserve"`
		OwnerBalance       float64 `json:"owner_balance"`
		DeployerBalance    int     `json:"deployer_balance"`
		BurnBalance        int     `json:"burn_balance"`
		LockBalance        int64   `json:"lock_balance"`
		TopHolders         []struct {
			Address string  `json:"address"`
			Balance float64 `json:"balance"`
		} `json:"top_holders"`
		Locks []struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Url     string `json:"url"`
			Balance int64  `json:"balance"`
			EndTime int    `json:"end_time"`
		} `json:"locks"`
	} `json:"pools"`
}
