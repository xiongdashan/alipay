package models

type AliTradeItem struct {
	Target  string `json:"target"`
	Status  string `json:"status"`
	Stat    string `json:"stat"`
	Account string `Json:"account"`
	Result  struct {
		Summary struct {
			ExpendSum struct {
				Amount string `json:"amount"`
				Count  int    `json:"count"`
			} `json:"expendSum"`
			IncomeSum struct {
				Amount string `json:"amount"`
				Count  int    `json:"count"`
			} `json:"incomeSum"`
		} `json:"summary"`
		Detail []struct {
			CashierChannels      string `json:"cashierChannels"`
			ActualChargeAmount   string `json:"actualChargeAmount"`
			TradeTime            string `json:"tradeTime"`
			TradeNo              string `json:"tradeNo"`
			ChargeRate           string `json:"chargeRate"`
			OtherAccountFullname string `json:"otherAccountFullname"`
			TransMemo            string `json:"transMemo"`
			SignProduct          string `json:"signProduct"`
			Balance              string `json:"balance"`
			TransDate            string `json:"transDate"`
			OrderNo              string `json:"orderNo"`
			OtherAccount         string `json:"otherAccount"`
			AccountLogID         string `json:"accountLogId"`
			OtherAccountEmail    string `json:"otherAccountEmail"`
			Action               struct {
				NeedDetail bool `json:"needDetail"`
			} `json:"action"`
			AccountType   string `json:"accountType"`
			GoodsTitle    string `json:"goodsTitle"`
			TradeAmount   string `json:"tradeAmount"`
			DepositBankNo string `json:"depositBankNo"`
		} `json:"detail"`
		Paging struct {
			SizePerPage int `json:"sizePerPage"`
			TotalItems  int `json:"totalItems"`
			Current     int `json:"current"`
		} `json:"paging"`
	} `json:"result"`
	IsEntOperator                                      bool `json:"isEntOperator"`
	OrgSpringframeworkValidationBindingResultQueryForm struct {
		NestedPath           string `json:"nestedPath"`
		MessageCodesResolver struct {
			Prefix string `json:"prefix"`
		} `json:"messageCodesResolver"`
	} `json:"org.springframework.validation.BindingResult.queryForm"`
	AccountDetailForm struct {
		SortTarget        string `json:"sortTarget"`
		PrecisionQueryKey string `json:"precisionQueryKey"`
		PageSize          string `json:"pageSize"`
		EndDateInput      string `json:"endDateInput"`
		Type              string `json:"type"`
		PageNum           string `json:"pageNum"`
		StartDateInput    string `json:"startDateInput"`
		SortType          string `json:"sortType"`
		ShowType          string `json:"showType"`
		QueryEntrance     string `json:"queryEntrance"`
		BillUserID        string `json:"billUserId"`
	} `json:"accountDetailForm"`
	QueryForm struct {
		SortTarget        string `json:"sortTarget"`
		PrecisionQueryKey string `json:"precisionQueryKey"`
		PageSize          string `json:"pageSize"`
		EndDateInput      string `json:"endDateInput"`
		Type              string `json:"type"`
		PageNum           string `json:"pageNum"`
		StartDateInput    string `json:"startDateInput"`
		SortType          string `json:"sortType"`
		ShowType          string `json:"showType"`
		QueryEntrance     string `json:"queryEntrance"`
		BillUserID        string `json:"billUserId"`
	} `json:"queryForm"`
}
