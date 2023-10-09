package data

import (
	"fmt"
)

type EquationCaptor struct {
	Name    string
	Captors []CaptorValue
}

func (equationCaptor *EquationCaptor) String() string {
	return fmt.Sprintf("%s with %d captors", equationCaptor.Name, len(equationCaptor.Captors))
}

func (equationCaptor *EquationCaptor) GetCaptor(i int) *CaptorValue {
	return &equationCaptor.Captors[i]
}
