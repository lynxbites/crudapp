package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var pool *pgxpool.Pool

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("CONNSTR")
	fmt.Print("DATABASE STRING: ", connStr)
	if err != nil {
		log.Fatalf("Error when parsing .env")
	}
	pool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
}

// Tasks godoc
// @Summary      GET
// @Description  Get all tasks
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200 {object} []Task "OK"
// @Failure      400  "Bad Request"
// @Failure      404  "Not Found"
// @Failure      405  "Method not allowed"
// @Router       /tasks [get]
func GetTasks() ([]Task, error) {
	rows, err := pool.Query(context.Background(), "select * from tasks")

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[Task])
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return []Task{}, err
	}
	return tasks, nil
}

// Tasks godoc
// @Summary      DELETE
// @Description  Delete task by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  "OK"
// @Failure      400  "Bad Request"
// @Failure      404  "Not Found"
// @Failure      405  "Method not allowed"
// @Router       /tasks/{id} [delete]
func DeleteTaskById(id int) error {

	if !IsExistById(id) {
		return errors.New("row doesn't exist")
	}

	_, err := pool.Exec(context.Background(), "delete from tasks where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Tasks godoc
// @Summary      POST
// @Description  Add a new task to the list
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        Task   body      TaskPost  true  "Task Info"
// @Success      200  "OK"
// @Failure      400  "Bad Request"
// @Failure      500  "Internal Server Error"
// @Router       /tasks [post]
func PostTask(task TaskPost) error {
	_, err := pool.Exec(context.Background(), "insert into tasks (title, description) values ($1,$2)", task.Title, task.Description)
	if err != nil {
		return err
	}
	return nil
}

// Tasks godoc
// @Summary      PUT
// @Description  Update an existing task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param   	 id      path     int     true  "Id of a task."
// @Param        Task   body      TaskPut  true  "Task Info"
// @Success      200  "OK"
// @Failure      400  "Bad Request"
// @Router       /tasks/{id} [put]
func PutTask(id int, task TaskPut) error {

	if !IsExistById(id) {
		return errors.New("row doesn't exist")
	}

	_, err := pool.Exec(context.Background(), "update tasks set title=$2, description=$3, status=$4, updated_at=$5 where id=$1", id, task.Title, task.Description, task.Status, time.Now().Local())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return nil
}

func IsExistById(id int) bool {
	var exists bool
	pool.QueryRow(context.Background(), "select exists (select * from tasks where id = $1)", id).Scan(&exists)
	return exists
}
