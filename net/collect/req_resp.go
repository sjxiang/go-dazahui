package collect




// 请求 Body
type Request struct {
	Source string `json:"source"`
	Words []string `json:"words"`
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
}


// 响应 Body
type Response struct {
	Details []Details `json:"details"`
	BaseResp BaseResp `json:"base_resp"`
}

type Details struct {
	Detail string `json:"detail"`
	Extra  string `json:"extra"`
}

type BaseResp struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
