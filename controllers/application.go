package controllers

import (
	"fmt"

	middleware "github.com/go-openapi/runtime/middleware"
	api "github.com/kathra-project/kathra-codegen-helm/restapi/operations/"
	svc "github.com/kathra-project/kathra-codegen-helm/services"
)

func GetTemplates() api.GetTemplatesHandlerFunc {
	return api.GetTemplatesHandler(func(params api.GetTemplatesParams) middleware.Responder {
		templates, err := svc.GetTemplates()
		if err != nil {
			fmt.Println(err)
			//return api.NewGetTemplatesServerError().WithPayload("Get code generates internal error")
			return api.NewGetTemplatesOK().WithPayload(nil)
		} else {
			return api.NewGetTemplatesOK().WithPayload(templates)
		}
	})
}
