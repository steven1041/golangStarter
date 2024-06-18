package controller

import "golangStarter/models"

type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`
	Message string                  `json:"message"`
	Data    []*models.ApiPostDetail `json:"data"`
}
