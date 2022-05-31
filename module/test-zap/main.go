package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("这是第一个zap log", "url", "hello")
	sugar.Infof("这是第一个zap log %s", "url:hello")
	sugar.Info("这是第一个zap log", zap.String("url", "hello"), zap.Int64("time", time.Now().Unix()))
}
