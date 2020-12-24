package service

import (
	"server/utils"
)

func StartHeaderService(config *utils.MyConfig) {
	StartHeader(config)
}

func StartFollowerService(config *utils.MyConfig) {
	StartFollower(config)
}
