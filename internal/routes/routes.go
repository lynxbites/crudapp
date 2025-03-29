package routes

import (
	"crudapp/internal/db"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Flussen/swagger-fiber-v3"
	"github.com/gofiber/fiber/v3"
	_ "github.com/swaggo/files"
)

func NewRouter() *fiber.App {
	router := fiber.New()
	router.Get("/swagger/*", swagger.HandlerDefault)
	router.Get("/tasks", func(c fiber.Ctx) error {

		tasks, err := db.GetTasks()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		json, err := json.Marshal(tasks)
		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.SendString(string(json))
	})

	router.Delete("/tasks/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return fiber.ErrBadRequest
		}

		err = db.DeleteTaskById(id)
		if err != nil {
			return fiber.ErrNotFound
		}

		return c.SendStatus(fiber.StatusOK)
	})

	router.Post("/tasks", func(c fiber.Ctx) error {
		var taskPost db.TaskPost
		reqBody := c.Request().Body()

		fmt.Printf("reqBody: %v\n", string(reqBody))
		err := json.Unmarshal(reqBody, &taskPost)
		if err != nil {

			return fiber.ErrBadRequest
		}

		if !taskPost.Valid() {
			return fiber.ErrBadRequest
		}

		fmt.Printf("taskPost: %v\n", taskPost)
		err = db.PostTask(taskPost)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return fiber.ErrInternalServerError
		}

		return c.SendStatus(fiber.StatusOK)
	})

	router.Put("/tasks/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return fiber.ErrBadRequest
		}

		var taskPut db.TaskPut
		reqBody := c.Request().Body()

		fmt.Printf("reqBody: %v\n", string(reqBody))

		err = json.Unmarshal(reqBody, &taskPut)
		if err != nil {

			return fiber.ErrBadRequest
		}

		if !taskPut.Valid() {
			return fiber.ErrBadRequest
		}
		err = db.PutTask(id, taskPut)
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendStatus(fiber.StatusOK)
	})

	return router
}
