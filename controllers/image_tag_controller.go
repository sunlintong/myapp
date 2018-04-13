package controllers

import (

)

type ImageTagController struct {
	BaseController
}

type ImageTagRequest struct {
	Image_Name string `json:"image_name"`
	EventType string `json:"event_type"`
}

const (
	DeleteTag = "deteletag"
	AddTag = "addtag"
)