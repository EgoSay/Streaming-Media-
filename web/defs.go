/**
 * @Author codeAC
 * @Time: 2018/8/15 17:44
 * @Description
 */
package main

type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecognized   = Err{Error: "api not recognized, bad resquest", ErrorCode: "001"}
	ErrorRequestBodyParseFailed = Err{Error: "request body is not correct", ErrorCode: "002"}
	ErrorInternalFaults         = Err{Error: "internal service error", ErrorCode: "003"}
)
