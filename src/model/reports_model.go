package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	FaultID  int                `bson:"fault_id"`
	SensorID int                `bson:"sensor_id"`
	Status   int                `bson:"status"`
	Tags     []string           `bson:"tags"`
}
