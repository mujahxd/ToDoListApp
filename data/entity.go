package data

type User struct {
	userID    int
	Fullname  string
	Username  string
	Password  string
	Create_at int
	Update_at int
}
type Task struct {
	taksID      int
	Name        string
	Status      bool
	Describtion string
	Create_at   int
	Update_at   int
}

type Plan struct {
	planID      int
	Name        string
	Describtion string
	Create_at   int
	Update_at   int
}
