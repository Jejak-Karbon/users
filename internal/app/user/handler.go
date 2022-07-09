package user

import (
	"strconv"
	
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) GetByID(c echo.Context) error {

	id,_:= strconv.Atoi(c.Param("id"))
	result, err := h.service.FindByID(c.Request().Context(), uint(id))
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Update(c echo.Context) error {
	payload := new(dto.UpdatePasswordRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	payloadToken := middleware.GetIDFromToken(c)
	var user_id uint = uint(payloadToken.(float64))

	result, err := h.service.Update(c.Request().Context(), user_id, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}