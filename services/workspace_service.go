package services

import (
	"awacs_smart_api_service/graph/model"

	"github.com/brkelkar/common_utils/logger"

	db "github.com/brkelkar/common_utils/databases"
)

//GetWorksapceDetails get workspace details by workspace id
func GetWorksapceDetails(Workspace *model.Workspaces, workspaceID string) (err error) {
	err = db.DB["smartdb"].Where("Workspace='" + workspaceID + "'").Find(&Workspace).Error
	if err != nil {
		logger.Error("Order details by workspaceId: ", err)
		return
	}
	return
}
