package audiohandler

type Server struct {
	handler Handler
	getter  GetWriter
}
