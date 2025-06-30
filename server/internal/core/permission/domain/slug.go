package domain

import (
	"strings"

	"github.com/server/pkg/utils"
)

type PermissionSlug string 

// converts permission back to its underlying type
func (s PermissionSlug) String() string {
	return string(s)
}

// converts the loosly provided slug to a slug we can parse
// example: app/resource.veiw -> app_resource_veiw
func ToSlug(poorlyfmtslug string) string {
	dilimeters := []string{ 
		utils.ParamDilimeter,
		utils.QueryDilimeter,
		utils.KeyDilimeter,
	}
	
	for _, dilimeter := range dilimeters {
		poorlyfmtslug = strings.Replace(poorlyfmtslug, dilimeter, utils.RelationDilimeter, 2) // replace the first two instances
	}	

	return poorlyfmtslug
}

// TODO move slug into its own thing and finsh the func
// validate func
func (s PermissionSlug) IsValid() error {
	if (s.String() == "") {
		return nil
	}

	return nil
}