package cache

import (
	"github.com/shubhamsnehi/golang-api-testing/entity"
)

type PostCache interface {
	Set(key string, value *entity.User)
	Get(key string) *entity.User
}
