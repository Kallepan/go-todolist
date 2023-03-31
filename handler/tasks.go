package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kallepan/go-backend/db"
	"github.com/kallepan/go-backend/models"
)

var taskIDKey = "taskID"

func tasks(router chi.Router) {
	router.With(paginate).Get("/", listTasks)
	router.Post("/", createTask)
	router.Route("/{taskID}", func(r chi.Router) {
		r.Use(taskContext)
		r.Get("/", getTask)
		r.Put("/", updateTask)
		r.Delete("/", deleteTask)
	})
}

func taskContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		taskID := chi.URLParam(r, "taskID")
		if taskID == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("task ID is missing")))
			return
		}

		id, err := strconv.Atoi(taskID)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("task ID is invalid")))
			return
		}

		ctx := context.WithValue(r.Context(), taskIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createTask(w http.ResponseWriter, r *http.Request) {
	task := &models.Task{}

	if err := render.Bind(r, task); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := dbInstance.AddTask(task); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	render.Status(r, http.StatusCreated)
	if err := render.Render(w, r, task); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	page_size := r.Context().Value("per_page").(int)
	page := r.Context().Value("page").(int)
	tasks, err := dbInstance.GetAllTasks(page, page_size)

	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, tasks); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.Context().Value(taskIDKey).(int)
	task, err := dbInstance.GetTask(taskId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
			return
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &task); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.Context().Value(taskIDKey).(int)
	err := dbInstance.DeleteTask(taskId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
			return
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.Context().Value(taskIDKey).(int)
	taskData := models.Task{}

	if err := render.Bind(r, &taskData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	task, err := dbInstance.UpdateTask(taskId, taskData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &task); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	render.Status(r, http.StatusAccepted)
}

// Paginate middleware
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 1
		}

		perPage, err := strconv.Atoi(r.URL.Query().Get("per_page"))
		if err != nil {
			perPage = 10
		}

		ctx := context.WithValue(r.Context(), "page", page)
		ctx = context.WithValue(ctx, "per_page", perPage)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
