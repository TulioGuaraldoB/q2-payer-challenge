package swagger

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/docs"
	_ "github.com/TulioGuaraldoB/q2-payer-challenge/docs"
)

func SetupSwagger() {
	docs.SwaggerInfo.Version = "v1.0"
	docs.SwaggerInfo.Title = "q2-payer-challenge"
	docs.SwaggerInfo.Description = "REST Api for Mid-Level Back-End challenge"
	docs.SwaggerInfo.BasePath = "api/v1/"
}
