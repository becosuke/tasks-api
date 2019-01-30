package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var once sync.Once
var conf *Config

type Config struct {
	NowDatetime  time.Time
	NowTimestamp int64
	TasksEnv     string
	TasksPrefix  string
	Image
	Memcache
	DatabaseMaster Database
	DatabaseSlave  Database
	GrpcHost       string
	GrpcPort       int
	GrpcAddr       string
	RestPort       int
	RestAddr       string
}

type Image struct {
	Protocol string
	Host     string
}

type Memcache struct {
	Host []string
	Port int
	URL  []string
}

type Database struct {
	User string
	Pass string
	Host []string
	Port int
	URL  []string
}

type AWS struct {
	S3 struct {
		Endpoint string
		Region   string
	}
	DynamoDB struct {
		Endpoint string
		Region   string
	}
}

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func GetConfig() *Config {
	once.Do(func() {
		var err error

		var now time.Time
		if tmp := os.Getenv("TASKS_NOW"); tmp == "" {
			now = time.Now()
		} else {
			if parsed, err := time.Parse("2006-01-02 15:04:05", tmp); err != nil {
				now = time.Now()
			} else {
				now = parsed
			}
		}

		var tasksEnv string
		if tasksEnv = os.Getenv("TASKS_ENV"); tasksEnv == "" {
			tasksEnv = "local"
		}

		var tasksPrefix string
		if tasksPrefix = os.Getenv("TASKS_PREFIX"); tasksPrefix == "" {
			tasksPrefix = tasksEnv
		}

		var imageProtocol string
		if imageProtocol = os.Getenv("TASKS_IMAGE_PROTOCOL"); imageProtocol == "" {
			imageProtocol = "https"
		}

		var imageHost string
		if imageHost = os.Getenv("TASKS_IMAGE_HOST"); imageHost == "" {
			imageHost = "localhost"
		}

		var memcacheHost []string
		if memcacheHost = strings.Split(os.Getenv("TASKS_MEMCACHE_HOST"), ","); len(memcacheHost) == 0 {
			memcacheHost = []string{"localhost"}
		}

		var memcachePort int
		if memcachePort, err = strconv.Atoi(os.Getenv("TASKS_MEMCACHE_PORT")); err != nil {
			memcachePort = 11211
		}

		memcacheURL := make([]string, 0, len(memcacheHost))
		for _, val := range memcacheHost {
			memcacheURLTmp := fmt.Sprintf("%s:%d", val, memcachePort)
			memcacheURL = append(memcacheURL, memcacheURLTmp)
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

		databaseURLMaster := make([]string, 0, len(databaseHostMaster))
		for _, val := range databaseHostMaster {
			databaseURLMaster = append(
				databaseURLMaster,
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

		databaseURLSlave := make([]string, 0, len(databaseHostSlave))
		for _, val := range databaseHostSlave {
			databaseURLSlave = append(
				databaseURLSlave,
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
			NowDatetime:  now,
			NowTimestamp: now.Unix(),
			TasksEnv:     tasksEnv,
			TasksPrefix:  tasksPrefix,
			Image: Image{
				Protocol: imageProtocol,
				Host:     imageHost,
			},
			Memcache: Memcache{
				Host: memcacheHost,
				Port: memcachePort,
				URL:  memcacheURL,
			},
			DatabaseMaster: Database{
				User: databaseUser,
				Pass: databasePass,
				Host: databaseHostMaster,
				Port: databasePort,
				URL:  databaseURLMaster,
			},
			DatabaseSlave: Database{
				User: databaseUser,
				Pass: databasePass,
				Host: databaseHostSlave,
				Port: databasePort,
				URL:  databaseURLSlave,
			},
			GrpcHost: grpcHost,
			GrpcPort: grpcPort,
			GrpcAddr: grpcHost + ":" + strconv.Itoa(grpcPort),
			RestPort: restPort,
			RestAddr: ":" + strconv.Itoa(restPort),
		}
	})

	return conf
}
