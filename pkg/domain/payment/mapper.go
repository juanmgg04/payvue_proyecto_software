package payment

import "math"

func ToPaymentResponse(pwd PaymentWithDebt) PaymentResponse {
	remainingInstallments := 0
	if pwd.DebtInstallmentAmount > 0 {
		remainingInstallments = int(math.Floor(pwd.DebtRemainingAmount / pwd.DebtInstallmentAmount))
	}

	receiptURL := ""
	if pwd.ReceiptFilename != "" {
		receiptURL = "/api/v1/payments/receipt/" + pwd.ReceiptFilename
	}

	return PaymentResponse{
		ID:                    pwd.ID,
		Amount:                pwd.Amount,
		Date:                  pwd.Date.Format("2006-01-02"),
		DebtName:              pwd.DebtName,
		RemainingInstallments: remainingInstallments,
		RemainingAmount:       pwd.DebtRemainingAmount,
		ReceiptURL:            receiptURL,
	}
}

func ToPaymentListResponse(payments []PaymentWithDebt) PaymentListResponse {
	responses := make([]PaymentResponse, len(payments))
	for i, payment := range payments {
		responses[i] = ToPaymentResponse(payment)
	}

	return PaymentListResponse{
		Payments: responses,
	}
}
