package main

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/usecase"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

func Test(t *testing.T) {
	// result, err := usecase.CrawlStockInfo("s_mc6d8baz2dz62ln10s")
	// if err != nil {
	// 	panic("huhu")
	// }
	// for _, stock := range result {
	// 	os.WriteFile("./test.txt", []byte(stock.String()), 0777)
	// }

	err := godotenv.Load()
	if err != nil {
		utils.OnError(err)
		return
	}

	snapshotId, err := usecase.TriggerFullStockCrawl()
	if err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println(snapshotId)
	}
}
