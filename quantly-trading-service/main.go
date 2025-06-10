package main

import (
	"fmt"
	"os"

	"github.com/trmviet0801/quantly/data"
	"github.com/trmviet0801/quantly/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	// Configure the encoder to use colors
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Colors for levels (e.g., ERROR in red)
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // Optional: format time

	// Create a core with console output
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),             // Lock for concurrent safety
		zap.NewAtomicLevelAt(zap.InfoLevel), // Set the minimum log level
	)

	// Replace the global logger
	logger := zap.New(core, zap.AddCaller()) // Add caller info for debugging
	zap.ReplaceGlobals(logger)
}

func main() {
	var stocks []*models.Stock = data.GetStocksFinancialIndexes()
	for _, stock := range stocks {
		fmt.Println(stock)
	}
}
