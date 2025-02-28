package resolver

import "github.com/zhenyanesterkova/nepblog/internal/gql/model"

func NewInternalErrorProblem() model.InternalErrorProblem {
	return model.InternalErrorProblem{Message: "internal server error"}
}
