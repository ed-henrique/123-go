package utils

import "flag"

func GetFlags() int64 {
	durationFlag := flag.Int64("time", 0, "Starting time to countdown from in milliseconds")
	flag.Parse()

	return *durationFlag
}
