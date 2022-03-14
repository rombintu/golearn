package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rombintu/golearn/config"
	"github.com/rombintu/golearn/internal/server"
)

// My code is comments
func main() {
	configPath := flag.String("config", "./config/server.toml", "Path to config file")
	logPath := flag.String("logfile", "./logs/server.log", "Path to log file")
	sqlitePath := flag.String("sqlite", "", "Path to sqlite3 database")
	dev := flag.Bool("dev", false, "Dev mode")

	flag.Parse()

	config := config.GetConfig(*configPath)

	if *dev {
		config.Postgres.Sqlite3 = *sqlitePath
		config.Default.LogFile = *logPath
	}

	s := server.NewApp(config)

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exitCh
		log.Println("Exit with 0")
		s.CloseLogFile()
		s.Store.Close()
		os.Exit(0)
	}()

	if err := s.Start(); err != nil {
		log.Fatalf("%v", err)
	}
}
