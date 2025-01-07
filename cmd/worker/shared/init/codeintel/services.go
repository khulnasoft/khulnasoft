package codeintel

import (
	workerdb "github.com/khulnasoft/khulnasoft/cmd/worker/shared/init/db"
	"github.com/khulnasoft/khulnasoft/internal/codeintel"
	"github.com/khulnasoft/khulnasoft/internal/observation"
)

// InitServices initializes and returns code intelligence services.
func InitServices(observationCtx *observation.Context) (codeintel.Services, error) {
	db, err := workerdb.InitDB(observationCtx)
	if err != nil {
		return codeintel.Services{}, err
	}

	codeIntelDB, err := InitDB(observationCtx)
	if err != nil {
		return codeintel.Services{}, err
	}

	return codeintel.NewServices(codeintel.ServiceDependencies{
		DB:             db,
		CodeIntelDB:    codeIntelDB,
		ObservationCtx: observationCtx,
	})
}
