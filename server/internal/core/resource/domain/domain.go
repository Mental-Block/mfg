package domain

import "github.com/server/pkg/metadata"

type Resource struct {
	Id            	string        
	URN           	string        
	Name          	string        
	Title         	string        
	ProjectId     	string        
	NamespaceId   	string        
	PrincipalId   	string
	PrincipalType 	string
	Metadata      	metadata.Metadata         
}

type SanitizedResource struct {
	Id            	string        
	URN           	string        
	Name          	string        
	Title         	string        
	ProjectId     	string        
	NamespaceId   	string        
	PrincipalId   	string
	PrincipalType 	string
	Metadata      	metadata.Metadata 
}