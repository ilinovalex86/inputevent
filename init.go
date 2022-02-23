package inputevent

type Event struct {
	Method string
	Event  string
	Key    string
	Code   string
	CorX   int
	CorY   int
	Ctrl   bool
	Shift  bool
}
