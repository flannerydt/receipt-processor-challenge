package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Items        []Item `json:"items" binding:"required,dive"`
	Total        string `json:"total" binding:"required"`
}
type Item struct {
	Description string `json:"shortDescription" binding:"required"`
	Price       string `json:"price" binding:"required"`
}
type ID struct {
	ID string `json:"id"`
}
type Points struct {
	Points int `json:"points"`
}

var (
	collection = map[string]*Receipt{}
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/receipts/process", ProcessReceipts).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var receipt Receipt
	json.Unmarshal([]byte(reqBody), &receipt)
	id := ID{
		ID: uuid.New().String(),
	}
	collection[id.ID] = &receipt
	json.NewEncoder(w).Encode(id)
	return
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	var pointsEarned = 0
	vars := mux.Vars(r)
	requestID := vars["id"]
	var receipt Receipt
	receipt = *collection[requestID]

	floatTotal, _ := strconv.ParseFloat(receipt.Total, 32)

	pointsEarned += retailerNamePoints(receipt)
	pointsEarned += roundDollarPoints(floatTotal)
	pointsEarned += quarterTotalPoints(floatTotal)
	pointsEarned += everyTwoItemsPoints(receipt.Items)
	pointsEarned += itemDescriptionPoints(receipt.Items)
	pointsEarned += oddDayPoints(receipt.PurchaseDate)
	pointsEarned += purchaseTimePoints(receipt.PurchaseTime)

	pointsResp := Points{
		Points: pointsEarned,
	}
	json.NewEncoder(w).Encode(pointsResp)
	return
}

// One point for every alphanumeric character in the retailer name
func retailerNamePoints(receipt Receipt) int {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	matches := re.FindAllString(receipt.Retailer, -1)
	var cleanedRetailer = ""
	for _, match := range matches {
		cleanedRetailer += match
	}
	return len(cleanedRetailer)
}

// 50 points if the total is a round dollar amount with no cents
func roundDollarPoints(floatTotal float64) int {
	var dollarAmt int
	dollarAmt = int(floatTotal)
	if float64(dollarAmt) == floatTotal {
		return 50
	}
	return 0
}

// 25 points if the total is a multiple of 0.25
func quarterTotalPoints(floatTotal float64) int {
	remainder := math.Mod(floatTotal, 0.25)
	if remainder == 0 {
		return 25
	}
	return 0
}

// 5 points for every two items on the receipt
func everyTwoItemsPoints(items []Item) int {
	numOfItems := len(items)
	pairs := int(numOfItems / 2)
	return (pairs * 5)
}

// if the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned
func itemDescriptionPoints(items []Item) int {
	var i = 0
	var pointVal = 0
	for i < len(items) {
		item := items[i]
		description := strings.Trim(item.Description, " ")
		if (len(description) % 3) == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points := float64(price) * 0.2
			if math.Trunc(points) == points {
				pointVal += int(points)
			} else {
				pointVal += int(points) + 1
			}
		}
		i += 1
	}
	return pointVal
}

// 6 points if the day in the purchase date is odd
func oddDayPoints(purchaseDate string) int {
	parsedDate := strings.Split(purchaseDate, "-")
	day, _ := strconv.Atoi(parsedDate[2])
	if (day % 2) != 0 {
		return 6
	}
	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func purchaseTimePoints(purchaseTime string) int {
	parsedTime := strings.Split(purchaseTime, ":")
	hour, _ := strconv.Atoi(parsedTime[0])
	if (hour >= 14) && (hour < 16) {
		return 10
	}
	return 0
}
