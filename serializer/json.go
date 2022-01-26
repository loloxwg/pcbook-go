package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJson(message proto.Message) (string, error) {
	marshalOptions := protojson.MarshalOptions{
		Indent:          " ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
		UseEnumNumbers:  false,
	}
	return marshalOptions.Format(message), nil

}
