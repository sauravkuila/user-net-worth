package external

import "time"

type HoldingsInfo struct {
	Symbol    string
	Quantity  int64
	Isin      string
	AvgPrice  float64
	UpdatedOn time.Time
}

type IDirectCustomerDetailRequest struct {
	SessionToken string
	AppKey       string
}

type IDirectCustomerDetailResponse struct {
	Success struct {
		ExgTradeDate struct {
			Nse string `json:"NSE"`
			Bse string `json:"BSE"`
			Fno string `json:"FNO"`
			Ndx string `json:"NDX"`
		} `json:"exg_trade_date"`
		ExgStatus struct {
			Nse string `json:"NSE"`
			Bse string `json:"BSE"`
			Fno string `json:"FNO"`
			Ndx string `json:"NDX"`
		} `json:"exg_status"`
		SegmentsAllowed struct {
			Trading     string `json:"Trading"`
			Equity      string `json:"Equity"`
			Derivatives string `json:"Derivatives"`
			Currency    string `json:"Currency"`
		} `json:"segments_allowed"`
		IdirectUserid           string `json:"idirect_userid"`
		SessionToken            string `json:"session_token"`
		IdirectUserName         string `json:"idirect_user_name"`
		IdirectORDTYP           string `json:"idirect_ORD_TYP"`
		IdirectLastloginTime    string `json:"idirect_lastlogin_time"`
		MfHoldingModePopupFlg   string `json:"mf_holding_mode_popup_flg"`
		CommodityExchangeStatus string `json:"commodity_exchange_status"`
		CommodityTradeDate      string `json:"commodity_trade_date"`
		CommodityAllowed        string `json:"commodity_allowed"`
	} `json:"Success"`
	Status int         `json:"Status"`
	Error  interface{} `json:"Error"`
}

type IDirectHoldingsResponse struct {
	Success []struct {
		StockCode             string      `json:"stock_code"`
		ExchangeCode          string      `json:"exchange_code"`
		Quantity              string      `json:"quantity"`
		AveragePrice          string      `json:"average_price"`
		BookedProfitLoss      string      `json:"booked_profit_loss"`
		CurrentMarketPrice    string      `json:"current_market_price"`
		ChangePercentage      string      `json:"change_percentage"`
		AnswerFlag            string      `json:"answer_flag"`
		ProductType           string      `json:"product_type"`
		ExpiryDate            interface{} `json:"expiry_date"`
		StrikePrice           interface{} `json:"strike_price"`
		Right                 interface{} `json:"right"`
		CategoryIndexPerStock interface{} `json:"category_index_per_stock"`
		Action                interface{} `json:"action"`
		RealizedProfit        interface{} `json:"realized_profit"`
		UnrealizedProfit      interface{} `json:"unrealized_profit"`
		OpenPositionValue     interface{} `json:"open_position_value"`
		PortfolioCharges      interface{} `json:"portfolio_charges"`
	} `json:"Success"`
	Status int         `json:"Status"`
	Error  interface{} `json:"Error"`
}

type IDirectDematHoldingsResponse struct {
	Success []struct {
		StockCode              string `json:"stock_code"`
		StockISIN              string `json:"stock_ISIN"`
		Quantity               string `json:"quantity"`
		DematTotalBulkQuantity string `json:"demat_total_bulk_quantity"`
		DematAvailQuantity     string `json:"demat_avail_quantity"`
		BlockedQuantity        string `json:"blocked_quantity"`
		DematAllocatedQuantity string `json:"demat_allocated_quantity"`
	} `json:"Success"`
	Status int         `json:"Status"`
	Error  interface{} `json:"Error"`
}

type ZerodhaPortfolioResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Tradingsymbol      string  `json:"tradingsymbol"`
		Exchange           string  `json:"exchange"`
		InstrumentToken    int     `json:"instrument_token"`
		Isin               string  `json:"isin"`
		Product            string  `json:"product"`
		Price              float64 `json:"price"`
		Quantity           int64   `json:"quantity"`
		UsedQuantity       int64   `json:"used_quantity"`
		T1Quantity         int64   `json:"t1_quantity"`
		RealisedQuantity   int64   `json:"realised_quantity"`
		AuthorisedQuantity int64   `json:"authorised_quantity"`
		AuthorisedDate     string  `json:"authorised_date"`
		Authorisation      struct {
		} `json:"authorisation"`
		OpeningQuantity     int64   `json:"opening_quantity"`
		ShortQuantity       int64   `json:"short_quantity"`
		CollateralQuantity  int64   `json:"collateral_quantity"`
		CollateralType      string  `json:"collateral_type"`
		Discrepancy         bool    `json:"discrepancy"`
		AveragePrice        float64 `json:"average_price"`
		LastPrice           float64 `json:"last_price"`
		ClosePrice          float64 `json:"close_price"`
		Pnl                 float64 `json:"pnl"`
		DayChange           float64 `json:"day_change"`
		DayChangePercentage float64 `json:"day_change_percentage"`
	} `json:"data"`
}

type ZerodhaLoginResponse struct {
	Status string `json:"status"`
	Data   struct {
		UserID      string   `json:"user_id"`
		RequestID   string   `json:"request_id"`
		TwofaType   string   `json:"twofa_type"`
		TwofaTypes  []string `json:"twofa_types"`
		TwofaStatus string   `json:"twofa_status"`
		Profile     struct {
			UserName      string `json:"user_name"`
			UserShortname string `json:"user_shortname"`
			AvatarURL     string `json:"avatar_url"`
		} `json:"profile"`
	} `json:"data"`
	Message   string `json:"message"`
	ErrorType string `json:""error_type`
}
