package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/wonyus/update-status-service/controllers"
	"github.com/wonyus/update-status-service/initials"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var ginLambda *chiadapter.ChiLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	initials.InitDB()
	r := initials.InitHttpServer()
	client := initials.InitialMqttClient(controllers.MessagePubHandler)
	controllers.DefaultSubscribeHandler(client)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("healty")
		w.Write([]byte("healty"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("ping")
		w.Write([]byte("pong"))
	})
	r.Post("/update-status", controllers.UpdateStatus)

}

func main() {
	port := os.Getenv("PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
