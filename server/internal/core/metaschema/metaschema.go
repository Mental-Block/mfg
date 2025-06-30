package metaschema

type MetaSchema struct {
	MetaSchemaId    string    
	Name      		string    			
	Schema    		string  
}

func DBModelToMetaSchema(model MetaSchemaModel) (*MetaSchema, error) {
	return &MetaSchema{
		MetaSchemaId:    model.Id,
		Name:      		 model.Name,
		Schema:  		 model.Schema,
	}, nil
}
