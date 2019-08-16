package modules

import "strconv"

// Atouint converts string to uint
func Atouint(s string) (uint, error) {
	ui64, err := strconv.ParseUint(s, 10, 64)
	return uint(ui64), err
}
