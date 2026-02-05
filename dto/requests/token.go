package requests

type TokenReq struct {
	Duration  int `json:"duration"`
	LateAfter int `json:"late_after"`
}

type SubmitToken struct {
	TokenCode string `json:"token_code"`
}
