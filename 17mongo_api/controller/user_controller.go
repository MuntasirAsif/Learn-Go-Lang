package controller
import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muntasirashif/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// Insert one student
func insertOneStudent(student model.Student) {
	result, err := studentCollection.InsertOne(context.Background(), student)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted student with ID:", result.InsertedID)
}

// Get all students
func getAllStudents() []primitive.M {
	cursor, err := studentCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var students []primitive.M

	for cursor.Next(context.Background()) {
		var student bson.M
		if err := cursor.Decode(&student); err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	defer cursor.Close(context.Background())
	return students
}

// Get one student by ID
func getOneStudent(studentId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}

	var student primitive.M
	err := studentCollection.FindOne(context.Background(), filter).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}

	return student
}

// Update student by ID
func updateOneStudent(studentId string, student model.Student) {
	id, _ := primitive.ObjectIDFromHex(studentId)

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": student,
	}

	result, err := studentCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Matched %v document and updated %v document\n",
		result.MatchedCount, result.ModifiedCount)
}

// Delete one student
func deleteOneStudent(studentId string) {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}

	result, err := studentCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v student\n", result.DeletedCount)
}

// Delete all students
func deleteAllStudents() int64 {
	result, err := studentCollection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v students\n", result.DeletedCount)
	return result.DeletedCount
}

////////////////////////
// HTTP Controllers  //
////////////////////////

// POST /students
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student model.Student
	_ = json.NewDecoder(r.Body).Decode(&student)

	insertOneStudent(student)
	json.NewEncoder(w).Encode(student)
}

// GET /students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	students := getAllStudents()
	json.NewEncoder(w).Encode(students)
}

// GET /students/{id}
func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	student := getOneStudent(params["id"])

	json.NewEncoder(w).Encode(student)
}

// PUT /students/{id}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var student model.Student
	_ = json.NewDecoder(r.Body).Decode(&student)

	updateOneStudent(params["id"], student)
	json.NewEncoder(w).Encode(params["id"])
}

// DELETE /students/{id}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	deleteOneStudent(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

// DELETE /students
func DeleteAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	count := deleteAllStudents()
	json.NewEncoder(w).Encode(count)
}
