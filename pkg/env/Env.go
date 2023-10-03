package env

import "os"

type Env struct {
	InfluxToken string
	InfluxUrl   string
	InfluxOrg   string
}

func (e *Env) String() string {
	return "InfluxToken: " + e.InfluxToken + "\n" +
		"InfluxUrl: " + e.InfluxUrl + "\n" +
		"InfluxOrg: " + e.InfluxOrg + "\n"
}

func Read() Env {
	env := Env{
		InfluxToken: os.Getenv("INFLUX_TOKEN"),
		InfluxUrl:   os.Getenv("INFLUX_URL"),
		InfluxOrg:   os.Getenv("INFLUX_ORG"),
	}
	return env
}
