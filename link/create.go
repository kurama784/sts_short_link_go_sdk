package link

type CreateDto struct {
	Code        string `json:"code"`
	Domain      string `json:"domain"`
	ExpiredAt   *int64 `json:"expired_at"`
	RedirectURL string `json:"redirect_url"`
}
