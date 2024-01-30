package main

import (
	"context"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/wonyus/update-status-service/controllers"
	"github.com/wonyus/update-status-service/initials"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var ginLambda *chiadapter.ChiLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	initials.InitDB()
	r := initials.InitHttpServer()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/update-status", controllers.UpdateStatus)

	ginLambda = chiadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
