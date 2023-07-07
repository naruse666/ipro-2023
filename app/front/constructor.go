package main 

import "os"

type backendHosts struct {
	SuicideHost string
}

func NewBackendHosts () *backendHosts {
	hosts := new(backendHosts)

	// suicide host
	hosts.SuicideHost = os.Getenv("SUICIDE_HOST")
	if hosts.SuicideHost == "" {
		hosts.SuicideHost = "localhost"
	}

	return hosts
}
