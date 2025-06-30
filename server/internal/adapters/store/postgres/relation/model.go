package relation_store

import (
	"database/sql"
	"time"

	namespace "github.com/server/internal/core/namespace/domain"

	"github.com/server/internal/core/relation/domain"
	"github.com/server/pkg/utils"
)

type RelationModel struct {
	Id                   utils.UUID          				`db:"id"`
	RelationName         domain.RelationName         		`db:"relation_name"`
	SubjectId            sql.NullString         			`db:"subject_id"`				// can be nil
	SubjectNamespace     namespace.NameSpaceName            `db:"subject_namespace_name"`	
	SubjectSubRelation   *domain.RelationName 				`db:"subject_subrelation_name"` // can be nil
	ObjectId             utils.UUID         				`db:"object_id"`
	ObjectNamespace      namespace.NameSpaceName         	`db:"object_namespace_name"`
	CreatedDT            time.Time      					`db:"created_dt"`
	CreatedBy			 string								`db:"created_by"`
	UpdatedDT            sql.NullTime       				`db:"updated_dt"` // can be nil
	UpdatedBy			 sql.NullString						`db:"updated_by"` // can be nil
	DeletedDT            sql.NullTime   					`db:"deleted_at"` // can be nil
	DeletedBy			 sql.NullString						`db:"deleted_dt"` // can be nil
}

func (r RelationModel) Transform() *domain.SanitizedRelation {
	return &domain.SanitizedRelation{
		Id: r.Id,
		RelationName: r.RelationName,
		Subject: domain.SanitizedSubject{
			Id:              r.SubjectId.String,
			Namespace:       r.SubjectNamespace,
			SubRelationName: *r.SubjectSubRelation,
		},
		Object: domain.SanitizedObject{
			Id:        r.ObjectId,
			Namespace: r.ObjectNamespace,
		},
	}
}
