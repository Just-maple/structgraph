package utils

import (
	`net/http`
)

type HttpClient struct {
	*http.Client
}
