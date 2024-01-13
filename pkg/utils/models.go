package utils

type PaytmMoneyMfTickerResponse struct {
	Meta struct {
		Code           string `json:"code"`
		Message        string `json:"message"`
		DisplayMessage string `json:"displayMessage"`
		ResponseID     string `json:"responseId"`
		RequestID      string `json:"requestId"`
	} `json:"meta"`
	Data struct {
		PortfolioDetails     any `json:"portfolioDetails"`
		TransactionRules     any `json:"transactionRules"`
		SipInvestmentDetails struct {
			MinimumInvestment int `json:"minimumInvestment"`
			RecommendedDate   int `json:"recommendedDate"`
		} `json:"sipInvestmentDetails"`
		SipExists        bool `json:"sipExists"`
		SchemeBookmarked any  `json:"schemeBookmarked"`
		FolioInProgress  bool `json:"folioInProgress"`
		PageLoad         struct {
			NfoInfo      any `json:"nfoInfo"`
			FundManagers []struct {
				CanonicalURL string `json:"canonicalUrl"`
				ID           int    `json:"id"`
				Name         string `json:"name"`
				Designation  string `json:"designation"`
				PhotoURL     string `json:"photoUrl"`
			} `json:"fundManagers"`
			Ratings             []any  `json:"ratings"`
			FundRank            any    `json:"fundRank"`
			CategoryWiseRanking any    `json:"categoryWiseRanking"`
			Isin                string `json:"isin"`
			Name                string `json:"name"`
			Image               string `json:"image"`
			SipAllowed          any    `json:"sipAllowed"`
			LumpsumAllowed      bool   `json:"lumpsumAllowed"`
			SchemeTags          []struct {
				DisplayName string `json:"displayName"`
				Params      []struct {
					Key    string   `json:"key"`
					Values []string `json:"values"`
				} `json:"params"`
				TagType  string `json:"tagType"`
				DeepLink any    `json:"deepLink"`
				IconURL  any    `json:"iconUrl"`
			} `json:"schemeTags"`
			CategoryTags []struct {
				DisplayName string `json:"displayName"`
				Params      []struct {
					Key    string   `json:"key"`
					Values []string `json:"values"`
				} `json:"params"`
				TagType  string `json:"tagType"`
				DeepLink any    `json:"deepLink"`
				IconURL  any    `json:"iconUrl"`
			} `json:"categoryTags"`
			SipCardInfo struct {
				SuggestedDate     int  `json:"suggestedDate"`
				MinimumInvestment int  `json:"minimumInvestment"`
				SipAllowed        bool `json:"sipAllowed"`
			} `json:"sipCardInfo"`
			DefaultInvestment           int     `json:"defaultInvestment"`
			CanonicalURL                string  `json:"canonicalUrl"`
			BuyAllowed                  bool    `json:"buyAllowed"`
			SellAllowed                 bool    `json:"sellAllowed"`
			InvestmentClause            string  `json:"investmentClause"`
			AbsoluteDayReturn           float64 `json:"absoluteDayReturn"`
			AbsoluteDayReturnPercentage float64 `json:"absoluteDayReturnPercentage"`
			NavVal                      float64 `json:"navVal"`
			FormattedNavRecordTimeStamp string  `json:"formattedNavRecordTimeStamp"`
			RedemptionClause            any     `json:"redemptionClause"`
			InvestNote                  any     `json:"investNote"`
			Category                    string  `json:"category"`
		} `json:"PAGE_LOAD"`
		Holdings []struct {
			Name string `json:"name"`
			List []struct {
				Value      float64 `json:"value"`
				Name       string  `json:"name"`
				Percentage float64 `json:"percentage"`
			} `json:"list"`
			Count int `json:"count"`
		} `json:"HOLDINGS"`
		Riskometer struct {
			RiskometerID   int    `json:"riskometerId"`
			Name           string `json:"name"`
			RiskometerName string `json:"riskometerName"`
			ID             int    `json:"id"`
		} `json:"RISKOMETER"`
		Dividend          any `json:"DIVIDEND"`
		InvestmentReturns struct {
			FundReturns struct {
				AbsoluteReturns []struct {
					Name       string  `json:"name"`
					Value      float64 `json:"value"`
					Percentage float64 `json:"percentage"`
				} `json:"absoluteReturns"`
				Returns []struct {
					Name       string `json:"name"`
					Value      any    `json:"value"`
					Percentage any    `json:"percentage"`
				} `json:"returns"`
			} `json:"fundReturns"`
			Performance struct {
				FixedDepositPercentage       float64 `json:"fixedDepositPercentage"`
				SavingsBankAccountPercentage float64 `json:"savingsBankAccountPercentage"`
				DefaultInvestmentValue       int     `json:"defaultInvestmentValue"`
				InvestmentPerformance        []struct {
					Bucket  string  `json:"bucket"`
					Direct  float64 `json:"direct"`
					Regular float64 `json:"regular"`
				} `json:"investmentPerformance"`
				BucketNames []struct {
					Name     string `json:"name"`
					Selected bool   `json:"selected"`
					Value    int    `json:"value"`
				} `json:"bucketNames"`
				HigherReturnsPercentage int `json:"higherReturnsPercentage"`
			} `json:"performance"`
		} `json:"INVESTMENT_RETURNS"`
		HighReturnFundsFromThisAmc []struct {
			Name          string  `json:"name"`
			Selected      bool    `json:"selected"`
			SchemeReturns float64 `json:"schemeReturns,omitempty"`
			Funds         []struct {
				AmcImage            string `json:"amcImage"`
				CanonicalURL        string `json:"canonicalUrl"`
				CategoryDisplayName string `json:"categoryDisplayName"`
				RiskType            struct {
					Name    string `json:"name"`
					ColorID int    `json:"colorId"`
				} `json:"riskType"`
				UserRating           any     `json:"userRating"`
				Aum                  float64 `json:"aum"`
				Isin                 string  `json:"isin"`
				FundName             string  `json:"fundName"`
				MinimumInvestment    int     `json:"minimumInvestment"`
				Returns              float64 `json:"returns"`
				CategoryReturnsRange struct {
					Min float64 `json:"min"`
					Max float64 `json:"max"`
				} `json:"categoryReturnsRange"`
			} `json:"funds"`
		} `json:"HIGH_RETURN_FUNDS_FROM_THIS_AMC"`
		FundInformation struct {
			FundDetails struct {
				CashHoldingPercentage  float64 `json:"cashHoldingPercentage"`
				RecordTimestamp        int64   `json:"recordTimestamp"`
				ExitLoadClause         string  `json:"exitLoadClause"`
				LaunchDate             int64   `json:"launchDate"`
				OfferPrice             int     `json:"offerPrice"`
				AssetDate              int64   `json:"assetDate"`
				Name                   string  `json:"name"`
				FundType               string  `json:"fundType"`
				Plan                   string  `json:"plan"`
				AssetSize              float64 `json:"assetSize"`
				ExpenseRatioPercentage float64 `json:"expenseRatioPercentage"`
				SchemeBenchmark        string  `json:"schemeBenchmark"`
				SchemeDocumentsLink    string  `json:"schemeDocumentsLink"`
			} `json:"fundDetails"`
			AmcDetails struct {
				Name              string  `json:"name"`
				AmcCode           string  `json:"amcCode"`
				Image             string  `json:"image"`
				AmcID             int     `json:"amcId"`
				CanonicalURL      string  `json:"canonicalUrl"`
				ActiveFundsAumSum float64 `json:"activeFundsAumSum"`
				Telephone         string  `json:"telephone"`
				Fax               string  `json:"fax"`
				Website           string  `json:"website"`
				Email             string  `json:"email"`
				ActiveFundsCount  int     `json:"activeFundsCount"`
				Address           string  `json:"address"`
			} `json:"amcDetails"`
		} `json:"FUND_INFORMATION"`
		TopCategoryFunds []struct {
			Name          string  `json:"name"`
			Selected      bool    `json:"selected"`
			SchemeReturns float64 `json:"schemeReturns,omitempty"`
			Funds         []struct {
				AmcImage            string `json:"amcImage"`
				CanonicalURL        string `json:"canonicalUrl"`
				CategoryDisplayName string `json:"categoryDisplayName"`
				RiskType            struct {
					Name    string `json:"name"`
					ColorID int    `json:"colorId"`
				} `json:"riskType"`
				UserRating           any     `json:"userRating"`
				Aum                  float64 `json:"aum"`
				Isin                 string  `json:"isin"`
				FundName             string  `json:"fundName"`
				MinimumInvestment    int     `json:"minimumInvestment"`
				Returns              float64 `json:"returns"`
				CategoryReturnsRange struct {
					Min float64 `json:"min"`
					Max float64 `json:"max"`
				} `json:"categoryReturnsRange"`
			} `json:"funds"`
		} `json:"TOP_CATEGORY_FUNDS"`
	} `json:"data"`
}
