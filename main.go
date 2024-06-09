package main

import (
	_ "github.com/lib/pq"
	"github.com/wonyus/update-status-service/contexts"
	"github.com/wonyus/update-status-service/controllers"
	"github.com/wonyus/update-status-service/pkg/client"
	"github.com/wonyus/update-status-service/pkg/db"
)

func main() {
	var messageReceived = make(chan struct{})

	ctx := contexts.NewResource()
	db := db.NewPGDB(ctx)
	repo := controllers.NewRepository(ctx, db)
	mqttHandler := client.NewMqttHandler(ctx, repo, messageReceived)

	mqttHandler.InitialMqttClient(mqttHandler.MessageSubHandler)

	for {
		<-messageReceived
	}
}

// queryAPI := client.QueryAPI(org)
// query := `from(bucket: "Switch")
//             |> range(start: -10m)
//             |> filter(fn: (r) => r._measurement == "measurement1")`
// results, err := queryAPI.Query(context.Background(), query)
// if err != nil {
//     log.Fatal(err)
// }
// for results.Next() {
//     fmt.Println(results.Record())
// }
// if err := results.Err(); err != nil {
//     log.Fatal(err)
// }
// query = `from(bucket: "Switch")
//               |> range(start: -10m)
//               |> filter(fn: (r) => r._measurement == "measurement1")
//               |> mean()`
// results, err = queryAPI.Query(context.Background(), query)
// if err != nil {
//     log.Fatal(err)
// }
// for results.Next() {
//     fmt.Println(results.Record())
// }
// if err := results.Err(); err != nil {
//     log.Fatal(err)
// }
