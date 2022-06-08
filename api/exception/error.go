package exception

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

type Model struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func New(message string) Model {
	panic(Model{Message: message})
}
