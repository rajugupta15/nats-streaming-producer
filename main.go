package main

// Import packages
import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	nats "github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
	createjson "github.com/rajugupta15/nats-streaming-producer/createJson"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	clusterID := "test-cluster"
	clientID := "nats-streaming-producer"
	metricsURL := "localhost:8080"
	natsURL := "localhost:4222"
	nc, err := nats.Connect(
		natsURL,
		nats.MaxReconnects(30),
		nats.ReconnectWait(1*time.Second),
		nats.DisconnectHandler(func(nc *nats.Conn) {
			fmt.Printf("[WARN] Client[%s] disconnected.\n", clientID)
			_, err := http.Get(metricsURL + "/prodisconncount")
			if err != nil {
				log.Printf("[ERROR] %s", err)
			}
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Printf("[WARN] Client[%s] reconnected to %v.\n", clientID, nc.ConnectedUrl())
			_, err := http.Get(metricsURL + "/proreconncount")
			if err != nil {
				log.Printf("[ERROR] %s", err)
			}
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			fmt.Printf("[WARN] Client[%s] connection to %v closed: %q\n", clientID, nc.ConnectedUrl(), nc.LastError())
			_, err := http.Get(metricsURL + "/proconnclosscount")
			if err != nil {
				log.Printf("[ERROR] %s", err)
			}
		}),
	)
	natsConnection, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, natsURL)
	}
	defer natsConnection.Close()
	subject := "foo"
	// i := 0
	// for i = 0; i < 1000; i++ {
	for {
		json := createjson.CreateJson(time.Now().UnixNano(), randStringBytes(15))
		err := natsConnection.Publish(subject, []byte(json))
		if err != nil {
			log.Printf("[ERROR] Publishing message '%s': %s", json, err)
		} else {
			log.Println("Published message on subject" + subject)
			_, err := http.Get(metricsURL + "/promesgcount")
			if err != nil {
				log.Printf("[ERROR] %s", err)
			}
		}
	}
	// time.Sleep(3600 * time.Second)
}
