package config

type Config struct {
	Address         string
	PingCount       int
	PingInterval    int
	RestartInterval int
	FailTimes       int
}
