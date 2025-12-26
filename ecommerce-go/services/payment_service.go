package services

import "errors"

func ProcessPayment(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid payment amount")
	}

	return nil
}
