package prefrence

type Preference struct {
	PrefrenceId  string  	   
	Name         string    
	Value        string    
	ResourceType string   
	ResourceId   string  
}

func DBModelToPrefrence(model PreferenceModel) (*Preference, error) {
	return &Preference{
		PrefrenceId:   model.PrefrenceId,
		Name:          model.Name,
		Value:         model.Value,
		ResourceType:  model.ResourceType,
		ResourceId:    model.ResourceId,
	}, nil
}
