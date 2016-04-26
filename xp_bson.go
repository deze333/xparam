package xparam

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//------------------------------------------------------------
// Xparam access for bson.ObjectId
//------------------------------------------------------------

// Gets key value as if it should be bson.ObjectId.
func (xp XP) As_ObjectId(key string) (id *bson.ObjectId) {

	if val, ok := xp[key]; ok {

		// Is it already ObjectId ?
		if oid, ok := val.(bson.ObjectId); ok {
			id = &oid
			return
		}

		// Treat as string
		valstr := fmt.Sprint(val)
		if bson.IsObjectIdHex(valstr) {
			bsonId := bson.ObjectIdHex(valstr)
			id = &bsonId
		}
	}
	return
}

// Gets key value as bson.ObjectId.
func (xp XP) MustBe_ObjectId(key string) (id bson.ObjectId, err error) {

	if val, ok := xp[key]; ok {
		valstr := fmt.Sprint(val)
		if bson.IsObjectIdHex(valstr) {
			id = bson.ObjectIdHex(valstr)
		} else {
			err = errors.New("Id parameter is not ObjectId: " + valstr)
		}
	} else {
		err = errors.New("Missing parameter: id")
	}
	return
}
