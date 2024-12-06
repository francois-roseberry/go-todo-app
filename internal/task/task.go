package task

var taskNames = []string{
	"Call mom",
	"Sort Pokemon cards",
	"Conquer the world",
	"Register for next Software Crafters meetup",
	"Go shopping for Christmas gifts",
	"Buy milk",
	"Print latest model",
	"Finish HTMX app",
	"Ship printed model to customer X",
	"Water the plants",
	"Visit cathedral",
	"Climb Mt Everest",
	"Learn manarin",
	"Go on a walk",
	"Clean the house",
	"Go bungee jumping",
	"Drink a beer on Oktoberfest",
	"Cook spaghetti",
}

type Task struct {
	Id   int
	Name string
	Done bool
}

func NewTask(index int) *Task {
	return &Task{
		Id:   index,
		Name: taskNames[index%len(taskNames)],
	}
}

func TaskList() []*Task {
	tasks := []*Task{}
	for i := 1; i < 10; i++ {
		tasks = append(tasks, NewTask(i))
	}
	return tasks
}
