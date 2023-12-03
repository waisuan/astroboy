package dependencies

import "context"

var ctx = context.Background()

type Dependencies struct {
	Config   *Config
	CacheCli ICache
	SqsCli   *SqsCli
	KafkaCli *KafkaCli
	Db       IDatabase
}

func Init() *Dependencies {
	cfg := LoadEnv()

	//sqsCli := NewSqsCli(cfg)
	//
	//kafkaCli := NewKafkaCli(cfg)
	//
	//cacheCli := NewCache(cfg)

	db := InitDB(cfg)

	return &Dependencies{
		Config: cfg,
		//SqsCli:   sqsCli,
		//KafkaCli: kafkaCli,
		//CacheCli: cacheCli,
		Db: db,
	}
}
