package port

type IdGenerator interface {
	GenerateID() string
}
