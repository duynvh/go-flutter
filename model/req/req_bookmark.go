package req

type ReqBookMark struct {
	RepoName string `json:"repo,omitempty" validate:"required"`
}