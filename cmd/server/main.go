package main

import (
	"log"

	"github.com/YacineMK/Rill/internal/config"
	"github.com/YacineMK/Rill/internal/server"
)

var cfg = config.GetConfig()

func main() {
	rmtpSrv := server.NewRtmpServer(cfg.RTMP.Port)
	go func() {
		if err := rmtpSrv.Start(); err != nil {
			log.Fatal("RTMP server error:", err)
		}
	}()

	httpSrv := server.NewHttpServer(cfg.HTTP.Port)
	if err := httpSrv.Start(); err != nil {
		log.Fatal("HTTP server error:", err)
	}
}
