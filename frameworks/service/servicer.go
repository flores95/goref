package service

type (
	// HealthIndicator is used for anything who's system health can be verified
	HealthIndicator interface {
		Health() (Status, error)
	}

	// Status represents the status/health of a service
	Status int

	// Servicer represents a servicer (something that provides a service).
	// This is typically an API, REST or Microservice endpoint
	Servicer interface {
		HealthIndicator
		Start()
	}
)

// constants for Status o
const (
	Down Status = 0
	Up   Status = 1
)

// String is used to give a text representation of a status
func (status Status) String() string {
	levels := [...]string{
		"Down",
		"Up",
	}
	return levels[status]
}
