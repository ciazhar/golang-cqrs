package error

type Status string

const (
	WrongInput Status = "01"
	NotExist   Status = "02"
)

type Error struct {
	Error  string `json:"error"`
	Status Status `json:"status"`
}

func New(err error) Error {
	return Error{Error: err.Error()}
}

func NewS(err error, status Status) Error {
	return Error{Error: err.Error(), Status: status}
}
