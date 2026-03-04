package main

import (
	"errors"
	"fmt"
)

type Task struct {
	id    int
	title string
	done  bool
}

func (t Task) showTask() string {
	return fmt.Sprintf("Id: %d | Başlık: %s | Tamamlanma durumu: %t", t.id, t.title, t.done)
}

type TaskManager struct {
	tasks  []*Task
	nextID int
}

func (tm *TaskManager) add(title string) *Task {
	tm.nextID += 1
	newTask := Task{id: tm.nextID, title: title, done: false}
	tm.tasks = append(tm.tasks, &newTask)
	return &newTask
}
func (tm TaskManager) list() {
	for i, task := range tm.tasks {
		fmt.Println(i, task.showTask())
	}
}

func (tm *TaskManager) MarkDone(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].id == id {
			tm.tasks[i].done = true
			return nil
		}
	}
	return errors.New("Bulunamadı")
}

func (tm *TaskManager) Delete(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].id == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("Bulunamadı")
}

func (tm TaskManager) Find(id int) (*Task, error) {
	for i := range tm.tasks {
		if tm.tasks[i].id == id {
			return tm.tasks[i], nil
		}
	}
	return nil, errors.New("Task Bulunamadı")
}

func main() {
	task1 := &Task{id: 1, title: "Birinci görev", done: false}
	tasks := []*Task{
		task1,
	}
	tm := TaskManager{tasks: tasks,
		nextID: 1}

	tm.add("Selamün aleyküm")
	finded, err := tm.Find(1)
	if err != nil {
		fmt.Println("Ararken hata oluştu ; ", err)
	} else {
		fmt.Println("Bulundu : ", finded.showTask())
	}

	tm.MarkDone(1)

	tm.list()
}
