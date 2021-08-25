package adress_book

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContactDB struct {
	db *mongo.Database
}
type Contact struct {
	Name      string `db:"name"`
	Phone     string `db:"phone"`
	GroupName string `db:"groupName"`
}

func (u *ContactDB) AddContact(name string, phone string, groupName string) (string, error) {

	obj, err := u.db.Collection("contacts").InsertOne(context.Background(), bson.M{
		"name":      name,
		"phone":     phone,
		"groupName": groupName,
	})
	if err != nil {
		return "", err
	}
	fmt.Printf("inserted document with ID %v\n", obj.InsertedID)
	return obj.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *ContactDB) AssignToGroup(id string, groupName string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	opts := options.Update()
	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$set", bson.D{{"groupName", groupName}}}}

	_, err = u.db.Collection("contacts").UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}
	fmt.Println("matched and replaced an existing document")
	return nil
}

func (u *ContactDB) ListAllByGroup(groupName string) ([]Contact, error) {
	c, err := u.db.Collection("contacts").Find(context.Background(), bson.M{"groupName": groupName})
	if err != nil {
		return nil, err
	}
	var contacts []bson.M
	if err = c.All(context.Background(), &contacts); err != nil {
		return nil, err
	}
	res := make([]Contact, 0, len(contacts))
	for _, v := range contacts {
		con := Contact{
			Name:      v["name"].(string),
			Phone:     v["phone"].(string),
			GroupName: v["groupName"].(string),
		}
		res = append(res, con)
	}
	return res, nil
}

func ConnToDB(uri string) (*ContactDB, error) {

	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = c.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &ContactDB{
		db: c.Database("adress_book"),
	}, nil

}
