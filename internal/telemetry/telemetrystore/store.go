package telemetrystore

import (
	"github.com/khulnasoft/khulnasoft/internal/database"
	"github.com/khulnasoft/khulnasoft/internal/telemetry"
	"github.com/khulnasoft/khulnasoft/internal/telemetry/telemetrystore/teestore"
)

// New creates a default EventStore. Most callers should not use this directly and instead use
// `telemetryrecorder.New`.
//
// The current default tees events to both the legacy event_logs table, as well
// as the new Telemetry Gateway export queue.
func New(exportQueue database.TelemetryEventsExportQueueStore, eventLogs database.EventLogStore) telemetry.EventsStore {
	return teestore.NewStore(exportQueue, eventLogs)
}
