package types

import "time"

// Executor describes an executor instance that has recently connected to Khulnasoft.
type Executor struct {
	ID              int
	Hostname        string
	QueueName       string
	QueueNames      []string
	OS              string
	Architecture    string
	DockerVersion   string
	ExecutorVersion string
	GitVersion      string
	IgniteVersion   string
	SrcCliVersion   string
	FirstSeenAt     time.Time
	LastSeenAt      time.Time
}
