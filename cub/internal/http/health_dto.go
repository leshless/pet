package http

import (
	api "github.com/leshless/pet/cub/api/http/v1"
	healthlogic "github.com/leshless/pet/cub/internal/logic/health"
	"github.com/leshless/pet/cub/internal/model"
)

var healthStatusFromModel = map[model.HealthStatus]api.HealthStatus{
	model.HealthStatusUnknown:    api.UNKNOWN,
	model.HealthStatusServing:    api.SERVING,
	model.HealthStatusNotServing: api.NOTSERVING,
}

func checkHealthRequestToDTO(req api.CheckHealthRequestObject) healthlogic.CheckArg {
	return healthlogic.NewCheckArg()
}

func checkHealthResponseFromDTO(res healthlogic.CheckRes) api.CheckHealth200JSONResponse {
	status, ok := healthStatusFromModel[res.Status]
	if !ok {
		status = api.UNKNOWN
	}

	return api.CheckHealth200JSONResponse{
		Status: status,
	}
}
