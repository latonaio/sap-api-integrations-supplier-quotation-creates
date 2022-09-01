package requests

type Header struct {
		SupplierQuotation              string `json:"SupplierQuotation"`
		CompanyCode                    string `json:"CompanyCode"`
		PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
		PurchasingDocumentType         string `json:"PurchasingDocumentType"`
		Supplier                       string `json:"Supplier"`
		CreatedByUser                  string `json:"CreatedByUser"`
		CreationDate                   string `json:"CreationDate"`
		Language                       string `json:"Language"`
		DocumentCurrency               string `json:"DocumentCurrency"`
		IncotermsClassification        string `json:"IncotermsClassification"`
		IncotermsTransferLocation      string `json:"IncotermsTransferLocation"`
		IncotermsVersion               string `json:"IncotermsVersion"`
		IncotermsLocation1             string `json:"IncotermsLocation1"`
		IncotermsLocation2             string `json:"IncotermsLocation2"`
		PaymentTerms                   string `json:"PaymentTerms"`
		CashDiscount1Days              int    `json:"CashDiscount1Days"`
		CashDiscount2Days              int    `json:"CashDiscount2Days"`
		CashDiscount1Percent           int    `json:"CashDiscount1Percent"`
		CashDiscount2Percent           int    `json:"CashDiscount2Percent"`
		NetPaymentDays                 int    `json:"NetPaymentDays"`
		PricingProcedure               string `json:"PricingProcedure"`
		PurchasingOrganization         string `json:"PurchasingOrganization"`
		PurchasingGroup                string `json:"PurchasingGroup"`
		PurchasingDocumentOrderDate    string `json:"PurchasingDocumentOrderDate"`
		AbsoluteExchangeRate           int    `json:"AbsoluteExchangeRate"`
		ExchRateIsIndirectQuotation    bool   `json:"ExchRateIsIndirectQuotation"`
		EffectiveExchangeRate          int    `json:"EffectiveExchangeRate"`
		ExchangeRateIsFixed            bool   `json:"ExchangeRateIsFixed"`
		PurContrValidityStartDate      string `json:"PurContrValidityStartDate"`
		PurContrValidityEndDate        string `json:"PurContrValidityEndDate"`
		IsEndOfPurposeBlocked          string `json:"IsEndOfPurposeBlocked"`
		PurchasingDocumentDeletionCode string `json:"PurchasingDocumentDeletionCode"`
		RequestForQuotation            string `json:"RequestForQuotation"`
		SupplierQuotationExternalID    string `json:"SupplierQuotationExternalID"`
		QuotationSubmissionDate        string `json:"QuotationSubmissionDate"`
		QuotationLatestSubmissionDate  string `json:"QuotationLatestSubmissionDate"`
		BindingPeriodValidityEndDate   string `json:"BindingPeriodValidityEndDate"`
		QtnLifecycleStatus             string `json:"QtnLifecycleStatus"`
		FollowOnDocumentCategory       string `json:"FollowOnDocumentCategory"`
		PurgDocFollowOnDocumentType    string `json:"PurgDocFollowOnDocumentType"`
}
