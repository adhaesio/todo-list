package todo_app

import "time"

// Структура задачи
type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompleteAt  *time.Time
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,

		CreatedAt:  time.Now(),
		CompleteAt: nil,
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()
	t.Completed = true
	t.CompleteAt = &completeTime
}

func (t *Task) Uncomplete() {

	t.Completed = false
	t.CompleteAt = nil
}
