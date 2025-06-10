package models

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/utils"
	"go.uber.org/zap"
)

type Stock struct {
	Name            string  `binding:"required" json:"name"`
	Symbol          string  `gorm:"primaryKey" binding:"required" json:"stock_symbol"`
	IpoYear         int16   `json:"ipo_year"`
	Country         string  `binding:"required" json:"country"`
	CurrentPrice    float64 `json:"current_price"`
	PriceChange     float64 `json:"price_change"`
	ChangePercent   float32 `json:"change_percent"`
	OpenPrice       float64 `json:"open_price"`
	DayRange        float64 `json:"day_range"`
	DayLow          float64 `json:"day_low"`
	DayHigh         float64 `json:"day_high"`
	Volume          int64   `json:"volume"`
	LatestTradeTime string  `json:"latest_trade_time"`
	Ticker          string  `json:"ticker"`
	Exchange        string  `json:"exchange"`
	Industry        string  `binding:"required" json:"industry"`
	Sector          string  `binding:"required" json:"sector"`
	Employees       int32   `json:"employees"`
	Headquarters    string  `json:"headquarters"`
	MarketCap       float64 `json:"market_cap"`
	PERatioTtm      float32 `json:"pe_ratio_ttm"`
	EPSTtm          float32 `json:"eps_ttm"`
}

func (stock *Stock) GetFinancialIndex(wg *sync.WaitGroup) {
	defer wg.Done()

	err := godotenv.Load()
	utils.OnLogError(err, "can not load environment variables")

	url := os.Getenv("STOCK_RESOURCE_URL") + stock.Symbol

	c := colly.NewCollector(
		colly.AllowedDomains(os.Getenv("STOCK_BASED_URL")),
		colly.MaxRequests(5),
	)

	//CurrentPrice
	//PriceChange
	//ChangePercent
	c.OnHTML("span[data-testid]", func(e *colly.HTMLElement) {
		switch e.Attr("data-testid") {
		case "qsp-price":
			currentPrice, err := strconv.ParseFloat(strings.TrimSpace(e.Text), 64)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not parse current price: %v", stock.Symbol, e.Text))
			stock.CurrentPrice = currentPrice
		case "qsp-price-change":
			changePrice, err := strconv.ParseFloat(utils.RemoveSpecialSymbol(e.Text), 64)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not parse change price: %v", stock.Symbol, e.Text))
			stock.PriceChange = changePrice
		case "qsp-price-change-percent":
			changePercent, err := strconv.ParseFloat(utils.RemoveSpecialSymbol(e.Text), 32)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not parse change percent: %v", stock.Symbol, e.Text))
			stock.ChangePercent = float32(changePercent)
		}
	})

	//OpenPrice
	//DayRange
	//DayLow
	//DayHigh
	//MarketCap
	//Volume
	c.OnHTML("fin-streamer", func(e *colly.HTMLElement) {
		switch e.Attr("data-field") {
		case "regularMarketOpen":
			openPrice, err := strconv.ParseFloat(strings.TrimSpace(e.Text), 64)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not parse open price: %v", stock.Symbol, e.Text))
			stock.OpenPrice = openPrice
		case "regularMarketDayRange":
			dayRangeArr := strings.Split(e.Text, " ")

			dayLow, err := strconv.ParseFloat(utils.RemoveSpecialSymbol(dayRangeArr[0]), 64)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not parse day low: %v", stock.Symbol, dayRangeArr[0]))

			dayHigh, err := strconv.ParseFloat(utils.RemoveSpecialSymbol(dayRangeArr[2]), 64)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not parse day high: %v", stock.Symbol, dayRangeArr[2]))

			stock.DayHigh = dayHigh
			stock.DayLow = dayLow
			stock.DayRange = dayHigh - dayLow
		case "regularMarketVolume":
			volume, err := strconv.ParseInt(utils.RemoveSpecialSymbol(e.Text), 10, 64)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not parse voume: %v", stock.Symbol, e.Text))
			stock.Volume = volume
		}
	})

	//PE
	//EPS
	c.OnHTML("li", func(e *colly.HTMLElement) {
		label := e.ChildText("span[title]")
		value := e.ChildText("fin-streamer")

		switch label {
		case "PE Ratio (TTM)":
			pe, err := strconv.ParseFloat(strings.TrimSpace(value), 32)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not pe ratio: %v", stock.Symbol, value))
			stock.PERatioTtm = float32(pe)
		case "EPS (TTM)":
			eps, err := strconv.ParseFloat(strings.TrimSpace(value), 32)
			utils.OnLogError(err, fmt.Sprintf("[%v] can not eps: %v", stock.Symbol, value))
			stock.EPSTtm = float32(eps)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		zap.L().Info("Get new data",
			zap.String("Symbol", stock.Symbol),
			zap.String("URL", url),
		)
	})

	c.OnError(func(r *colly.Response, err error) {
		zap.L().Error("Request Error",
			zap.String("Symbol", stock.Symbol),
			zap.String("URL", r.Request.URL.String()),
		)
		r.Request.Retry()
	})

	c.Visit(url)
}
