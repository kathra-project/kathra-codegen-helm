package controllers

import (
	"fmt"
	"os"

	middleware "github.com/go-openapi/runtime/middleware"
	api "github.com/kathra-project/kathra-codegen-helm/restapi/operations"
	svc "github.com/kathra-project/kathra-codegen-helm/services"
)

func GetTemplates() api.GetTemplatesHandler {
	return api.GetTemplatesHandlerFunc(func(params api.GetTemplatesParams) middleware.Responder {

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

func GenerateFromTemplate() api.GenerateFromTemplateHandler {
	return api.GenerateFromTemplateHandlerFunc(func(params api.GenerateFromTemplateParams) middleware.Responder {
		filePath, err := svc.GenerateFilesFromTemplate(params.CodeGenTemplate)
		if err != nil {
			fmt.Println(err)
			return api.NewGenerateFromTemplateInternalServerError().WithPayload("Error generation")
		} else {
			var file, fileErr = os.Open(filePath)
			if fileErr != nil {
				if os.IsNotExist(fileErr) {
					return api.NewGenerateFromTemplateInternalServerError().WithPayload("Error generation")
				} else if os.IsPermission(fileErr) {
					return api.NewGenerateFromTemplateInternalServerError().WithPayload("Error generation")
				}

				return api.NewGenerateFromTemplateInternalServerError().WithPayload("Error generation")
			}
			return api.NewGenerateFromTemplateOK().WithPayload(file)
		}
	})
}
