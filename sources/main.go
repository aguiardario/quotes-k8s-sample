//Package main se encarga de retornar quotes de autores famosos.
// Si no hay frases cargadas, popula la base de datos con 6 quotes.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quotes = [6]string{
	"Nunca emprenderíamos nada si quisiéramos asegurar por anticipado el éxito de nuestra empresa (Napoleón Bonaparte)",
	"El verdadero emprendedor no es un soñador, es un hacedor (Nolan Bushnell)",
	"Si no puedes hacer grandes cosas, haz pequeñas cosas a lo grande (Napoleon Hill)",
	"Cuando soplan vientos de cambio, unos buscan refugio y se ponen a salvo y otros construyen molinos y se hacen ricos (Claus Möller)",
	"Un hombre con una nueva idea es un loco hasta que ésta triunfa (Mark Twain)",
	"Un líder es alguien a quien sigues a un lugar al que no irías por ti mismo (Joel Arthur Barker)",
}

//Quote es la estructura que se serializa y se retorna en formato json.
type Quote struct {
	ID    primitive.ObjectID `bson:"_id, omitempty"`
	Quote string             `json:"quote"`
}

func getConnectionString() string {

	port := "27017"
	host := "localhost"
	user := ""
	pwd := ""
	dbName := "quotes"

	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if os.Getenv("USER_DB") != "" {
		user = os.Getenv("USER_DB")
	}

	if os.Getenv("PWD_DB") != "" {
		pwd = os.Getenv("PWD_DB")
	}

	if user != "" {
		user = user + ":" + pwd + "@"
	}

	// fmt.Println("mongodb://%s:%s@%s:%s", user, pwd, host, port)
	fmt.Println("mongodb://" + user + "@" + host + ":" + port + "/" + dbName)
	return fmt.Sprintf("mongodb://" + user + host + ":" + port + "/" + dbName)
}

func home(w http.ResponseWriter, r *http.Request) {

	uriMongo := getConnectionString()

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	clientOpts := options.Client().ApplyURI(uriMongo)
	client, err := mongo.Connect(nil, clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Congratulations, you're already connected to MongoDB!")

	collection := client.Database("quotes").Collection("quotes")

	rand.Seed(time.Now().UnixNano())

	count, _ := collection.EstimatedDocumentCount(context.TODO())

	fmt.Println(count)
	magic := rand.Intn(int(count))

	cur, err := collection.Find(context.TODO(), bson.D{})
	defer cur.Close(context.TODO())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error")
	}

	var results []Quote
	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Quote
		err := cur.Decode(&elem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error")
		}

		results = append(results, elem)
	}

	quote := results[magic]

	// res := &Quote{
	// 	Quote: quote.Quote,
	// }

	jsonResult, err := json.Marshal(quote)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error")
		return
	}

	log.Println(string(jsonResult))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)

}

func connect() {

	uriMongo := getConnectionString()

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	clientOpts := options.Client().ApplyURI(uriMongo)

	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Congratulations, you're already connected to MongoDB!")

	collection := client.Database("quotes").Collection("quotes")

	count, _ := collection.EstimatedDocumentCount(context.TODO())

	if count == 0 {

		for _, q := range quotes {

			quote := Quote{
				Quote: q,
			}

			insertResult, err := collection.InsertOne(context.TODO(), quote)

			if err != nil {
				log.Fatal(err)
			}

			log.Println("Death Star had been inserted: ", insertResult.InsertedID)
		}

	}

	defer client.Disconnect(context.TODO())

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func main() {

	connect()
	http.HandleFunc("/", home)
	http.HandleFunc("/check", HealthCheckHandler)
	// p := properties.MustLoadFile("config.properties", properties.UTF8)
	// port := ":" + p.GetString("port", "3000")
	port := ":3000"
	log.Print("=========================> Servidor escuchando en puerto:", port)
	http.ListenAndServe(port, nil)
}
