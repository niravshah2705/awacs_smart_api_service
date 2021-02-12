package services

import (
	"awacs_smart_api_service/dal"
	"awacs_smart_api_service/graph/model"

	"github.com/brkelkar/common_utils/logger"
)

//GetWorksapceDetails get workspace details by workspace id
func GetWorksapceDetails(Workspace *model.Workspaces, workspaceID string) (err error) {
	clouse := "Workspace='" + workspaceID + "'"
	err = dal.GetWorksapceDetails(Workspace, clouse)
	if err != nil {
		logger.Error("Order details by workspaceId: ", err)
		return
	}
	return
}
