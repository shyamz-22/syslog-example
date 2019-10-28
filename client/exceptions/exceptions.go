package exceptions

import "log"

func LogAndIgnore(log *log.Logger, err error) {
	if  err != nil {
		log.Printf("Error: %v", err)
	}
	return
}
