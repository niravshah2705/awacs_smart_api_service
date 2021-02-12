package dal

import (
	"awacs_smart_api_service/graph/model"

	db "github.com/brkelkar/common_utils/databases"
)

//GetWorksapceDetails get workspace data
func GetWorksapceDetails(Workspace *model.Workspaces, clouse string) (err error) {	
	err = db.DB["smartdb"].Where(clouse).Find(&Workspace).Error
	return
}