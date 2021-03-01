package queue_name_enum

type QueueNameEnum struct {
	Name string
}

func InsertMany() QueueNameEnum {
	return QueueNameEnum{
		Name: "InsertMany",
	}
}
