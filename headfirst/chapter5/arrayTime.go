package main

import (
	"fmt"
	"time"
)

func main() {
	var time_list [3]time.Time
	time_list[0] = time.Unix(13253864390, 0)
	time_list[1] = time.Unix(13257864390, 0)
	time_list[2] = time.Unix(18953864390, 0)

	fmt.Println(time_list[0], time_list[1], time_list[2])
}