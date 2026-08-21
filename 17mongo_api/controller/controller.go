package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/muntasirashif/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://muntasirashifee_db_user:ICSfLSgGaS5I1CYS@cluster0.5k5nuph.mongodb.net/?appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"
const studentCol = "students"

// Most Important consts for MongoDB connection
var collection *mongo.Collection
var studentCollection *mongo.Collection


func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(connectionString + "&tls=true")
		//SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	// clientOptions.SetTLSConfig(&tls.Config{
	// 	MinVersion: tls.VersionTLS12,
	// })
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongo connect error:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping error:", err)
	}
	db :=client.Database(dbName)
	collection = db.Collection(colName)
	studentCollection = db.Collection(studentCol)
	log.Println("MongoDB connected successfully")
}





// MongoDB helper function
// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, error := collection.InsertOne(context.Background(), movie)

	if error != nil {
		log.Fatal(error)
	}
	log.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := primitive.M{"_id": id}
	update := primitive.M{"$set": bson.M{"watched": true}}

	result, error := collection.UpdateOne(context.Background(), filter, update)
	if error != nil {
		log.Fatal(error)
	}

	log.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := primitive.M{"_id": id}
	result, error := collection.DeleteOne(context.Background(), filter)
	if error != nil {
		log.Fatal(error)
	}

	log.Printf("Deleted %v documents in the movies collection\n", result.DeletedCount)
}

// delete all records
func deleteAllMovies() int64 {
	result, error := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if error != nil {
		log.Fatal(error)
	}

	log.Printf("Deleted %v documents in the movies collection\n", result.DeletedCount)
	return result.DeletedCount
}

// get all movies from db
func getAllMovies() []primitive.M {
	cur, error := collection.Find(context.Background(), bson.D{{}})
	if error != nil {
		log.Fatal(error)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		error := cur.Decode(&movie)
		if error != nil {
			log.Fatal(error)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

/// acctual controller functions called in main.go will be here ///

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allmovies := getAllMovies()
	log.Println("All movies: ", allmovies)

	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
