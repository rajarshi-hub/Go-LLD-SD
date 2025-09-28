package ParkingPricing

import (
	"fmt"
	"time"
)

// Prices in paise
type ParkingPricing interface {
	CalculatePrice(inTime, outTime *time.Time) int
}

type BikeParkingPricing struct {
	BasePrice    int
	PricePerHour int
}

func (bp *BikeParkingPricing) CalculatePrice(inTime, outTime *time.Time) int {
	dur := outTime.Sub(*inTime)
	fmt.Println(dur.Hours(), dur.Hours() < 0)
	if dur.Hours() < 1 {
		return bp.BasePrice
	}
	return bp.PricePerHour*int(dur.Hours()-1) + bp.PricePerHour

}

type CarParkingPricing struct {
	BasePrice    int
	PricePerHour int
}

func (cp *CarParkingPricing) CalculatePrice(inTime, outTime *time.Time) int {
	dur := outTime.Sub(*inTime)
	if dur.Hours() < 0 {
		return cp.BasePrice
	}
	return cp.PricePerHour*int(dur.Hours()-1) + cp.PricePerHour

}

type LoyaltyPointedParkingPricing struct {
	baseModel     ParkingPricing
	LoyaltyPoints float64 // b/w 0 & 1
}

func (lp *LoyaltyPointedParkingPricing) CalculatePrice(inTime, outTime *time.Time) int {
	if lp.LoyaltyPoints >= 0 {
		return int(float64(lp.baseModel.CalculatePrice(inTime, outTime)) * (1.0 - lp.LoyaltyPoints))
	}
	return lp.baseModel.CalculatePrice(inTime, outTime)
}

type WeekendPricing struct {
	baseModel   ParkingPricing
	WeekendToll int
}

func (lp *WeekendPricing) CalculatePrice(inTime, outTime *time.Time) int {
	return lp.WeekendToll * lp.baseModel.CalculatePrice(inTime, outTime)
}
