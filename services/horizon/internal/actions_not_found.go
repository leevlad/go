package horizon

import (
	"github.com/leevlad/go/services/horizon/internal/actions"
	"github.com/leevlad/go/support/render/problem"
)

// Interface verification
var _ actions.JSONer = (*NotFoundAction)(nil)

// NotFoundAction renders a 404 response
type NotFoundAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotFoundAction) JSON() error {
	problem.Render(action.R.Context(), action.W, problem.NotFound)
	return action.Err
}
