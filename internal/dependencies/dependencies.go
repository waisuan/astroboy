package dependencies

import "context"

var ctx = context.Background()

type Dependencies struct {
	Config   *Config
	CacheCli ICache
	DB       IDatabase
}

func Init() *Dependencies {
	cfg := LoadEnv()

	cacheCli := NewCache(cfg)

	db := InitDB(cfg)

	return &Dependencies{
		Config: cfg,
		//SqsCli:   sqsCli,
		//KafkaCli: kafkaCli,
		CacheCli: cacheCli,
		DB:       db,
	}
}
