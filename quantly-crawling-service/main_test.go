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

	// snapshotId, err := usecase.TriggerFullStockCrawl()
	// if err != nil {
	// 	fmt.Println("failed")
	// } else {
	// 	fmt.Println(snapshotId)
	// }

	// snapshot, _ := usecase.GetProcessStatus("s_mc4r3u432je0vzrco1")
	// fmt.Println(snapshot.String())

	// ctx := context.Background()

	// db := db.Database{}
	// rdb := db.GetDatabase()

	// rdb.Set(ctx, "foo", "bar", 0)
	// result, err := rdb.Get(ctx, "foo").Result()
	// if err != nil {
	// 	fmt.Println("error")
	// } else {
	// 	fmt.Println(result)
	// }

	snapshotOverviews, err := usecase.GetAllSnapshotOverview()
	if err != nil {
		fmt.Println("failed")
	} else {
		for _, snapshot := range snapshotOverviews {
			fmt.Println(snapshot.String())
		}
	}
}
