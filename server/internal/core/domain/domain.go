package domain

type Domain struct {
	DomainId        string   
	OrgId     		string   
	Name      		string  
	Token     		string   
	State     		string   
}

func DBModelToDomain(model DomainModel) (*Domain, error) {
	return &Domain{
		DomainId:   model.DomainId,
		OrgId:		model.OrgId,  
		Name:       model.Name,     		
		Token:      model.Token,
		State:      model.State,
	}, nil
}
