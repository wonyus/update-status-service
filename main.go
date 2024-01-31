package main

import (
	"context"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/wonyus/update-status-service/controllers"
	"github.com/wonyus/update-status-service/initials"
	"github.com/wonyus/update-status-service/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var ginLambda *chiadapter.ChiLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	initials.InitDB()
	r := initials.InitHttpServer()
	client := initials.InitialMqttClient(controllers.MessagePubHandler)
	controllers.DefaultSubscribeHandler(client)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Post("/update-status", controllers.UpdateStatus)

	// Check the MODE environment variable to decide whether to use local server or AWS Lambda
	if utils.Strip(os.Getenv("MODE")) != "dev" {
		ginLambda = chiadapter.New(r)
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// If running in local mode, start the HTTP server
	if utils.Strip(os.Getenv("MODE")) == "dev" {
		http.ListenAndServe(":8080", nil)
	} else {
		// Otherwise, start the Lambda function
		lambda.Start(Handler)
	}
}
