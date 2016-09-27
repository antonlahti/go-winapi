package winapi

import (
	"strconv"
)

func Errstr(code int) string {
	return "Error code:" + strconv.Itoa(code)
}
