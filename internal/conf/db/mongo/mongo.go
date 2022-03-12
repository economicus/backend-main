package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const maxIter = 20

type DB struct {
	conf
	DB *mongo.Database
}

func New() *DB {
	md := DB{}
	md.conf = newConf()
	md.GetMongoDB()
	md.testDBConnection()
	md.Info()
	return &md
}

func (md *DB) GetMongoDB() {
	var err error
	dsn := md.GetDSN()
	log.Println(dsn)

	for i := 0; i < maxIter; i++ {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
		if err == nil {
			md.DB = client.Database(md.name)
			return
		} else {
			log.Printf("MongoDB connection has failed. Sleeping for a second...: %v", err)
			time.Sleep(1 * time.Second)
		}
	}
	log.Panicf("error while connecting mongoDB with dsn '%s': %v", dsn, err)
}

func (md *DB) testDBConnection() {
	for i := 0; i < maxIter; i++ {
		if err := md.DB.Client().Ping(context.TODO(), readpref.Primary()); err == nil {
			log.Println("[INFO] Pinged mongoDB successfully")
			return
		} else {
			log.Printf("[INFO] MongoDB has not been prepared yet. Sleeping for a second...: %v", err)
			time.Sleep(1 * time.Second)
		}
	}
	log.Panicf("error while testing connection: tried %d times but failed", maxIter)
}
