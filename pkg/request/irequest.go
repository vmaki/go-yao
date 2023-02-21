package request

type IRequest interface {
	Generate(data interface{}) string
}
