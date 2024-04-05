package kadi

import (
	"github.com/google/uuid"
)

type app struct {
	appID string
}

var kadigo app

func KadiGo(appID string) app {
	if appID != "" {
		appID = uuid.New().String()
	}
	InitTokenAuth()
	kadigo = app{appID: appID}
	return kadigo
}
