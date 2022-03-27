package service

import "github.com/mazzoleni-gabriel/transaction-risk/risk"

const (
	lowRisk    = 0
	mediumRisk = 1
	highRisk   = 2

	mediumRiskTransactionSpend = 500000
	highRiskTransactionSpend   = 1000000

	mediumRiskTotalSpend = 1000000
	highRiskTotalSpend   = 2000000

	mediumRiskCreditCards = 1
	highRiskCreditCards   = 2
)

func RateRisks(request risk.TransactionsRequest) risk.RiskRatingsResponse {
	userTotalSpendMap := make(map[int]int)
	userCardsMap := make(map[int]map[int]bool)

	var transactionRisks []int

	for _, transaction := range request.Transactions {
		userTotalSpendMap[transaction.UserID] += transaction.AmountUsCents

		cardsMap := userCardsMap[transaction.UserID]
		if cardsMap == nil {
			userCardsMap[transaction.UserID] = make(map[int]bool)
		}
		userCardsMap[transaction.UserID][transaction.CardID] = true

		transactionRisk := resolveTransactionRisk(transaction, userTotalSpendMap, userCardsMap)
		transactionRisks = append(transactionRisks, transactionRisk)
	}

	return buildRiskRatingsResponse(transactionRisks)
}

func buildRiskRatingsResponse(transactionRisks []int) risk.RiskRatingsResponse {
	riskTranslations := map[int]string{
		lowRisk:    "low",
		mediumRisk: "medium",
		highRisk:   "high",
	}

	response := risk.RiskRatingsResponse{}

	for _, transactionRisk := range transactionRisks {
		response.RiskRatings = append(response.RiskRatings, riskTranslations[transactionRisk])
	}

	return response
}

func resolveTransactionRisk(transaction risk.Transaction, spendMap map[int]int, cardsMap map[int]map[int]bool) int {
	risk := 0

	if transaction.AmountUsCents > mediumRiskTransactionSpend && risk < mediumRisk {
		risk = mediumRisk
	}

	if transaction.AmountUsCents > highRiskTransactionSpend && risk < highRisk {
		risk = highRisk
	}

	if spendMap[transaction.UserID] > mediumRiskTotalSpend && risk < mediumRisk {
		risk = mediumRisk
	}

	if spendMap[transaction.UserID] > highRiskTotalSpend && risk < highRisk {
		risk = highRisk
	}

	if len(cardsMap[transaction.UserID]) > mediumRiskCreditCards && risk < mediumRisk {
		risk = mediumRisk
	}

	if len(cardsMap[transaction.UserID]) > highRiskCreditCards && risk < highRisk {
		risk = highRisk
	}

	return risk
}
