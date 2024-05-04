package main

import "time"

type MeResponse struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	CompanyName  string     `json:"company_name"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	PlanId       int        `json:"plan_id"`
	Reports      int        `json:"reports"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	NextBillDate *time.Time `json:"next_bill_date"`
	ApiToken     string     `json:"api_token"`
	CardFirstSix *string    `json:"card_first_six"`
	CardLastFour *string    `json:"card_last_four"`
	CardType     *string    `json:"card_type"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type SearchCompanyResponse struct {
	Company []struct {
		Id          string `json:"id"`
		Title       string `json:"title"`
		Identifiers []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"identifiers"`
		Name         string `json:"name"`
		NameEn       string `json:"name_en"`
		Inactive     bool   `json:"inactive"`
		Jurisdiction string `json:"jurisdiction"`
		Addresses    []struct {
			Raw   string `json:"raw"`
			State string `json:"state"`
		} `json:"addresses"`
		Structure string `json:"structure"`
	} `json:"company"`
	Count            int `json:"count"`
	PaginationCount  int `json:"pagination_count"`
	PaginationOffset int `json:"pagination_offset"`
}

type GetCompanyResponse struct {
	Id            int    `json:"id"`
	Identifier    string `json:"identifier"`
	Title         string `json:"title"`
	Name          string `json:"name"`
	NameEn        string `json:"name_en"`
	NameNative    string `json:"name_native"`
	LegalForm     string `json:"legal_form"`
	CompanySize   string `json:"company_size"`
	Jurisdiction  string `json:"jurisdiction"`
	Industry      string `json:"industry"`
	Status        string `json:"status"`
	OwnershipType string `json:"ownership_type"`
	VatPayer      bool   `json:"vat_payer"`
	Inactive      bool   `json:"inactive"`
	Addresses     []struct {
		Raw          string      `json:"raw"`
		State        string      `json:"state"`
		Path         string      `json:"path"`
		PropertyType interface{} `json:"property_type"`
	} `json:"addresses"`
	Contacts struct {
		Website []struct {
			Id    int    `json:"id"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"website"`
		Phone []struct {
			Id    int    `json:"id"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"phone"`
		Email []struct {
			Id    int    `json:"id"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"email"`
	} `json:"contacts"`
	Financials []struct {
		Year    int   `json:"year"`
		Taxes   int64 `json:"taxes"`
		Revenue int64 `json:"revenue"`
	} `json:"financials"`
	Headcount []struct {
		Count string    `json:"count"`
		Date  time.Time `json:"date"`
	} `json:"headcount"`
	Identifiers []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"identifiers"`
	MainIndustryCode interface{} `json:"main_industry_code"`
	IndustryCodes    []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
		IsMain      bool   `json:"is_main"`
	} `json:"industry_codes"`
	Officers []struct {
		Identifier interface{} `json:"identifier"`
		Name       string      `json:"name"`
		NameEn     string      `json:"name_en"`
		NameNative string      `json:"name_native"`
		Role       string      `json:"role"`
		StartDate  *time.Time  `json:"start_date"`
		EndDate    interface{} `json:"end_date"`
		Inactive   bool        `json:"inactive"`
		Contacts   interface{} `json:"contacts"`
		Risks      interface{} `json:"risks"`
		HashBin    string      `json:"hash_bin"`
	} `json:"officers"`
	Shareholders []struct {
		ShareholderCompanyId int         `json:"shareholder_company_id"`
		Identifier           string      `json:"identifier"`
		Name                 string      `json:"name"`
		NameEn               string      `json:"name_en"`
		NameNative           string      `json:"name_native"`
		Share                interface{} `json:"share"`
		ShareCapital         interface{} `json:"share_capital"`
		Inactive             bool        `json:"inactive"`
		StartDate            interface{} `json:"start_date"`
		EndDate              interface{} `json:"end_date"`
		IsPerson             bool        `json:"is_person"`
		Contacts             interface{} `json:"contacts"`
		Risks                interface{} `json:"risks"`
		HashBin              string      `json:"hash_bin"`
	} `json:"shareholders"`
	Risks struct {
		Company []struct {
			Status  string      `json:"status"`
			Object  string      `json:"object"`
			Source  string      `json:"source"`
			Records interface{} `json:"records"`
		} `json:"company"`
		Shareholders []struct {
			Status  string      `json:"status"`
			Object  string      `json:"object"`
			Source  string      `json:"source"`
			Records interface{} `json:"records"`
		} `json:"shareholders"`
		Officers []struct {
			Status  string      `json:"status"`
			Object  string      `json:"object"`
			Source  string      `json:"source"`
			Records interface{} `json:"records"`
		} `json:"officers"`
	} `json:"risks"`
	Capital struct {
		ShareCapital  int64       `json:"share_capital"`
		PaidUpCapital interface{} `json:"paid_up_capital"`
		Currency      interface{} `json:"currency"`
	} `json:"capital"`
	ManagingCompanyId   interface{} `json:"managing_company_id"`
	ManagingCompanyName interface{} `json:"managing_company_name"`
	Branches            []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		NameEn string `json:"name_en"`
	} `json:"branches"`
	Names                []interface{} `json:"names"`
	Export               []string      `json:"export"`
	Import               []string      `json:"import"`
	LastUpdateDate       time.Time     `json:"last_update_date"`
	DissolutionDate      interface{}   `json:"dissolution_date"`
	IncorporationDate    time.Time     `json:"incorporation_date"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
	AlternativeCompanies interface{}   `json:"alternative_companies"`
	Structure            string        `json:"structure"`
	Aliases              []struct {
		Name string `json:"name"`
	} `json:"aliases"`
}

type GetCompanyMetaResponse struct {
	International struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"international"`
	Founders struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"founders"`
	RegDepartment struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"reg_department"`
	RegStatus struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"reg_status"`
	BranchesExistence struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"branches_existence"`
	Investors struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"investors"`
	CreationMethod struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"creation_method"`
	Registrations struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"registrations"`
	Scoring struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			RetrievedAt time.Time `json:"retrieved_at"`
			Score       int       `json:"score"`
			Group       string    `json:"group"`
			BadRate     float64   `json:"bad_rate"`
		} `json:"data"`
	} `json:"scoring"`
	Certs struct {
		Key   string        `json:"key"`
		Value string        `json:"value"`
		Data  []interface{} `json:"data"`
	} `json:"certs"`
	ConformityCerts struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"conformity_certs"`
	Debtors struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			EpNumber     string      `json:"ep_number"`
			Active       bool        `json:"active"`
			EpStartDate  time.Time   `json:"ep_start_date"`
			EpEndDate    interface{} `json:"ep_end_date"`
			BanStartDate interface{} `json:"ban_start_date"`
			BanEndDate   interface{} `json:"ban_end_date"`
			DebtorName   string      `json:"debtor_name"`
			Amount       int         `json:"amount"`
			Source       string      `json:"source"`
			Currency     string      `json:"currency"`
		} `json:"data"`
	} `json:"debtors"`
	Stores struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			StoreName string      `json:"store_name"`
			Address   interface{} `json:"address"`
			License   interface{} `json:"license"`
		} `json:"data"`
	} `json:"stores"`
	Jobs struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			Title        string    `json:"title"`
			Salary       string    `json:"salary"`
			JobType      string    `json:"job_type"`
			WorkSchedule string    `json:"work_schedule"`
			Experience   string    `json:"experience"`
			Education    string    `json:"education"`
			StartDate    time.Time `json:"start_date"`
		} `json:"data"`
	} `json:"jobs"`
	Trademarks struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			TrademarkName   string    `json:"trademark_name"`
			RegNumber       string    `json:"reg_number"`
			RegDate         time.Time `json:"reg_date"`
			ExpDate         time.Time `json:"exp_date"`
			PdfUrl          string    `json:"pdf_url"`
			Classifications []string  `json:"classifications"`
			SourceUrl       string    `json:"source_url"`
			BulletinDate    time.Time `json:"bulletin_date"`
			BulletinNumber  string    `json:"bulletin_number"`
		} `json:"data"`
	} `json:"trademarks"`
	QuasiContracts struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			LotName      string    `json:"lot_name"`
			CustomerName string    `json:"customer_name"`
			Price        float64   `json:"price"`
			Date         time.Time `json:"date"`
			SourceUrl    string    `json:"source_url"`
			Source       string    `json:"source"`
		} `json:"data"`
	} `json:"quasi_contracts"`
	PosTerminals struct {
		Key   string        `json:"key"`
		Value string        `json:"value"`
		Data  []interface{} `json:"data"`
	} `json:"pos_terminals"`
	Places struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			PlaceName string   `json:"place_name"`
			Area      string   `json:"area"`
			City      string   `json:"city"`
			Address   string   `json:"address"`
			Contacts  []string `json:"contacts"`
		} `json:"data"`
	} `json:"places"`
	BankAccounts struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			Iban      string      `json:"iban"`
			BankName  string      `json:"bank_name"`
			BankBic   string      `json:"bank_bic"`
			DateOpen  time.Time   `json:"date_open"`
			Holder    string      `json:"holder"`
			CreatedAt time.Time   `json:"created_at"`
			UpdatedAt time.Time   `json:"updated_at"`
			Type      interface{} `json:"type"`
		} `json:"data"`
	} `json:"bank_accounts"`
	Marketplaces struct {
		Key   string        `json:"key"`
		Value string        `json:"value"`
		Data  []interface{} `json:"data"`
	} `json:"marketplaces"`
	Counterparties struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			SupplierIdentifier string `json:"supplier_identifier"`
			CustomerIdentifier string `json:"customer_identifier"`
			Source             string `json:"source"`
		} `json:"data"`
	} `json:"counterparties"`
	Websites struct {
		Key   string `json:"key"`
		Value string `json:"value"`
		Data  []struct {
			Value          string      `json:"value"`
			Domain         string      `json:"domain"`
			Title          string      `json:"title"`
			Entities       []string    `json:"entities"`
			SocialNetworks interface{} `json:"social_networks"`
			Contacts       []string    `json:"contacts"`
		} `json:"data"`
	} `json:"websites"`
	GovLoans struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"gov_loans"`
	Farmers struct {
		Key   string        `json:"key"`
		Value string        `json:"value"`
		Data  []interface{} `json:"data"`
	} `json:"farmers"`
	Names struct {
		Key   string        `json:"key"`
		Value string        `json:"value"`
		Data  []interface{} `json:"data"`
	} `json:"names"`
	Registries struct {
		Key   string        `json:"key"`
		Value string        `json:"value"`
		Data  []interface{} `json:"data"`
	} `json:"registries"`
	Assets struct {
		Key   string      `json:"key"`
		Value string      `json:"value"`
		Data  interface{} `json:"data"`
	} `json:"assets"`
}

type GetCompanyCourtCasesResponse struct {
	CourtCases []struct {
		ExternalId   string        `json:"external_id"`
		CaseNumber   string        `json:"case_number"`
		Instance     string        `json:"instance"`
		CaseType     string        `json:"case_type"`
		StartDate    string        `json:"start_date"`
		EndDate      interface{}   `json:"end_date"`
		Court        string        `json:"court"`
		Category     string        `json:"category"`
		Judge        string        `json:"judge"`
		Sides        []string      `json:"sides"`
		Defendants   []interface{} `json:"defendants"`
		Plaintiffs   []interface{} `json:"plaintiffs"`
		Jurisdiction string        `json:"jurisdiction"`
	} `json:"court_cases"`
	CaseTypeMeta []struct {
		Type  string `json:"type"`
		Total int    `json:"total"`
	} `json:"case_type_meta"`
	Meta []struct {
		Type  string `json:"type"`
		Total int    `json:"total"`
	} `json:"meta"`
	YearMeta []struct {
		Type  string `json:"type"`
		Total int    `json:"total"`
	} `json:"year_meta"`
	Total int `json:"total"`
}

type GetCompanyDepartmentsResponse struct {
	Jurisdiction string   `json:"jurisdiction"`
	Departments  []string `json:"departments"`
}

type GetCompanyGovContractsResponse struct {
	GovernmentContract []struct {
		SupplierBiin       string      `json:"supplier_biin"`
		SupplierBik        string      `json:"supplier_bik"`
		SupplierIik        string      `json:"supplier_iik"`
		SupplierBankNameKz string      `json:"supplier_bank_name_kz"`
		SupplierBankNameRu string      `json:"supplier_bank_name_ru"`
		SignReasonDocName  string      `json:"sign_reason_doc_name"`
		CustomerBin        string      `json:"customer_bin"`
		CustomerBik        string      `json:"customer_bik"`
		CustomerIik        string      `json:"customer_iik"`
		CustomerBankNameKz string      `json:"customer_bank_name_kz"`
		CustomerBankNameRu string      `json:"customer_bank_name_ru"`
		ContractSumWnds    float64     `json:"contract_sum_wnds"`
		SignDate           time.Time   `json:"sign_date"`
		EcEndDate          time.Time   `json:"ec_end_date"`
		DescriptionKz      string      `json:"description_kz"`
		DescriptionRu      string      `json:"description_ru"`
		SupplierName       interface{} `json:"supplier_name"`
		CustomerName       interface{} `json:"customer_name"`
		CustomerStatsnetId interface{} `json:"customer_statsnet_id"`
		SupplierStatsnetId interface{} `json:"supplier_statsnet_id"`
	} `json:"government_contract"`
	Count            int `json:"count"`
	PaginationCount  int `json:"pagination_count"`
	PaginationOffset int `json:"pagination_offset"`
}

type GetCompanyEventsResponse struct {
	Events []struct {
		Jurisdiction string    `json:"jurisdiction"`
		Date         time.Time `json:"date"`
		Attr         string    `json:"attr"`
		Before       string    `json:"before"`
		After        string    `json:"after"`
	} `json:"events"`
	Meta []struct {
		Type  string `json:"type"`
		Count int    `json:"count"`
	} `json:"meta"`
	Total            int `json:"total"`
	PaginationCount  int `json:"pagination_count"`
	PaginationOffset int `json:"pagination_offset"`
}
type GetCompanyRelationsResponse struct {
	Links []struct {
		Id                       int    `json:"id"`
		Name                     string `json:"name"`
		SourceKey                int64  `json:"source_key"`
		Type                     string `json:"type"`
		ConnectionType           string `json:"connection_type"`
		ConnectionId             int    `json:"connection_id"`
		ConnectionName           string `json:"connection_name"`
		RootCompanyId            int    `json:"root_company_id"`
		ConnectionHashIdentifier string `json:"connection_hash_identifier"`
		ConnectionKey            int64  `json:"connection_key"`
		IsPerson                 bool   `json:"is_person"`
	} `json:"links"`
	Meta struct {
		Subsidiaries int `json:"subsidiaries"`
		Branches     int `json:"branches"`
		Contacts     int `json:"contacts"`
		Shareholders int `json:"shareholders"`
		Officers     int `json:"officers"`
		Affiliations int `json:"affiliations"`
	} `json:"meta"`
	Total int `json:"total"`
}

type GetCompanyByIdentifierResponse struct {
	Id            int    `json:"id"`
	Identifier    string `json:"identifier"`
	Title         string `json:"title"`
	Name          string `json:"name"`
	NameEn        string `json:"name_en"`
	NameNative    string `json:"name_native"`
	LegalForm     string `json:"legal_form"`
	CompanySize   string `json:"company_size"`
	Jurisdiction  string `json:"jurisdiction"`
	Industry      string `json:"industry"`
	Status        string `json:"status"`
	OwnershipType string `json:"ownership_type"`
	VatPayer      bool   `json:"vat_payer"`
	Inactive      bool   `json:"inactive"`
	Addresses     []struct {
		Raw          string      `json:"raw"`
		State        string      `json:"state"`
		Path         string      `json:"path"`
		PropertyType interface{} `json:"property_type"`
	} `json:"addresses"`
	Contacts struct {
		Website []struct {
			Id    int    `json:"id"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"website"`
		Phone []struct {
			Id    int    `json:"id"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"phone"`
		Email []struct {
			Id    int    `json:"id"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"email"`
	} `json:"contacts"`
	Financials []struct {
		Year    int   `json:"year"`
		Taxes   int64 `json:"taxes"`
		Revenue int64 `json:"revenue"`
	} `json:"financials"`
	Headcount []struct {
		Count string    `json:"count"`
		Date  time.Time `json:"date"`
	} `json:"headcount"`
	Identifiers []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"identifiers"`
	MainIndustryCode interface{} `json:"main_industry_code"`
	IndustryCodes    []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
		IsMain      bool   `json:"is_main"`
	} `json:"industry_codes"`
	Officers []struct {
		Identifier interface{} `json:"identifier"`
		Name       string      `json:"name"`
		NameEn     string      `json:"name_en"`
		NameNative string      `json:"name_native"`
		Role       string      `json:"role"`
		StartDate  *time.Time  `json:"start_date"`
		EndDate    interface{} `json:"end_date"`
		Inactive   bool        `json:"inactive"`
		Contacts   interface{} `json:"contacts"`
		Risks      interface{} `json:"risks"`
		HashBin    string      `json:"hash_bin"`
	} `json:"officers"`
	Shareholders []struct {
		ShareholderCompanyId int         `json:"shareholder_company_id"`
		Identifier           string      `json:"identifier"`
		Name                 string      `json:"name"`
		NameEn               string      `json:"name_en"`
		NameNative           string      `json:"name_native"`
		Share                interface{} `json:"share"`
		ShareCapital         interface{} `json:"share_capital"`
		Inactive             bool        `json:"inactive"`
		StartDate            interface{} `json:"start_date"`
		EndDate              interface{} `json:"end_date"`
		IsPerson             bool        `json:"is_person"`
		Contacts             interface{} `json:"contacts"`
		Risks                interface{} `json:"risks"`
		HashBin              string      `json:"hash_bin"`
	} `json:"shareholders"`
	Risks struct {
		Company []struct {
			Status  string      `json:"status"`
			Object  string      `json:"object"`
			Source  string      `json:"source"`
			Records interface{} `json:"records"`
		} `json:"company"`
		Shareholders []struct {
			Status  string      `json:"status"`
			Object  string      `json:"object"`
			Source  string      `json:"source"`
			Records interface{} `json:"records"`
		} `json:"shareholders"`
		Officers []struct {
			Status  string      `json:"status"`
			Object  string      `json:"object"`
			Source  string      `json:"source"`
			Records interface{} `json:"records"`
		} `json:"officers"`
	} `json:"risks"`
	Capital struct {
		ShareCapital  int64       `json:"share_capital"`
		PaidUpCapital interface{} `json:"paid_up_capital"`
		Currency      interface{} `json:"currency"`
	} `json:"capital"`
	ManagingCompanyId   interface{} `json:"managing_company_id"`
	ManagingCompanyName interface{} `json:"managing_company_name"`
	Branches            []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		NameEn string `json:"name_en"`
	} `json:"branches"`
	Names                []interface{} `json:"names"`
	Export               []string      `json:"export"`
	Import               []string      `json:"import"`
	LastUpdateDate       time.Time     `json:"last_update_date"`
	DissolutionDate      interface{}   `json:"dissolution_date"`
	IncorporationDate    time.Time     `json:"incorporation_date"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
	AlternativeCompanies interface{}   `json:"alternative_companies"`
	Structure            string        `json:"structure"`
	Aliases              []struct {
		Name string `json:"name"`
	} `json:"aliases"`
}
