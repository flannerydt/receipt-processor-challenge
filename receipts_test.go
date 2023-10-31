package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetailerNamePoints(t *testing.T) {
	receipt := Receipt{
		Retailer: "Target",
	}
	result := retailerNamePoints(receipt)
	assert.Equal(t, 6, result)
}

func TestRetailerNamePointsBadChars(t *testing.T) {
	receipt := Receipt{
		Retailer: "Targe*t)0",
	}
	result := retailerNamePoints(receipt)
	assert.Equal(t, 7, result)
}

func TestRoundDollarPoints(t *testing.T) {
	result := roundDollarPoints(50.00)
	assert.Equal(t, 50, result)

}

func TestRoundDollarPointsNotRound(t *testing.T) {
	result := roundDollarPoints(50.21)
	assert.Equal(t, 0, result)
}

func TestQuarterTotalPoints(t *testing.T) {
	result := quarterTotalPoints(0.50)
	assert.Equal(t, 25, result)
}

func TestQuarterTotalPointsRemainder(t *testing.T) {
	result := quarterTotalPoints(0.53)
	assert.Equal(t, 0, result)
}

func TestEveryTwoItemsPointsEven(t *testing.T) {
	item1 := Item{
		Price:       "25.2",
		Description: "Item description",
	}
	item2 := Item{
		Price:       "17.2",
		Description: "Item description",
	}
	var items = []Item{item1, item2}
	result := everyTwoItemsPoints(items)
	assert.Equal(t, 5, result)
}

func TestEveryTwoItemsPointsOdd(t *testing.T) {
	item1 := Item{
		Price:       "25.2",
		Description: "Item description",
	}
	items := []Item{item1}
	result := everyTwoItemsPoints(items)
	assert.Equal(t, 0, result)
}

func TestItemDescriptionPointsSuccess(t *testing.T) {
	item1 := Item{
		Price:       "10.0",
		Description: " Itemdescription   ",
	}
	items := []Item{item1}
	result := itemDescriptionPoints(items)
	assert.Equal(t, 2, result)
}

func TestItemDescriptionPointsSuccessOdd(t *testing.T) {
	item1 := Item{
		Price:       "10.2",
		Description: " Itemdescription   ",
	}
	items := []Item{item1}
	result := itemDescriptionPoints(items)
	assert.Equal(t, 3, result)
}

func TestItemDescriptionPointsFail(t *testing.T) {
	item1 := Item{
		Price:       "10.0",
		Description: " Item description   ",
	}
	items := []Item{item1}
	result := itemDescriptionPoints(items)
	assert.Equal(t, 0, result)
}

func TestOddDayPointsOdd(t *testing.T) {
	purchaseDate := "2023-07-07"
	result := oddDayPoints(purchaseDate)
	assert.Equal(t, 6, result)
}

func TestOddDayPointsEven(t *testing.T) {
	purchaseDate := "2023-07-08"
	result := oddDayPoints(purchaseDate)
	assert.Equal(t, 0, result)
}

func TestPurchaseTimePointsSuccess(t *testing.T) {
	purchaseTime := "14:33"
	result := purchaseTimePoints(purchaseTime)
	assert.Equal(t, 10, result)
}

func TestPurchaseTimePointsFail(t *testing.T) {
	purchaseTime := "08:33"
	result := purchaseTimePoints(purchaseTime)
	assert.Equal(t, 0, result)
}
