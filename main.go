package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	config   *luciferConfig
	version  string
	name     string
	cmdStart = cli.Command{
		Name:   "start",
		Usage:  "you can use file as params ex. ./telegrambot start -e example.env",
		Action: start,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "env-file,e",
			},
			cli.BoolFlag{
				Name:  "debug,d",
				Usage: "open debug mode",
			},
		},
	}
	logger *logrus.Logger
)

func init() {
	logger = logrus.New()
}

type luciferConfig struct {
	// Debug            bool
	// LogDebug         bool
	// RedisIP          string
	// RedisPassword    string
	// RedisDB          int
	// RedisMaxIdle     int
	// RedisMaxActive   int
	// RedisMaxRetries  int
	// RedisTimeout     int
	// RedisPoolSize    int
	// RedisMinIdle     int
	// RedisPoolTimeout int
	// RedisIdleTimeout int
	// RedisDialTimeout int

	// MySQLSlaveHost            string
	// MySQLSlavePort            string
	// MySQLSlaveUserName        string
	// MySQLSlavePassword        string
	// MySQLSlaveDBName          string
	// MySQLSlaveDBParam         string
	// MySQLSlaveMaxIdle         int
	// MySQLSlaveMaxConn         int
	// MySQLSlaveConnMaxLifetime int
	// MySQLSlaveSingularTable   bool
	// MySQLSlaveLogMode         bool

	// MySQLMasterHost            string
	// MySQLMasterPort            string
	// MySQLMasterUserName        string
	// MySQLMasterPassword        string
	// MySQLMasterDBName          string
	// MySQLMasterDBParam         string
	// MySQLMasterMaxIdle         int
	// MySQLMasterMaxConn         int
	// MySQLMasterConnMaxLifetime int
	// MySQLMasterSingularTable   bool
	// MySQLMasterLogMode         bool

	// GrpcListenPort string
	// ReadyPath      string
	// HealthPort     string

	TelegramBool       bool
	TelegramToken      string
	TelegramChatIDInfo string
	TelegramChatIDWarn string
	TelegramAddress    string
}

func envInit(c *cli.Context) (err error) {
	config = &luciferConfig{}
	if c.String("env-file") != "" {
		envLoadErr := godotenv.Load(c.String("env-file"))
		if envLoadErr != nil {
			return envLoadErr
		}
	}
	// config.Debug = c.Bool("debug")
	// logDebug := os.Getenv("LOG_DEBUG")
	// if logDebug == "true" {
	// 	config.LogDebug = true
	// }

	// config.RedisIP = os.Getenv("REDIS_IP")
	// if config.RedisIP == "" {
	// 	err = errors.New("env REDIS_IP empty")
	// 	return
	// }
	// config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	// config.RedisDB, err = strconv.Atoi(os.Getenv("REDIS_DB"))
	// if err != nil {
	// 	err = errors.New("env RedisDB empty")
	// 	return
	// }

	// config.RedisMaxIdle, err = strconv.Atoi(os.Getenv("REDIS_MAX_IDLE"))
	// if err != nil {
	// 	config.RedisMaxIdle = 10
	// }
	// config.RedisMaxActive, err = strconv.Atoi(os.Getenv("REDIS_MAX_ACTIVE"))
	// if err != nil {
	// 	config.RedisMaxActive = 100
	// }
	// config.RedisPoolSize, err = strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	// if err != nil {
	// 	logger.Info("redis pool size empty: ", err)
	// }
	// config.RedisMinIdle, err = strconv.Atoi(os.Getenv("REDIS_MIN_IDLE"))
	// if err != nil {
	// 	logger.Info("env RedisMinIdle empty", err)
	// }
	// config.RedisMaxRetries, err = strconv.Atoi(os.Getenv("REDIS_MAX_RETRIES"))
	// if err != nil {
	// 	logger.Info("env RedisMaxRetries empty", err)
	// }
	// config.RedisDialTimeout, err = strconv.Atoi(os.Getenv("REDIS_DIAL_TIMEOUT"))
	// if err != nil {
	// 	logger.Info("env RedisDialTimeout empty", err)
	// }
	// config.RedisIdleTimeout, err = strconv.Atoi(os.Getenv("REDIS_IDLE_TIMEOUT"))
	// if err != nil {
	// 	logger.Info("env RedisIdleTimeout empty", err)
	// }
	// config.RedisPoolTimeout, err = strconv.Atoi(os.Getenv("REDIS_POOL_TIMEOUT"))
	// if err != nil {
	// 	logger.Info("env RedisPoolTimeout empty", err)
	// }

	// config.GrpcListenPort = os.Getenv("GRPC_LISTEN_PORT")
	// if config.GrpcListenPort == "" {
	// 	err = errors.New("env GrpcListenPort empty")
	// 	return
	// }
	// mysqlSlowTime := os.Getenv("MYSQL_SLOWTIME")
	// if mysqlSlowTime == "" {
	// 	err = errors.New("env MYSQL_SLOWTIME empty")
	// 	return
	// }

	// //mysql master address

	// //gorm env
	// //mysql shit slave address
	// config.MySQLSlaveHost = os.Getenv("MYSQL_SHIT_SLAVE_HOST")
	// if config.MySQLSlaveHost == "" {
	// 	err = errors.New("env MYSQL_SHIT_SLAVE_HOST empty")
	// 	return
	// }
	// config.MySQLSlavePort = os.Getenv("MYSQL_SHIT_SLAVE_PORT")
	// if config.MySQLSlavePort == "" {
	// 	err = errors.New("env MYSQL_SHIT_SLAVE_PORT empty")
	// 	return
	// }
	// config.MySQLSlaveUserName = os.Getenv("MYSQL_SHIT_SLAVE_USERNAME")
	// if config.MySQLSlaveUserName == "" {
	// 	err = errors.New("env MYSQL_SHIT_SLAVE_USERNAME empty")
	// 	return
	// }
	// config.MySQLSlavePassword = os.Getenv("MYSQL_SHIT_SLAVE_PASSWORD")
	// if config.MySQLSlavePassword == "" {
	// 	err = errors.New("env MYSQL_SHIT_SLAVE_PASSWORD empty")
	// 	return
	// }
	// config.MySQLSlaveDBName = os.Getenv("MYSQL_SHIT_SLAVE_DBNAME")
	// if config.MySQLSlaveDBName == "" {
	// 	err = errors.New("env MYSQL_SHIT_SLAVE_DBNAME empty")
	// 	return
	// }
	// config.MySQLSlaveDBParam = os.Getenv("MYSQL_SHIT_SLAVE_DBPARAM")
	// if config.MySQLSlaveDBParam == "" {
	// 	err = errors.New("env MYSQL_SHIT_SLAVE_DBPARAM empty")
	// 	return
	// }
	// config.MySQLSlaveMaxIdle, err = strconv.Atoi(os.Getenv("MYSQL_SHIT_SLAVE_MAXIDLE"))
	// if err != nil {
	// 	config.MySQLSlaveMaxIdle = 10
	// 	return
	// }
	// config.MySQLSlaveMaxConn, err = strconv.Atoi(os.Getenv("MYSQL_SHIT_SLAVE_MAXCONN"))
	// if err != nil {
	// 	config.MySQLSlaveMaxConn = 100
	// 	return
	// }
	// config.MySQLSlaveConnMaxLifetime, err = strconv.Atoi(os.Getenv("MYSQL_SHIT_SLAVE_CONNMAXLIFETIME"))
	// if err != nil {
	// 	config.MySQLSlaveConnMaxLifetime = 30
	// 	return
	// }
	// if os.Getenv("MYSQL_SHIT_SLAVE_SINGULARTABLE") == "true" {
	// 	config.MySQLSlaveSingularTable = true
	// } else {
	// 	config.MySQLSlaveSingularTable = false
	// }
	// if os.Getenv("MYSQL_SHIT_SLAVE_LOGMODE") == "true" {
	// 	config.MySQLSlaveLogMode = true
	// } else {
	// 	config.MySQLSlaveLogMode = false
	// }

	// //mysql shit master address
	// config.MySQLMasterHost = os.Getenv("MYSQL_SHIT_MASTER_HOST")
	// if config.MySQLMasterHost == "" {
	// 	err = errors.New("env MYSQL_SHIT_MASTER_HOST empty")
	// 	return
	// }
	// config.MySQLMasterPort = os.Getenv("MYSQL_SHIT_MASTER_PORT")
	// if config.MySQLMasterPort == "" {
	// 	err = errors.New("env MYSQL_SHIT_MASTER_PORT empty")
	// 	return
	// }
	// config.MySQLMasterUserName = os.Getenv("MYSQL_SHIT_MASTER_USERNAME")
	// if config.MySQLMasterUserName == "" {
	// 	err = errors.New("env MYSQL_SHIT_MASTER_USERNAME empty")
	// 	return
	// }
	// config.MySQLMasterPassword = os.Getenv("MYSQL_SHIT_MASTER_PASSWORD")
	// if config.MySQLMasterPassword == "" {
	// 	err = errors.New("env MYSQL_SHIT_MASTER_PASSWORD empty")
	// 	return
	// }
	// config.MySQLMasterDBName = os.Getenv("MYSQL_SHIT_MASTER_DBNAME")
	// if config.MySQLMasterDBName == "" {
	// 	err = errors.New("env MYSQL_SHIT_MASTER_DBNAME empty")
	// 	return
	// }
	// config.MySQLMasterDBParam = os.Getenv("MYSQL_SHIT_MASTER_DBPARAM")
	// if config.MySQLMasterDBParam == "" {
	// 	err = errors.New("env MYSQL_SHIT_MASTER_DBPARAM empty")
	// 	return
	// }
	// config.MySQLMasterMaxIdle, err = strconv.Atoi(os.Getenv("MYSQL_SHIT_MASTER_MAXIDLE"))
	// if err != nil {
	// 	config.MySQLMasterMaxIdle = 10
	// 	return
	// }

	// config.MySQLMasterMaxConn, err = strconv.Atoi(os.Getenv("MYSQL_SHIT_MASTER_MAXCONN"))
	// if err != nil {
	// 	config.MySQLMasterMaxConn = 100
	// 	return
	// }
	// config.MySQLMasterConnMaxLifetime, err = strconv.Atoi(os.Getenv("MYSQL_SHIT_MASTER_CONNMAXLIFETIME"))
	// if err != nil {
	// 	config.MySQLMasterConnMaxLifetime = 30
	// 	return
	// }
	// if os.Getenv("MYSQL_SHIT_MASTER_SINGULARTABLE") == "true" {
	// 	config.MySQLMasterSingularTable = true
	// } else {
	// 	config.MySQLMasterSingularTable = false
	// }
	// if os.Getenv("MYSQL_SHIT_MASTER_LOGMODE") == "true" {
	// 	config.MySQLMasterLogMode = true
	// } else {
	// 	config.MySQLMasterLogMode = false
	// }

	// //mysql shit master address
	// config.MySQLMasterHost = os.Getenv("MYSQL_LUCIFER_MASTER_HOST")
	// if config.MySQLMasterHost == "" {
	// 	err = errors.New("env MYSQL_LUCIFER_MASTER_HOST empty")
	// 	return
	// }
	// config.MySQLMasterPort = os.Getenv("MYSQL_LUCIFER_MASTER_PORT")
	// if config.MySQLMasterPort == "" {
	// 	err = errors.New("env MYSQL_LUCIFER_MASTER_PORT empty")
	// 	return
	// }
	// config.MySQLMasterUserName = os.Getenv("MYSQL_LUCIFER_MASTER_USERNAME")
	// if config.MySQLMasterUserName == "" {
	// 	err = errors.New("env MYSQL_LUCIFER_MASTER_USERNAME empty")
	// 	return
	// }
	// config.MySQLMasterPassword = os.Getenv("MYSQL_LUCIFER_MASTER_PASSWORD")
	// if config.MySQLMasterPassword == "" {
	// 	err = errors.New("env MYSQL_LUCIFER_MASTER_PASSWORD empty")
	// 	return
	// }
	// config.MySQLMasterDBName = os.Getenv("MYSQL_LUCIFER_MASTER_DBNAME")
	// if config.MySQLMasterDBName == "" {
	// 	err = errors.New("env MYSQL_LUCIFER_MASTER_DBNAME empty")
	// 	return
	// }
	// config.MySQLMasterDBParam = os.Getenv("MYSQL_LUCIFER_MASTER_DBPARAM")
	// if config.MySQLMasterDBParam == "" {
	// 	err = errors.New("env MYSQL_LUCIFER_MASTER_DBPARAM empty")
	// 	return
	// }
	// config.MySQLMasterMaxIdle, err = strconv.Atoi(os.Getenv("MYSQL_LUCIFER_MASTER_MAXIDLE"))
	// if err != nil {
	// 	config.MySQLMasterMaxIdle = 10
	// 	return
	// }

	// config.MySQLMasterMaxConn, err = strconv.Atoi(os.Getenv("MYSQL_LUCIFER_MASTER_MAXCONN"))
	// if err != nil {
	// 	config.MySQLMasterMaxConn = 100
	// 	return
	// }
	// config.MySQLMasterConnMaxLifetime, err = strconv.Atoi(os.Getenv("MYSQL_LUCIFER_MASTER_CONNMAXLIFETIME"))
	// if err != nil {
	// 	config.MySQLMasterConnMaxLifetime = 30
	// 	return
	// }
	// if os.Getenv("MYSQL_LUCIFER_MASTER_SINGULARTABLE") == "true" {
	// 	config.MySQLMasterSingularTable = true
	// } else {
	// 	config.MySQLMasterSingularTable = false
	// }
	// if os.Getenv("MYSQL_LUCIFER_MASTER_LOGMODE") == "true" {
	// 	config.MySQLMasterLogMode = true
	// } else {
	// 	config.MySQLMasterLogMode = false
	// }

	// //ready line
	// config.HealthPort = os.Getenv("HEALTH_PORT")
	// if config.HealthPort == "" {
	// 	config.HealthPort = "8787"
	// }
	// config.ReadyPath = os.Getenv("READY_PATH")
	// if config.ReadyPath == "" {
	// 	config.ReadyPath = "/ready"
	// }

	//telegram config
	if os.Getenv("TELEGRAM_BOOL") == "true" {
		config.TelegramBool = true
	} else {
		config.TelegramBool = false
	}

	config.TelegramAddress = os.Getenv("TELEGRAM_ADDRESS")
	if config.TelegramAddress == "" {
		logger.Info("env TelegramAddress empty", nil)
		err = errors.New("env TelegramAddress empty")
		return
	}

	config.TelegramChatIDInfo = os.Getenv("TELEGRAM_CHATID")
	if config.TelegramChatIDInfo == "" {
		logger.Info("env TelegramChatIDInfo empty", nil)
		err = errors.New("env TelegramChatIDInfo empty")
		return
	}

	config.TelegramChatIDWarn = os.Getenv("TELEGRAM_CHATID_WARN")
	if config.TelegramChatIDWarn == "" {
		logger.Info("env TelegramChatIDWarn empty", nil)
		err = errors.New("env TelegramChatIDWarn empty")
		return
	}

	config.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	if config.TelegramToken == "" {
		logger.Info("env TelegramToken empty", nil)
		err = errors.New("env TelegramToken empty")
		return
	}

	return nil
}

func main() {
	log.SetOutput(os.Stdout)
	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Compiled = time.Now()
	app.Usage = "telegram bot to receive message"
	app.UsageText = "telegrambot command [command options] [arguments...]"
	app.Commands = []cli.Command{
		cmdStart,
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal("app.run fail", err)
	}

}
