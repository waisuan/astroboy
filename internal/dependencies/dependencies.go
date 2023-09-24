package dependencies

type Dependencies struct {
	Config   *Config
	CacheCli *Cache
	SqsCli   *SqsCli
	KafkaCli *KafkaCli
}

func Init() *Dependencies {
	cfg := LoadEnv()

	sqsCli := NewSqsCli(cfg)

	kafkaCli := NewKafkaCli(cfg)

	cacheCli := NewCache(cfg)

	return &Dependencies{
		Config:   cfg,
		SqsCli:   sqsCli,
		KafkaCli: kafkaCli,
		CacheCli: cacheCli,
	}
}
