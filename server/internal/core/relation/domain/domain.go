package domain

import (
	namespace "github.com/server/internal/core/namespace/domain"
	"github.com/server/pkg/utils"
)

type Filter struct {
	Subject Subject
	Object  Object
}

type CheckPair struct {
	Relation Relation
	Status   bool
}

type Object struct {
	Id        string
	Namespace string
}

var SubjectAll = utils.EverythingDilimeter

type Subject struct {
	Id              string
	Namespace       string
	SubRelationName string
}

type Relation struct {
	Id           string
	RelationName string
	Object       Object
	Subject      Subject
}

func (r Relation) Transform()(SanitizedRelation, error){

	UUID, err := utils.ConvertStringToUUID(r.Id) 
	
	if err != nil {
		return SanitizedRelation{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	relationName := RelationName(r.RelationName)

	if err := relationName.IsValid(); err != nil {
		return SanitizedRelation{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, err.Error()) 
	}

	objectUUID, err := utils.ConvertStringToUUID(r.Object.Id) 
	
	if err != nil {
		return SanitizedRelation{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}

	objectNameSpace := namespace.NameSpaceName(r.Object.Namespace)

	if err := objectNameSpace.IsValid(); err != nil {
		return SanitizedRelation{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, err.Error()) 
	}

	if !utils.IsNullUUID(r.Subject.Id) && r.Subject.Id != SubjectAll {
		return SanitizedRelation{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, utils.ErrInvalidUUIDFormat, err)
	}
 
	subjectNameSpace := namespace.NameSpaceName(r.Subject.Namespace)

	if err := subjectNameSpace.IsValid(); err != nil {
		return SanitizedRelation{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, err.Error()) 
	}

	subjectSubRelationName :=  RelationName(r.Subject.SubRelationName)

	if err := subjectSubRelationName.IsValid(); err != nil {
		return SanitizedRelation{}, utils.NewErrorf(utils.ErrorCodeInvalidArgument, err.Error()) 
	}
	
	return SanitizedRelation{
		Id: UUID,
		RelationName: relationName,
		Object: SanitizedObject{
			Id: objectUUID,
			Namespace: objectNameSpace,
		},
		Subject: SanitizedSubject{
			Id: r.Subject.Id,
			Namespace: subjectNameSpace,
			SubRelationName: subjectSubRelationName,
		},
	}, nil
}

type SanitizedObject struct {
	Id        utils.UUID
	Namespace namespace.NameSpaceName
}

type SanitizedSubject struct {
	Id              string
	Namespace       namespace.NameSpaceName
	SubRelationName RelationName
}

type SanitizedRelation struct {
	Id           utils.UUID
	RelationName RelationName
	Object       SanitizedObject
	Subject      SanitizedSubject
}

func (r SanitizedRelation) Transform()(*Relation){
	return &Relation{
		Id: r.Id.String(),
		RelationName: r.RelationName.String(),
		Object: Object{
			Id: r.Object.Id.String(),
			Namespace: r.Object.Namespace.String(),
		},
		Subject: Subject{
			Id: r.Subject.Id,
			Namespace: r.Subject.Namespace.String(),
			SubRelationName: r.Subject.SubRelationName.String(),
		},
	}
}

