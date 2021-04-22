package model

import "strings"

const (
	// Namespace id.
	nid string = "pigeon-selftech-io"
)

type Urn interface {
	Nss() string
}

// StringUrn converts a urn to a string.
func StringUrn(urn Urn) string {
	parts := []string{
		"urn",
		nid,
		urn.Nss(),
	}
	return strings.Join(parts, ":")
}
