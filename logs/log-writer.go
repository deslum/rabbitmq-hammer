package logs

import (
	"fmt"
)

type LogWriter struct{}


func (o LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}
