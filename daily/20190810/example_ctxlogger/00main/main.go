package main

import (
	"time"

	"github.com/podhmo/ctxlog/zapctxlog"
)

func main() {
	log, _ := zapctxlog.New()
	log.With("now", time.Now()).Info("hai")
}
