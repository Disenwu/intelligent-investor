package server

import (
	"fmt"
	options "intelligent-investor/cmd/api-server/options"
	"intelligent-investor/internal/app/middleware"
	"intelligent-investor/internal/app/router/v1"
	"intelligent-investor/internal/pkg/log"
	"intelligent-investor/internal/pkg/service"
	"intelligent-investor/pkg/version"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var configFile string

type Config struct {
	JWTSecret      string                 `json:"jwtSecret"`
	Expiration     time.Duration          `json:"expiration"`
	Port           int                    `json:"port"`
	DatabaseConfig service.DatabaseConfig `json:"databaseConfig"`
	RedisConfig    service.RedisConfig    `json:"redisConfig"`
}

func NewServerCommand() *cobra.Command {
	opts := options.NewServerOptions()
	cmd := &cobra.Command{
		Use:           "intelligent-investor",
		Short:         "a web tool for invest to analize financial report",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		Args: cobra.NoArgs,
	}

	cobra.OnInitialize(onInitialize)
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/api-server.yaml)")
	opts.AddFlags(cmd.PersistentFlags())
	version.AddFlags(cmd.PersistentFlags())
	return cmd
}

func run(otps *options.ServerOptions) error {
	version.PrintAndExitIfRequested()

	log.Init()
	defer log.Sync()

	if err := viper.Unmarshal(otps); err != nil {
		return err
	}

	if err := otps.Validate(); err != nil {
		return err
	}

	cfg := FromServerOptions(otps)

	server, err := cfg.NewUnionServer()
	if err != nil {
		return err
	}

	return server.Run()
}

func FromServerOptions(otps *options.ServerOptions) *Config {
	// 从 Viper 获取数据库配置
	dbConfig := &service.DatabaseConfig{}
	if err := viper.UnmarshalKey("database", dbConfig); err != nil {
		log.Panicw("Failed to unmarshal database config", "error", err)
	}
	// 从 Viper 获取 Redis 配置
	redisConfig := &service.RedisConfig{}
	if err := viper.UnmarshalKey("redis", redisConfig); err != nil {
		log.Panicw("Failed to unmarshal redis config", "error", err)
	}
	authIgnorePaths := viper.GetStringSlice("auth.ignore-paths")
	middleware.IgnorePathsInit(authIgnorePaths)
	return &Config{
		JWTSecret:      otps.JWTSecret,
		Expiration:     otps.Expiration,
		Port:           otps.Port,
		DatabaseConfig: *dbConfig,
		RedisConfig:    *redisConfig,
	}
}

type UnionServer struct {
	cfg *Config
}

func (cfg *Config) NewUnionServer() (*UnionServer, error) {
	return &UnionServer{
		cfg: cfg,
	}, nil
}

func (s *UnionServer) Run() error {
	service.DatabaseInitialize(&s.cfg.DatabaseConfig)
	service.RedisInitialize(&s.cfg.RedisConfig)
	r := gin.Default()
	r.Use(gin.Recovery(), middleware.CORSMiddleware(), middleware.RequestIDMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.UserRouter(r)
	r.Run(fmt.Sprintf(":%d", s.cfg.Port))
	return nil
}
