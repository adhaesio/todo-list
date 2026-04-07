package todo_app

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}
func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.tasks[task.Title]; ok {
		l.mtx.Unlock()
		return ErrTaskAlreadyExists
	}
	l.tasks[task.Title] = task
	return nil
}
func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (l *List) ListTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Task, len(l.tasks))
	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}

func (l *List) UnCompleteTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Uncomplete()
	l.tasks[title] = task
	return l.tasks[title], nil
}

func (l *List) CompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Complete()
	l.tasks[title] = task
	return l.tasks[title], nil
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	delete(l.tasks, title)
	return nil
}

func (l *List) ListUnCompletedTasks() map[string]Task {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	unCompletedTasks := make(map[string]Task)
	for title, task := range l.tasks {
		if !task.Completed {
			unCompletedTasks[title] = task
		}
	}
	return unCompletedTasks

}

// // Структура списка дел
// type TodoList struct {
// 	tasks  []Task
// 	nextID int
// }

// // Конструктор (теперь без файла)
// func NewTodoList() *TodoList {
// 	return &TodoList{
// 		tasks:  []Task{}, // пустой список
// 		nextID: 1,        // начинаем с ID = 1
// 	}
// }

// func (t *Task) Complete() {
// 	doneTime:=time.Now()
// 	t.Completed=true
// 	t.DoneAt=&doneTime

// }

// // Метод добавления задачи
// func (tl *TodoList) Add(text string) int {
// 	task := Task{
// 		ID:        tl.nextID,
// 		Text:      text,
// 		Completed: false,
// 		CreatedAt: time.Now(),
// 	}
// 	tl.tasks = append(tl.tasks, task) // добавляем в массив
// 	tl.nextID++                       // увеличиваем счетчик
// 	return task.ID                    // возвращаем ID новой задачи
// }

// // Метод удаления
// func (tl *TodoList) Delete(id int) error {
// 	for i, task := range tl.tasks {
// 		if task.ID == id {
// 			// Удаляем задачу из среза
// 			// [1, 2, 3, 4, 5] → удаляем 3 → [1, 2, 4, 5]
// 			tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("задача с ID %d не найдена", id)
// }

// // Отметка как выполненной
// func (tl *TodoList) Complete(id int) error {
// 	for i := range tl.tasks {
// 		if tl.tasks[i].ID == id {
// 			tl.tasks[i].Completed = true
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("задача с ID %d не найдена", id)
// }

// // Показать все задачи
// func (tl *TodoList) ShowAll() {
// 	if len(tl.tasks) == 0 {
// 		fmt.Println("📭 Список дел пуст")
// 		return
// 	}

// 	fmt.Println("\n📋 Список дел:")
// 	for _, task := range tl.tasks {
// 		status := " "
// 		if task.Completed {
// 			status = "✓"
// 		}
// 		fmt.Printf("[%s] %d. %s\n",
// 			status,
// 			task.ID,
// 			task.Text,
// 		)
// 	}
// }
