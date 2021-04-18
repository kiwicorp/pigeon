package main

func validateParams(params Params) *Error {
	errs := make([]error, 3)
	if params.AccessToken == "" {
		errs = append(errs, ErrMissingAccessToken)
	}
	if params.RepoOwner == "" {
		errs = append(errs, ErrMissingRepoOwner)
	}
	if params.RepoName == "" {
		errs = append(errs, ErrMissingRepoName)
	}

	if len(errs) > 0 {
		return &Error{
			Inner: errs,
		}
	}
	return nil
}

type Params struct {
	AccessToken string `json:"access_token"`
	RepoOwner   string `json:"repo_owner"`
	RepoName    string `json:"repo_name"`
}
