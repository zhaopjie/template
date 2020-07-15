package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"lottery/models/mysql"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lottery/logger"
	"lottery/server"
	"lottery/setting"
)

func main() {

	// load config and logger
	var cfg setting.Config
	cfg.LoadConf()
	lg := logger.NewLogger(cfg.Log.Level, cfg.Log.OutPath)

	// init mysql
	mysql.LoadMysql(cfg, lg)

	//http server
	sev, err := server.NewHttpServer(lg)
	if err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", ":"+cfg.Core.Port)
	if err != nil {
		panic(err)
	}
	go sev.Start(l)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()
	// Stop the service gracefully.
	if err := sev.ShutDown(ctx); err != nil {
		log.Println(err)
		return
	}

	// Wait gorotine print shutdown message
	//time.Sleep(time.Second * 5)
	log.Println("done.")

}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func hashToken(token string) string {
	hashBytes := sha256.Sum256([]byte(token + "aaaaa"))
	return hex.EncodeToString(hashBytes[:])
}
