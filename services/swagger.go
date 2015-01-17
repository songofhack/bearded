package services

import (
	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/sirupsen/logrus"
)

func init() {
	swagger.LogInfo = logrus.Infof
}

func Swagger(wsContainer *restful.Container, apiPath, swaggerPath, swaggerFilepath string) {
	config := swagger.Config{
		WebServices: wsContainer.RegisteredWebServices(), // you control what services are visible
		ApiPath:     apiPath,

		// Optionally, specifiy where the UI is located
		SwaggerPath:     swaggerPath,
		SwaggerFilePath: swaggerFilepath,
	}
	swagger.RegisterSwaggerService(config, wsContainer)
}