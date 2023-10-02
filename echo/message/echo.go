package message

type Echo struct {
	Type string
	Body string
}

func (echo Echo) GetType() string {
	return "echo"
}

func (echo Echo) GetBody() []byte {
	return []byte(echo.Body)
}
