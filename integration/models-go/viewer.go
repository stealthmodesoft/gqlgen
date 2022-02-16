package models

import "github.com/stealthmodesoft/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
