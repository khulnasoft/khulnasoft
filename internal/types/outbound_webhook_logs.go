package types

import (
	"time"

	"github.com/khulnasoft/khulnasoft/internal/encryption"
)

type OutboundWebhookLog struct {
	ID                int64
	JobID             int64
	OutboundWebhookID int64
	SentAt            time.Time
	StatusCode        int
	Request           *EncryptableWebhookLogMessage
	Response          *EncryptableWebhookLogMessage
	Error             *encryption.Encryptable
}

const OutboundWebhookLogUnsentStatusCode int = 0
