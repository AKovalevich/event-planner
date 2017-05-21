package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/AKovalevich/event-planner/app"
)

//
func HashMd5(string string) string {
	var hash = md5.New()
	string = string + app.Config().Secret
	hash.Write([]byte(string))
	return hex.EncodeToString(hash.Sum(nil))
}
