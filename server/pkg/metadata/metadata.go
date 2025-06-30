package metadata

import (
	"encoding/json"

	"google.golang.org/protobuf/types/known/structpb"
)

// Metadata is a structure to store dynamic values. it could be use as an additional information of a specific entity
type Metadata map[string]any

// ToStructPB transforms Metadata to *structpb.Struct
func (m Metadata) ToStructPB() (*structpb.Struct, error) {
	newMap := make(map[string]any)

	for key, value := range m {
		newMap[key] = value
	}

	return structpb.NewStruct(newMap)
}

// Build transforms a Metadata from map[string]any
func Build(m map[string]any) Metadata {
	newMap := make(Metadata)
	for key, value := range m {
		newMap[key] = value
	}
	return newMap
}

// FromString transforms a Metadata from map[string]string
func FromString(m map[string]string) Metadata {
	newMap := make(Metadata)
	for key, value := range m {
		newMap[key] = value
	}
	return newMap
}

//  Unserializes metadata to map 
func UnMarshallMetadata[T Metadata](metadata []byte) (T, error) {
	newMetaData := T{}
	
	if len(metadata) > 0 {
		if err := json.Unmarshal(metadata, &newMetaData); err != nil {
			return nil, err
		}
	}

	return newMetaData, nil
}

//  Unserializes metadata to any array
func UnMarshallArray[T comparable](array []byte) ([]T, error) {
	newArray := []T{}

	if len(array) > 0 {
		if err := json.Unmarshal(array, &newArray); err != nil {
			return nil, err
		}
	}

	return newArray, nil
}

//loops through metadata and converts it to an array base of the given key
func MapToArray(metadata Metadata, key string) []string {
	var collection []string

	if val, ok := metadata[key]; ok && (val != nil) {
		for _, rawValue := range val.([]interface{}) {
			collection = append(collection, rawValue.(string))
		}
	}

	return collection 
}
