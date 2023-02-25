package main

import "os"

func GetenvOrElse(envVar string, fallback string) string {
	value := os.Getenv(envVar)
	if value == "" {
		value = fallback
	}
	return value
}
