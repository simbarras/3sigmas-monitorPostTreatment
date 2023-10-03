package data

import (
	"fmt"
	"time"
)

type ComputedMeasure struct {
	DateTime  time.Time
	Value     float64
	Captor    string
	Variables []string
}

func (m ComputedMeasure) String() string {
	return m.DateTime.Format("2006-01-02T15:04:05Z") + " " + m.Captor + " " + fmt.Sprintf("%f", m.Value)
}

func (m ComputedMeasure) Tags() map[string]string {
	tag := ""
	for _, variable := range m.Variables {
		tag += variable + ";"
	}
	return map[string]string{
		"captors": tag,
	}
}

func (m ComputedMeasure) Fields() map[string]interface{} {
	return map[string]interface{}{
		"value": m.Value,
	}
}

func (m ComputedMeasure) Measurement() string {
	return m.Captor
}

func (m ComputedMeasure) Date() time.Time {
	return m.DateTime
}
