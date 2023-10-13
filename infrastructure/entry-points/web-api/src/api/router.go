package api

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type BudgetRouter struct {
	engine  *gin.Engine
	handler BudgetHandler
}

func ConfigRouter(handler BudgetHandler) *BudgetRouter {
	router := &BudgetRouter{
		engine:  gin.Default(),
		handler: handler,
	}

	router.route(http.MethodGet, "/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "UP"})
	})
	router.route(http.MethodGet, "/favicon.ico", func(gc *gin.Context) {
		pwd, _ := os.Getwd()
		fileBytes, err := os.ReadFile(path.Join(pwd, "static", "assets", "favicon.ico"))
		if err != nil {
			gc.JSON(http.StatusInternalServerError, gin.H{"error": "error reading favicon.ico file."})
			return
		}

		contentType := http.DetectContentType(fileBytes)
		gc.Data(http.StatusOK, contentType, fileBytes)
	})

	expenseRoute := "/budget/expense"

	router.route(http.MethodGet, expenseRoute, router.handler.FindExpense)
	router.route(http.MethodPost, expenseRoute, router.handler.SaveExpense)
	router.route(http.MethodPut, expenseRoute, router.handler.UpdateExpense)
	router.route(http.MethodDelete, expenseRoute, router.handler.DeleteExpense)

	return router
}

func (r *BudgetRouter) route(method, path string, handle func(*gin.Context)) {
	switch method {
	case http.MethodGet:
		r.engine.GET(path, handle)
	case http.MethodPost:
		r.engine.POST(path, handle)
	case http.MethodPatch:
		r.engine.PATCH(path, handle)
	case http.MethodPut:
		r.engine.PUT(path, handle)
	case http.MethodDelete:
		r.engine.DELETE(path, handle)
	default:
		r.engine.GET(path, handle)
	}
}

func (r *BudgetRouter) Run() {
	r.engine.Run()
}
