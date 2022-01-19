// Package log is contains utilities for logging
package log

import "log"

func LogErr(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}
