package services

import (
	"errors"
	"time"

	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func ValidateCoupon(code string, amount float64) (*models.Coupon, float64, error) {
	coupon, err := repositories.GetCouponByCode(code)
	if err != nil {
		return nil, 0, errors.New("coupon not found")
	}

	if coupon.EndDate.Before(time.Now()) {
		return nil, 0, errors.New("coupon expired")
	}

	if coupon.StartDate.After(time.Now()) {
		return nil, 0, errors.New("coupon not yet active")
	}

	if coupon.MaxUsage > 0 && coupon.UsageCount >= coupon.MaxUsage {
		return nil, 0, errors.New("coupon usage limit reached")
	}

	if amount < coupon.MinAmount {
		return nil, 0, errors.New("purchase amount does not meet minimum")
	}

	var discount float64
	if coupon.DiscountType == "percentage" {
		discount = (amount * coupon.DiscountValue) / 100
	} else {
		discount = coupon.DiscountValue
	}

	return coupon, discount, nil
}

func ApplyCoupon(code string) error {
	return repositories.UpdateCouponUsage(code)
}
