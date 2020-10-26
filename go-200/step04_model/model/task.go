package model

import "time"

// TODO build the TaskStatus type as an int
// TaskStatus is the current processing status of a task

// TODO define the Task status enum as const using iota (StatusTodo, StatusInProgress, StatusDone)
const ()

// TODO build the TaskPriority type as an int
// TaskPriority is the priority of a task

// TODO define the Task Priority enum as const using iota (PriorityMinor, PriorityMedium, PriorityHigh)
const ()

// TODO add the Status and Priority enums, the Creation and Due Dates and the JSON ans BSON annotations

// Task is the structure to define a task to be done
type Task struct {
	ID          string `json:"id,omitempty" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	// TODO Status
	// TODO Priority
	// TODO Creation Date
	DueDate time.Time `json:"due_date" bson:"due_date"`
}

// TODO add a NewTask function to create a new pointer to a task when called
// TODO function takes 2 arguments as string : title and description
// TODO function sets a new UUID, a creation date and other default param initalisation
// NewTask builds a new task with a new ID of the Task as a string

// TODO add an Equal method for Task comparison, be careful with time.Time comparison
