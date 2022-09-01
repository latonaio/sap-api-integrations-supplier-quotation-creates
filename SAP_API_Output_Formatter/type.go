package sap_api_output_formatter

type SupplierQuotation struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	Product           string `json:"Product"`
	APISchema         string `json:"api_schema"`
	SupplierQuotation string `json:"supplier_quotation_no"`
	Deleted           string `json:"deleted"`
}

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

type Item struct {
	SupplierQuotation              string `json:"SupplierQuotation"`
	SupplierQuotationItem          string `json:"SupplierQuotationItem"`
	PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
	PurchasingDocumentItemText     string `json:"PurchasingDocumentItemText"`
	Material                       string `json:"Material"`
	ProductTypeCode                string `json:"ProductTypeCode"`
	ManufacturerMaterial           string `json:"ManufacturerMaterial"`
	SupplierMaterialNumber         string `json:"SupplierMaterialNumber"`
	ManufacturerPartNmbr           string `json:"ManufacturerPartNmbr"`
	Manufacturer                   string `json:"Manufacturer"`
	MaterialGroup                  string `json:"MaterialGroup"`
	Plant                          string `json:"Plant"`
	ManualDeliveryAddressID        string `json:"ManualDeliveryAddressID"`
	ReferenceDeliveryAddressID     string `json:"ReferenceDeliveryAddressID"`
	AddressID                      string `json:"AddressID"`
	ItemDeliveryAddressID          string `json:"ItemDeliveryAddressID"`
	IncotermsClassification        string `json:"IncotermsClassification"`
	IncotermsTransferLocation      string `json:"IncotermsTransferLocation"`
	IncotermsLocation1             string `json:"IncotermsLocation1"`
	IncotermsLocation2             string `json:"IncotermsLocation2"`
	ScheduleLineDeliveryDate       string `json:"ScheduleLineDeliveryDate"`
	ScheduleLineOrderQuantity      int    `json:"ScheduleLineOrderQuantity"`
	AwardedQuantity                int    `json:"AwardedQuantity"`
	PerformancePeriodStartDate     string `json:"PerformancePeriodStartDate"`
	PerformancePeriodEndDate       string `json:"PerformancePeriodEndDate"`
	OrderPriceUnit                 string `json:"OrderPriceUnit"`
	OrderPriceUnitToOrderUnitNmrtr int    `json:"OrderPriceUnitToOrderUnitNmrtr"`
	OrdPriceUnitToOrderUnitDnmntr  int    `json:"OrdPriceUnitToOrderUnitDnmntr"`
	OrderQuantityUnit              string `json:"OrderQuantityUnit"`
	OrderItemQtyToBaseQtyNmrtr     int    `json:"OrderItemQtyToBaseQtyNmrtr"`
	OrderItemQtyToBaseQtyDnmntr    int    `json:"OrderItemQtyToBaseQtyDnmntr"`
	PurgDocPriceDate               string `json:"PurgDocPriceDate"`
	BaseUnit                       string `json:"BaseUnit"`
	NetAmount                      int    `json:"NetAmount"`
	GrossAmount                    int    `json:"GrossAmount"`
	EffectiveAmount                int    `json:"EffectiveAmount"`
	NetPriceAmount                 int    `json:"NetPriceAmount"`
	NetPriceQuantity               int    `json:"NetPriceQuantity"`
	DocumentCurrency               string `json:"DocumentCurrency"`
	PurchaseRequisition            string `json:"PurchaseRequisition"`
	PurchaseRequisitionItem        string `json:"PurchaseRequisitionItem"`
	RequestForQuotation            string `json:"RequestForQuotation"`
	RequestForQuotationItem        string `json:"RequestForQuotationItem"`
	PurchasingInfoRecordUpdateCode string `json:"PurchasingInfoRecordUpdateCode"`
	PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
	PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
	LastChangeDateTime             string `json:"LastChangeDateTime"`
}
