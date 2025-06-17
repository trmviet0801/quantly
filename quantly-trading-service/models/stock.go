package models

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/utils"
	"go.uber.org/zap"
)

type Stock struct {
	Name            string    `gorm:"column:name" binding:"required"`
	Symbol          string    `gorm:"primaryKey;column:stock_symbol" binding:"required"`
	IpoYear         int16     `gorm:"column:ipo_year"`
	Country         string    `gorm:"column:country" binding:"required"`
	CurrentPrice    float64   `gorm:"column:current_price"`
	PriceChange     float64   `gorm:"column:price_change"`
	ChangePercent   float32   `gorm:"column:change_percent"`
	OpenPrice       float64   `gorm:"column:open_price"`
	DayRange        float64   `gorm:"column:day_range"`
	DayLow          float64   `gorm:"column:day_low"`
	DayHigh         float64   `gorm:"column:day_high"`
	Volume          int64     `gorm:"column:volume"`
	LatestTradeTime time.Time `gorm:"column:latest_trade_time"`
	Ticker          string    `gorm:"column:ticker"`
	Exchange        string    `gorm:"column:exchange"`
	Industry        string    `gorm:"column:industry" binding:"required"`
	Sector          string    `gorm:"column:sector" binding:"required"`
	Employees       int32     `gorm:"column:employees"`
	Headquarters    string    `gorm:"column:headquarters"`
	MarketCap       float64   `gorm:"column:market_cap"`
	PERatioTtm      float32   `gorm:"column:pe_ratio_ttm"`
	EPSTtm          float32   `gorm:"column:eps_ttm"`
}

func (stock *Stock) GetFinancialIndex(wg *sync.WaitGroup) {
	defer wg.Done()

	err := godotenv.Load()
	utils.OnLogError(err, "can not load environment variables")

	c := colly.NewCollector(
		colly.AllowedDomains(os.Getenv("STOCK_BASED_URL")),
		colly.MaxRequests(20),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "yahoo.",
		Parallelism: 1,
		Delay:       5 * time.Second,
		RandomDelay: 10 * time.Second,
	})

	userAgents := []string{
		// Chrome on Windows
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",

		// Chrome on macOS
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 13_4_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",

		// Firefox on Windows
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:125.0) Gecko/20100101 Firefox/125.0",

		// Firefox on macOS
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 13.4; rv:125.0) Gecko/20100101 Firefox/125.0",

		// Safari on macOS
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 13_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",

		// Edge on Windows
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.51",
		// iPhone Safari
		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",

		// Android Chrome
		"Mozilla/5.0 (Linux; Android 13; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Mobile Safari/537.36",

		// Samsung Browser
		"Mozilla/5.0 (Linux; Android 13; SAMSUNG SM-G991B) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/22.0 Chrome/124.0.0.0 Mobile Safari/537.36",
	}

	//url := os.Getenv("STOCK_RESOURCE_URL") + stock.Symbol

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
			zap.String("URL", os.Getenv("STOCK_RESOURCE_URL")+stock.Symbol),
		)
		r.Headers.Add("User-Agent", userAgents[rand.Intn(len(userAgents)-1)+1])
	})

	c.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == 429 {
			zap.L().Error("IP Banned",
				zap.String("Symbol", stock.Symbol),
				zap.String("URL", r.Request.URL.String()),
			)
		} else {
			zap.L().Error("Error",
				zap.String("Symbol", stock.Symbol),
				zap.Int("", r.StatusCode),
				zap.String("URL", r.Request.URL.String()),
			)
		}
		r.Request.Retry()
	})
	c.Visit(os.Getenv("STOCK_RESOURCE_URL") + stock.Symbol)
}
