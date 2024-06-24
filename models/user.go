package models

//bson is the format understood by mongodb
import "time"

// type LogData struct{ 
// 	Id bson.ObjectId  `json:"id" bson:"_id"` //we are basically telling that, in JSON it will look like this, and in mongoDB, it will look like this
// 	Name string       `json:"name" bson:"name"`
// 	Gender string     `json:"gender" bson:"gender"`
// 	Age int           `json:"age" bson:"age"`
// }

type LogData struct {
	Level       string    `json:"level"`
	Message     string    `json:"message"`
	ResourceID  string    `json:"resourceId"`
	Timestamp   time.Time `json:"timestamp"`
	TraceID     string    `json:"traceId"`
	SpanID      string    `json:"spanId"`
	Commit      string    `json:"commit"`
	Metadata    Metadata  `json:"metadata"`
}

type Metadata struct {
	ParentResourceID string `json:"parentResourceId"`
}