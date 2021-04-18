package main

type Result struct {
	Changes uint `json:"changes"`
}

func newResult(changes uint) *Result {
	return &Result{
		Changes: changes,
	}
}
