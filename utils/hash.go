package utils

import (
	"github.com/AKovalevich/event-planner/app"
	"crypto/md5"
	"encoding/hex"
)

//
func HashMd5(string string) string {
	var hash = md5.New()
	string = string + app.Config().Secret
	hash.Write([]byte(string))
	return hex.EncodeToString(hash.Sum(nil))
}
