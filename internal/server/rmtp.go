package server

import (
	"log"

	"github.com/nareix/joy4/format/rtmp"
)

type RtmpServer struct {
	Port       string
	rtmpServer *rtmp.Server
}

func NewRtmpServer(port string) *RtmpServer {
	s := &RtmpServer{
		Port:       port,
		rtmpServer: &rtmp.Server{},
	}

	return s
}

func (s *RtmpServer) Start() error {
	log.Println("📡 RTMP server listening on", s.Port)
	s.rtmpServer.Addr = ":" + s.Port
	return s.rtmpServer.ListenAndServe()
}
