package engine

type Request struct {
	Url   string
	Parse func(str string) RequestResult
}

type RequestResult struct {
	Items []interface{}
	R     []Request
}

