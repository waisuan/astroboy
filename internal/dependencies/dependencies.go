package dependencies

type Dependencies struct {
	Config   *Config
	CacheCli ICache
	SqsCli   *SqsCli
	KafkaCli *KafkaCli
	Db       *DB
}

func Init() *Dependencies {
	cfg := LoadEnv()

	sqsCli := NewSqsCli(cfg)

	kafkaCli := NewKafkaCli(cfg)

	cacheCli := NewCache(cfg)

	db := InitDB(cfg)

	return &Dependencies{
		Config:   cfg,
		SqsCli:   sqsCli,
		KafkaCli: kafkaCli,
		CacheCli: cacheCli,
		Db:       db,
	}
}
