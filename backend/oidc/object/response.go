package object

type Response struct {
	Msg  string `json:"msg"`
	Code int    `json:"status"`
}
