package main

import (
	"os"

	"github.com/trmviet0801/quantly/automate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	// Configure the encoder to use colors
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Colors for levels (e.g., ERROR in red)
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // Optional: format time

	// // Create a core with console output
	// core := zapcore.NewCore(
	// 	zapcore.NewConsoleEncoder(encoderConfig),
	// 	zapcore.Lock(os.Stdout),             // Lock for concurrent safety
	// 	zap.NewAtomicLevelAt(zap.InfoLevel), // Set the minimum log level
	// )

	// Log file
	logFile, err := os.OpenFile("./log/system_log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Fail to open log file")
	}
	fileWriter := zapcore.AddSync(logFile)
	consoleWriter := zapcore.AddSync(os.Stdout)

	// Encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// Level enablers
	consoleLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel // log INFO and above to console
	})
	fileLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel // only log ERROR and above to file
	})

	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, consoleLevel)
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, fileLevel)

	core := zapcore.NewTee(consoleCore, fileCore)

	// Replace the global logger
	logger := zap.New(core, zap.AddCaller()) // Add caller info for debugging
	zap.ReplaceGlobals(logger)
}

func main() {
	// var stocks []*models.Stock = data.GetStocksFinancialIndexes()
	// for _, stock := range stocks {
	// 	fmt.Println(stock)
	// }
	automate.AutomateController()
}
