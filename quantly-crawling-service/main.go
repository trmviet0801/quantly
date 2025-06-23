package main

import (
	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		utils.OnError(err)
		return
	}
}
