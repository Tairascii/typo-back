package typo_back

import "gopkg.in/mgo.v2/bson"

type UserModel struct {
	Id       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
}

type User struct {
	Id       string `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}
