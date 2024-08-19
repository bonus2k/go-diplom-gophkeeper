package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

var (
	srvAddr  string
	logLevel string
	dbFile   string
	crtFile  string
	confFile string
)

type config struct {
	SrvAddr  string `json:"conn_addr"`
	LogLevel string `json:"log_level"`
	CrtFile  string `json:"crt_file"`
	DBFile   string `json:"log_file"`
}

func parseFlags() {
	var cfg *config
	flag.StringVar(&confFile, "cfg", "config_s.cfg", "config file path")
	cfg, err := parseConfig(confFile)
	if err != nil {
		flag.StringVar(&srvAddr, "a", "localhost:3200", "server address")
		flag.StringVar(&logLevel, "ll", "info", "log level")
		flag.StringVar(&dbFile, "db", "test.db", "db path")
		flag.StringVar(&crtFile, "crt", "private.pem", "certificate x509 path")
		cfg = &config{
			SrvAddr:  srvAddr,
			LogLevel: logLevel,
			DBFile:   dbFile,
			CrtFile:  crtFile,
		}
	} else {
		srvAddr = cfg.SrvAddr
		logLevel = cfg.LogLevel
		dbFile = cfg.DBFile
		crtFile = cfg.CrtFile
	}

	saveCfg(confFile, cfg)
}

func saveCfg(confFile string, cfg *config) {
	file, err := os.OpenFile(confFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	bytes, err := json.Marshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(bytes)
}

func parseConfig(confFile string) (*config, error) {
	var conf config
	file, err := os.Open(confFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buf, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
