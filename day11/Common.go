package main

import (
	"fmt"
	"strconv"
)


func ReduceZeros(s string) string {
	v, err := strconv.Atoi(s)
	if err == nil {
		return fmt.Sprintf("%d", v)
	}
	return s
}