package mylog

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// https://pkg.go.dev/go.uber.org/zap
// https://github.com/sandipb/zap-examples
// https://gist.github.com/calam1/c2673b6b0a53918df033d71bbf958b56
// https://github.com/bigwhite/experiments/tree/master/uber-zap-advanced-usage
// COLOR https://github.com/uber-go/zap/issues/648

func ExampleFromDoc() {
	// Using zap's preset constructors is the simplest way to get a feel for the
	// package, but they don't allow much customization.
	logger, _ := zap.NewProduction() // or NewProduction, or NewDevelopment, or NewExample
	defer logger.Sync()

	const url = "http://example.com"

	// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	// other structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	// In the unusual situations where every microsecond matters, use the
	// Logger. It's even faster than the SugaredLogger, but only supports
	// structured logging.
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func ExampleWithColor() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("example")
}
