package cuser

import (
	"ginp-api/internal/app/gapi/entity"
)

type RespondLogin struct {
	Token    string          `json:"token"`
	UserInfo *entity.User `json:"user_info"`
}


