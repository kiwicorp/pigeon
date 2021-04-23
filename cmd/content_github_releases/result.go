package main

type Result struct {
	Changed bool `json:"changed"`
}

func newResult(changed bool) *Result {
	return &Result{
		Changed: changed,
	}
}
