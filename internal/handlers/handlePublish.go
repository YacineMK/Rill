package handlers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/YacineMK/Rill/internal/config"
	"github.com/YacineMK/Rill/pkg/jwt"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/nareix/joy4/format/ts"
)

var cfg = config.GetConfig()

func HandlePublish(conn *rtmp.Conn) {
	path := conn.URL.Path
	if path == "" {
		log.Println("No stream path")
		conn.Close()
		return
	}

	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		log.Println("Invalid path")
		conn.Close()
		return
	}

	claims, err := jwt.DecodeJWT(parts[1])
	if err != nil {
		log.Println("Invalid key:", err)
		conn.Close()
		return
	}

	streamKey := claims.StreamID
	streamDir := filepath.Join(cfg.RTMP.TmpFolderPath, streamKey)

	if err := os.MkdirAll(streamDir, 0755); err != nil {
		log.Printf("Dir error: %v", err)
		conn.Close()
		return
	}

	log.Printf("Stream: %s", streamKey)

	codecs, err := conn.Streams()
	if err != nil {
		log.Printf("Error getting streams: %v", err)
		conn.Close()
		return
	}

	segmentIndex := 0
	segmentDuration := 10 * time.Second
	segmentStartTime := time.Now()

	var currentFile *os.File
	var tsMuxer *ts.Muxer

	for {
		pkt, err := conn.ReadPacket()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		if time.Since(segmentStartTime) >= segmentDuration || currentFile == nil {
			if currentFile != nil {
				currentFile.Close()
			}

			segmentName := fmt.Sprintf("segment_%d.ts", segmentIndex)
			segmentPath := filepath.Join(streamDir, segmentName)

			currentFile, err = os.Create(segmentPath)
			if err != nil {
				log.Printf("Error creating segment: %v", err)
				break
			}

			tsMuxer = ts.NewMuxer(currentFile)
			if err := tsMuxer.WriteHeader(codecs); err != nil {
				log.Printf("Error writing header: %v", err)
				break
			}

			playlistPath := filepath.Join(streamDir, "playlist.m3u8")
			updatePlaylist(playlistPath, segmentName, segmentIndex)

			log.Printf("Created segment: %s", segmentName)
			segmentIndex++
			segmentStartTime = time.Now()
		}

		if tsMuxer != nil {
			err = tsMuxer.WritePacket(pkt)
			if err != nil {
				log.Printf("Error writing packet: %v", err)
			}
		}
	}

	if currentFile != nil {
		currentFile.Close()
	}

	log.Printf("Stream ended: %s", streamKey)
}

func updatePlaylist(playlistPath, segmentName string, segmentIndex int) {
	content := "#EXTM3U\n#EXT-X-VERSION:3\n#EXT-X-TARGETDURATION:10\n#EXT-X-MEDIA-SEQUENCE:0\n"
	for i := 0; i <= segmentIndex; i++ {
		content += fmt.Sprintf("#EXTINF:10.0,\nsegment_%d.ts\n", i)
	}
	os.WriteFile(playlistPath, []byte(content), 0644)
}
