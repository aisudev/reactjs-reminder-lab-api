package delivery

import (
	"net/http"
	"reminder/domain"

	"github.com/labstack/echo"
)

type Handler struct {
	usecase domain.TodoUsecase
}

func NewHandler(e *echo.Group, u domain.TodoUsecase) *Handler {

	h := Handler{usecase: u}

	e.GET("/todo", h.GetAll)
	e.GET("/todo/:uuid", h.Get)
	e.POST("/todo", h.Create)
	e.PUT("/todo/:uuid", h.Update)
	e.DELETE("/todo/:uuid", h.Delete)

	return &h
}

func (h *Handler) Create(e echo.Context) error {

	todo := domain.Todo{}

	if err := e.Bind(&todo); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false, "error": err,
		})
	}

	if err := h.usecase.Create(&todo); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false, "error": err,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) GetAll(e echo.Context) error {
	todo := []domain.Todo{}
	var err error

	if todo, err = h.usecase.GetAll(); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false, "error": err,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": true, "data": todo,
	})

}

func (h *Handler) Get(e echo.Context) error {

	uuid := e.Param("uuid")

	var todo *domain.Todo
	var err error

	if todo, err = h.usecase.Get(uuid); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false, "error": err,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": true, "data": todo,
	})

}

func (h *Handler) Update(e echo.Context) error {

	uuid := e.Param("uuid")

	todo := domain.Todo{}

	if err := e.Bind(&todo); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false, "error": err,
		})
	}

	if err := h.usecase.Update(uuid, &todo); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false, "error": err,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": true,
	})

}

func (h *Handler) Delete(e echo.Context) error {

	uuid := e.Param("uuid")

	if err := h.usecase.Delete(uuid); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false, "error": err,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": true,
	})

}
