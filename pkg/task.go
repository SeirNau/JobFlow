package pkg

type Task struct {
	ID              string
	Title           string
	Priority        int
	Done            bool
	EstimateMinutes int
}

func (t *Task) Complete() {
	t.Done = true
}
