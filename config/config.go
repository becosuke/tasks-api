package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	nowFixation    bool
	nowDatetime    time.Time
	nowTimestamp   int64
	TasksEnv       string
	Memcache       Memcache
	DatabaseMaster Database
	DatabaseSlave  Database
	GrpcHost       string
	GrpcPort       int
	GrpcAddr       string
	RestPort       int
	RestAddr       string
}

type Memcache struct {
	Host []string
	Port int
	Url  []string
}

type Database struct {
	User string
	Pass string
	Host []string
	Port int
	Url  []string
}

var conf *Config

func init() {
	location := "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	var now time.Time
	var nowFixation bool
	if tmp := os.Getenv("TASKS_NOW_DATETIME"); tmp == "" {
		now = time.Now()
		nowFixation = false
	} else {
		if parsed, err := time.Parse("2006-01-02 15:04:05", tmp); err != nil {
			now = time.Now()
			nowFixation = false
		} else {
			now = parsed
			nowFixation = true
		}
	}

	var tasksEnv string
	if tasksEnv = os.Getenv("TASKS_ENV"); tasksEnv == "" {
		tasksEnv = "local"
	}

	var memcacheHost []string
	if memcacheHost = strings.Split(os.Getenv("TASKS_MEMCACHE_HOST"), ","); len(memcacheHost) == 0 {
		memcacheHost = []string{"localhost"}
	}

	var memcachePort int
	if memcachePort, err = strconv.Atoi(os.Getenv("TASKS_MEMCACHE_PORT")); err != nil {
		memcachePort = 11211
	}

	memcacheUrl := make([]string, 0, len(memcacheHost))
	for _, val := range memcacheHost {
		memcacheUrl = append(memcacheUrl, fmt.Sprintf("%s:%d", val, memcachePort))
	}

	var databaseUser string
	if databaseUser = os.Getenv("TASKS_DATABASE_USER"); databaseUser == "" {
		databaseUser = "root"
	}

	var databasePass string
	if databasePass = os.Getenv("TASKS_DATABASE_PASS"); databasePass == "" {
		databasePass = "root"
	}

	var databaseHostMaster []string
	if databaseHostMaster = strings.Split(os.Getenv("TASKS_DATABASE_HOST_MASTER"), ","); len(databaseHostMaster) == 0 {
		databaseHostMaster = []string{"localhost"}
	}

	var databaseHostSlave []string
	if databaseHostSlave = strings.Split(os.Getenv("TASKS_DATABASE_HOST_SLAVE"), ","); len(databaseHostSlave) == 0 {
		databaseHostSlave = []string{"localhost"}
	}

	var databasePort int
	if databasePort, err = strconv.Atoi(os.Getenv("TASKS_DATABASE_PORT")); err != nil {
		databasePort = 3306
	}

	databaseUrlMaster := make([]string, 0, len(databaseHostMaster))
	for _, val := range databaseHostMaster {
		databaseUrlMaster = append(
			databaseUrlMaster,
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%%s?charset=utf8mb4&parseTime=true&loc=%s",
				databaseUser,
				databasePass,
				val,
				databasePort,
				"Asia%%2FTokyo",
			),
		)
	}

	databaseUrlSlave := make([]string, 0, len(databaseHostSlave))
	for _, val := range databaseHostSlave {
		databaseUrlSlave = append(
			databaseUrlSlave,
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%%s?charset=utf8mb4&parseTime=true&loc=%s",
				databaseUser,
				databasePass,
				val,
				databasePort,
				"Asia%%2FTokyo",
			),
		)
	}

	var grpcHost string
	if grpcHost = os.Getenv("TASKS_GRPC_HOST"); grpcHost == "" {
		grpcHost = "0.0.0.0"
	}

	var grpcPort int
	if grpcPort, err = strconv.Atoi(os.Getenv("TASKS_GRPC_PORT")); err != nil {
		grpcPort = 50051
	}

	var restPort int
	if restPort, err = strconv.Atoi(os.Getenv("TASKS_REST_PORT")); err != nil {
		restPort = 80
	}

	conf = &Config{
		nowFixation:  nowFixation,
		nowDatetime:  now,
		nowTimestamp: now.Unix(),
		TasksEnv:     tasksEnv,
		Memcache: Memcache{
			Host: memcacheHost,
			Port: memcachePort,
			Url:  memcacheUrl,
		},
		DatabaseMaster: Database{
			User: databaseUser,
			Pass: databasePass,
			Host: databaseHostMaster,
			Port: databasePort,
			Url:  databaseUrlMaster,
		},
		DatabaseSlave: Database{
			User: databaseUser,
			Pass: databasePass,
			Host: databaseHostSlave,
			Port: databasePort,
			Url:  databaseUrlSlave,
		},
		GrpcHost: grpcHost,
		GrpcPort: grpcPort,
		GrpcAddr: grpcHost + ":" + strconv.Itoa(grpcPort),
		RestPort: restPort,
		RestAddr: ":" + strconv.Itoa(restPort),
	}

}

func GetConfig() *Config {
	return conf
}

func NowDatetime() time.Time {
	if GetConfig().nowFixation {
		return GetConfig().nowDatetime
	}

	return time.Now()
}

func NowTimestamp() int64 {
	if GetConfig().nowFixation {
		return GetConfig().nowTimestamp
	}

	return time.Now().Unix()
}
