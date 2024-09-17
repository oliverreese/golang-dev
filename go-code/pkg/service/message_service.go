package service

type MessageService struct {}

func(m *MessageService) GetMessage() string {
    return "Hello World From Container"
}