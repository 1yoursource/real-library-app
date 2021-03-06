package auto_incrementer

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session      = &mgo.Session{}
	collection   = &mgo.Collection{}
	idFieldName  = "id"
	seqFieldName = "seq"
)

type (
	AutoIncrementer struct {
		session      *mgo.Session
		collection   *mgo.Collection
		idFieldName  string
		seqFieldName string
	}
)

// Connect to database
func Connect(c *mgo.Collection, fieldNames ...string) {
	collection = c
	session = c.Database.Session
	if len(fieldNames) > 0 {
		idFieldName = fieldNames[0]
	}
	if len(fieldNames) > 1 {
		seqFieldName = fieldNames[1]
	}
}

// Next sequence of AutoIncrement
func Next(name string) uint64 {
	connectionCheck()
	result := bson.M{}
	if _, err := collection.Find(bson.M{idFieldName: name}).Apply(mgo.Change{
		Update:    bson.M{"$set": bson.M{idFieldName: name}, "$inc": bson.M{seqFieldName: 1}},
		Upsert:    true,
		ReturnNew: true,
	}, &result); err != nil {
		fmt.Println("Autoincrement error(1):", err.Error())
	}

	sec, ok := result[seqFieldName].(int)
	if ok {
		return uint64(sec)
	}

	sec2, _ := result[seqFieldName].(int64)
	return uint64(sec2)
}

// Cancel is decrement counter value
func Cancel(name string) {
	connectionCheck()
	if err := collection.Update(bson.M{"id": name}, bson.M{"$inc": bson.M{"seq": -1}}); err != nil {
		fmt.Println("Autoincrement error(2):", err.Error())
	}
}

func connectionCheck() {
	if err := session.Ping(); err != nil {
		session.Refresh()
	}
}

// Create new instance of AI
func create(c *mgo.Collection, fieldNames ...string) *AutoIncrementer {
	ai := &AutoIncrementer{
		collection:   c,
		session:      c.Database.Session,
		idFieldName:  "id",
		seqFieldName: "seq",
	}

	if len(fieldNames) > 0 {
		ai.idFieldName = fieldNames[0]
	}
	if len(fieldNames) > 1 {
		ai.seqFieldName = fieldNames[1]
	}
	return ai
}

func (ai *AutoIncrementer) Next(name string) uint64 {
	ai.connectionCheck()
	result := bson.M{}
	if _, err := ai.collection.Find(bson.M{ai.idFieldName: name}).Apply(mgo.Change{
		Update:    bson.M{"$set": bson.M{ai.idFieldName: name}, "$inc": bson.M{ai.seqFieldName: 1}},
		Upsert:    true,
		ReturnNew: true,
	}, &result); err != nil {
		fmt.Println("Autoincrement error(1):", err.Error())
	}

	sec, ok := result[ai.seqFieldName].(int)
	if ok {
		return uint64(sec)
	}

	sec2, _ := result[ai.seqFieldName].(int64)
	return uint64(sec2)
}

// Cancel is decrement counter value
func (ai *AutoIncrementer) Cancel(name string) {
	ai.connectionCheck()
	if err := ai.collection.Update(bson.M{ai.idFieldName: name}, bson.M{"$inc": bson.M{ai.seqFieldName: -1}}); err != nil {
		fmt.Println("Autoincrement error(2):", err.Error())
	}
}

func (ai *AutoIncrementer) connectionCheck() {
	if err := ai.session.Ping(); err != nil {
		ai.session.Refresh()
	}
}
