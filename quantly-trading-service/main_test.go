package main

import (
	"os"
	"testing"

	"github.com/trmviet0801/quantly/usecase"
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

func TestSubmitOrder(t *testing.T) {
	//automate.AutomateController()

	// order := dto.OrderDto{
	// 	Symbol:        "",
	// 	Qty:           "1",
	// 	Side:          "buy",
	// 	Type:          "limit", // e.g. "market", "limit"
	// 	TimeInForce:   "gtc",   // e.g. "gtc", "day", "opg"
	// 	LimitPrice:    "10.5",
	// 	ExtendedHours: false,
	// 	OrderClass:    "simple",
	// }

	// usecase.SubmitOrder(&order, "12c5d20e-aa3d-412b-985e-245a927a1be4")

	// data, err := usecase.GetAllOrdersOfAccount("da4337b4-1f79-4427-a47d-2f2044be6402")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(len(*data))
	// 	for _, item := range *data {
	// 		fmt.Println(item.String() + "\n")
	// 	}
	// }

	usecase.CancelOrder("12c5d20e-aa3d-412b-985e-245a927a1be4", "18315a8-44dc-4a39-a934-0da5c7696136")
}
