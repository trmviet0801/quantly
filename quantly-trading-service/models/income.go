package models

type Income struct {
	StockSymbol                                         string  `binding:"required" json:"stock_symbol"`
	TotalRevenue                                        float64 `json:"total_revenue"`
	OperationgRevenue                                   float64 `json:"operationg_revenue"`
	CostOfRevenue                                       float64 `json:"cost_of_revenue"`
	GrossProfit                                         float64 `json:"gross_profit"`
	OperatingExpense                                    float64 `json:"operating_expense"`
	SellingGeneralAndAdministrative                     float64 `json:"selling_general_and_administrative"`
	ResearchAndDevelopment                              float64 `json:"research_and_development"`
	OperatingIncome                                     float64 `json:"operating_income"`
	NetNonOperatingInterestIncomeExpense                float64 `json:"net_non_operating_interest_income_expense"`
	InterestIncomeNonOperating                          float64 `json:"interest_income_non_operating"`
	InterestExpnseNonOperating                          float64 `json:"interest_expense_non_operating"`
	OtherIncomeExpenseNet                               float64 `json:"other_income_expense_net"`
	SpecialIncomeCharges                                float64 `json:"special_income_charges"`
	RestructuringAndMergerAcquisition                   float64 `json:"restructuring_and_merger_acquisition"`
	OterNonOperatingIncomeExpense                       float64 `json:"oter_non_operating_income_expense"`
	PreTaxIncome                                        float64 `json:"pre_tax_income"`
	TexProvision                                        float64 `json:"tex_provision"`
	NetIncomeCommonStockholders                         float64 `json:"net_income_common_stockholders"`
	NetIncome                                           float64 `json:"net_income"`
	NetIncomeIncludingNonControllingInterest            float64 `json:"net_income_including_non_controlling_interest"`
	NetIncomeContinousOperations                        float64 `json:"net_income_continous_operations"`
	DilutedNiAvailableToCommonStockholders              float64 `json:"diluted_nia_available_to_common_stockholders"`
	BasicEps                                            float64 `json:"basic_eps"`
	DilutedEps                                          float64 `json:"diluted_eps"`
	BasicAverageShares                                  float64 `json:"basic_average_shares"`
	DilutedAverageshares                                float64 `json:"diluted_averageshares"`
	TotalOperatingIncomeAsReported                      float64 `json:"total_operating_income_as_reported"`
	TotalExpenses                                       float64 `json:"total_expenses"`
	NetIncomeFromContiuingAndDiscontinuedOperation      float64 `json:"net_income_from_contiuing_and_discontinued_operation"`
	NormalizedIncome                                    float64 `json:"normalized_income"`
	InterestIncome                                      float64 `json:"interest_income"`
	InterestExpense                                     float64 `json:"interest_expense"`
	NetInterestIncome                                   float64 `json:"net_interest_income"`
	Ebit                                                float64 `json:"ebit"`
	Ebitda                                              float64 `json:"ebitda"`
	ReconciledCostOfRevenue                             float64 `json:"reconciled_cost_of_revenue"`
	ReconciledDepreciation                              float64 `json:"reconciled_depreciation"`
	NetIncomeFromContinuingOperationNetMinorityInterest float64 `json:"net_income_from_continuing_operation_net_minority_interest"`
	TotalUnusualItemsExcludingGoodwill                  float64 `json:"total_unusual_items_excluding_goodwill"`
	TotalUnusualItems                                   float64 `json:"total_unusual_items"`
	NormalizedEditda                                    float64 `json:"normalized_ebitda"`
	TaxRateForCalcs                                     float64 `json:"tax_rate_for_calcs"`
	TexEffectOfUnusualItems                             float64 `json:"tex_effect_of_unusual_items"`
}
