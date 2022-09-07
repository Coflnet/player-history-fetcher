package mongo

type TemPlayer struct {
	ID struct {
		ProfileUUID string `bson:"profileUuid"`
		PlayerUUID  string `bson:"playerUuid"`
	} `bson:"_id"`
	GenericItems []interface{} `bson:"generic_items"`
	GenericPets  []interface{} `bson:"generic_pets"`
}
