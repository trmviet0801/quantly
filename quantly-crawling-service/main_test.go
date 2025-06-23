package quantlycrawlingservice

import (
	"os"
	"testing"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/usecase"
)

func Test(t *testing.T) {
	result, err := usecase.CrawlStockInfo("s_mc6d8baz2dz62ln10s")
	if err != nil {
		panic("huhu")
	}
	for _, stock := range result {
		os.WriteFile("./test.txt", []byte(stock.String()), 0777)
	}
}
