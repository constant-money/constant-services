package server

import (
	"fmt"

	"github.com/constant-money/constant-services/fcm-service/config"
)

// Init :
func Init() {
	conf := config.GetConfig()

	r := NewRouter()
	r.Run(fmt.Sprintf(":%d", conf.Port))
}
