package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("error of data parsing '%s': %v", data, err)
			continue
		}
		_, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error of data: %v", err)
			continue
		}
	}
}
