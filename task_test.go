package kanban

import (
    "testing"
)

func TestTaskCanExecute(t *testing.T) {
    task := Task {
        Instructions: []string {
            "echo hello",
        },
    }
    output, oops := task.Execute()
    if (oops != nil) || (output != "hello") {
        t.Error("wrong output")
    }
}
