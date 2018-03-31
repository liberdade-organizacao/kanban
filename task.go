package kanban

type Task struct {
    Instructions []string
}

// This function executes the Instructions from a task structure.
func (task *Task) Execute() (string, error) {
    return "hello", nil
}
