package service

import (
	"context"
	"encoding/json"
	"fmt"
	"goToDo/server/models"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	log "github.com/labstack/gommon/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {
	loadTheEnv()
	createDBInstance()
}

// load .env file
func loadTheEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func createDBInstance() {
	// DB connection string
	connectionString := os.Getenv("DB_URI")

	// Database Name
	dbName := os.Getenv("DB_NAME")

	// Collection name
	collName := os.Getenv("DB_COLLECTION_NAME")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

// GetAllTask get all the task route
func GetAllTask(c echo.Context) error {
	log.Info("GetAllTask")
	r := c.Request().Header
	w := c.Response()
	r.Set("Context-Type", "application/x-www-form-urlencoded")
	r.Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	return c.JSON(http.StatusOK, json.NewEncoder(w).Encode(payload))
}

// CreateTask create task route
func CreateTask(c echo.Context) error {
	r := c.Request()
	w := c.Response()
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	log.Printf("Request Body: %v", r.Body)
	insertOneTask(task)
	return c.JSON(http.StatusOK, json.NewEncoder(w).Encode(task))
}

// TaskComplete update task route
func TaskComplete(c echo.Context) error {
	w := c.Response()
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	param := c.Param("id")
	taskComplete(param)
	return c.JSON(http.StatusOK, json.NewEncoder(w).Encode(param))
}

// UndoTask undo the complete task route
func UndoTask(c echo.Context) error {
	w := c.Response()
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	param := c.Param("id")
	undoTask(param)
	return c.JSON(http.StatusOK, json.NewEncoder(w).Encode(param))
}

// DeleteTask delete one task route
func DeleteTask(c echo.Context) error {
	w := c.Response()
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// json.NewEncoder(w).Encode("Task not found")
	param := c.Param("id")
	deleteOneTask(param)
	return c.JSON(http.StatusOK, json.NewEncoder(w).Encode(param))
}

// DeleteAllTask delete all tasks route
func DeleteAllTask(c echo.Context) error {
	w := c.Response()
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTask()
	return c.JSON(http.StatusOK, json.NewEncoder(w).Encode(count))
	// json.NewEncoder(w).Encode("Task not found")

}

// get all task from the DB and return it
func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Inserted a Single Record ", insertResult.InsertedID)
}

// task complete method, update task's status to true
func taskComplete(task string) {
	log.Info(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("modified count: ", result.ModifiedCount)
}

// task undo method, update task's status to false
func undoTask(task string) {
	log.Info(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("modified count: ", result.ModifiedCount)
}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	log.Info(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Deleted Document", d.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}
