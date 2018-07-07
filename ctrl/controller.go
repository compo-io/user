package ctrl

import (
	"github.com/skamenetskiy/jsonapi"
	"github.com/compo-io/user/user"
	"github.com/pkg/errors"
	"github.com/compo-io/db"
)

type Controller struct {
	jsonapi.BaseController
}

// Methods implements jsonapi.Controller
func (c *Controller) Methods() jsonapi.ControllerMethods {
	return jsonapi.ControllerMethods{
		jsonapi.MethodGet: {
			"/id/:id":    c.GetByID,
			"/login/:id": c.GetByLogin,
		},
		jsonapi.MethodPost: {
			"/": c.Create,
		},
	}
}

func (c *Controller) GetByID(ctx *jsonapi.Ctx) *jsonapi.Result {
	id, err := ctx.GetParamUint64("id")
	if err != nil {
		return c.ErrInternalServerError(errors.New("failed to parse user id"))
	}
	u, err := db.Get().FindByPrimaryKeyFrom(user.UserTable, id)
	if err != nil {
		if db.IsErrNoRows(err) {
			return c.ErrNotFound(errors.Wrap(err, "user not found"))
		}
		return c.ErrInternalServerError(errors.Wrap(err, "failed to get user"))
	}
	return c.OK(u.(*user.User))
}

func (c *Controller) GetByLogin(ctx *jsonapi.Ctx) *jsonapi.Result {
	login := ctx.GetParamString("login")
	u, err := db.Get().FindOneFrom(user.UserTable, "login", login)
	if err != nil {
		if db.IsErrNoRows(err) {
			return c.ErrNotFound(errors.Wrap(err, "user not found"))
		}
		return c.ErrInternalServerError(errors.Wrap(err, "failed to get user"))
	}
	return c.OK(u.(*user.User))
}

func (c *Controller) Create(ctx *jsonapi.Ctx) *jsonapi.Result {
	u := new(user.User)
	if err := ctx.ReadJSON(u); err != nil {
		return c.ErrBadRequest(errors.Wrap(err, "failed to parse request json"))
	}
	if err := u.ValidateOnRegistration(); err != nil {
		return c.ErrBadRequest(errors.Wrap(err, ""))
	}
	if err := db.Get().Insert(u); err != nil {
		return c.ErrInternalServerError(errors.Wrap(err, "failed to save user"))
	}
	return c.OK(u)
}
