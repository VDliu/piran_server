package common

type Res struct {
	Code int
	Msg  string
	Result bool
}

func NewRes(code int, msg string,result bool) *Res {
	return &Res{Code: code, Msg: msg,Result:result}
}

func (res * Res)SetRes(code int, msg string,result bool) {
	res.Set(result)
	res.SetMsg(msg)
	res.SetCode(code)
}

func ( res *Res)SetCode(code int)  {
	res.Code = code
}

func ( res *Res)SetMsg(msg string)  {
	res.Msg = msg
}

func ( res *Res)Set(result bool)  {
	res.Result = result
}

func (err *Res) Error() string {
	return err.Msg
}
