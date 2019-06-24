package process

type ETHRequest struct {
	Jsonrpc 	string
	Method 		string
	Params		[]string
	Id 			string
}

type ETHListAccount struct {
	Jsonrpc 	string
	Id 			int64
	Result		[]string
}

type ETHNewAccount struct {
	Jsonrpc 	string
	Id 			int64
	Result		string
}

type ETHCommonResponse struct {
	Jsonrpc 	string
	Id 			int64
	Result		string
}