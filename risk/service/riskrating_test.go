package service_test

import (
	"github.com/mazzoleni-gabriel/transaction-risk/risk"
	"github.com/mazzoleni-gabriel/transaction-risk/risk/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RateRisks(t *testing.T) {
	t.Run("Should return expected risks", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 200000, CardID: 1},
				{ID: 2, UserID: 1, AmountUsCents: 600000, CardID: 1},
				{ID: 3, UserID: 1, AmountUsCents: 1100000, CardID: 1},
				{ID: 4, UserID: 2, AmountUsCents: 100000, CardID: 2},
				{ID: 5, UserID: 2, AmountUsCents: 100000, CardID: 3},
				{ID: 6, UserID: 2, AmountUsCents: 100000, CardID: 4},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"low",
				"medium",
				"high",
				"low",
				"medium",
				"high",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})
}

func Test_TransactionAmount(t *testing.T) {
	t.Run("Should return low risk when amount is less  or equal than $5000", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"low",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Should return medium risk when amount is more than $5000 and less than $10000", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 500001, CardID: 1},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"medium",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Should return HIGH risk when amount is more than $10000", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 1000001, CardID: 1},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"high",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})
}

func Test_TransactionsTotalAmount(t *testing.T) {
	t.Run("Should return low risk when total amount is less or equal than $10000", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"low",
				"low",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Should return medium risk when total amount is more than $10000 and less than $20000", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"low",
				"low",
				"medium",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Should return high risk when total amount is more than $20000", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 500000, CardID: 1},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"low",
				"low",
				"medium",
				"medium",
				"high",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})
}

func Test_TransactionsTotalCards(t *testing.T) {
	t.Run("Should return low risk when total cards is less than 2", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 1, CardID: 1},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"low",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Should return medium risk when total cards is equal to 2", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 1, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 1, CardID: 2},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"low",
				"medium",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Should return high risk when total cards is more than 2", func(t *testing.T) {

		input := risk.TransactionsRequest{
			Transactions: []risk.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 1, CardID: 1},
				{ID: 1, UserID: 1, AmountUsCents: 1, CardID: 2},
				{ID: 1, UserID: 1, AmountUsCents: 1, CardID: 3},
			},
		}

		expectedResponse := risk.RiskRatingsResponse{
			RiskRatings: []string{
				"low",
				"medium",
				"high",
			},
		}

		res := service.RateRisks(input)

		assert.Equal(t, expectedResponse, res)
	})
}
