package worker

import (
	"sync"

	"bookwormia/book/internal/cache"
	"bookwormia/book/internal/repository"
	"bookwormia/pkg/logger"
)

const (
	CreateOneBookWorker   = 1
	CreateBulkBooksWorker = 2
	EditOneBookWorker     = 3
	EditBulkBooksWorker   = 4
	DeleteOneBookWorker   = 5
	DeleteBulkBooksWorker = 6
)

type Worker struct {
	WorkerType      int
	RedisRepository cache.ICache
	BookRepository  repository.IBook
	stopChan        chan struct{}
}

func NewWorker(workerType int, c cache.ICache, bookRepository repository.IBook) *Worker {
	return &Worker{
		WorkerType:      workerType,
		RedisRepository: c,
		BookRepository:  bookRepository,
		stopChan:        make(chan struct{}),
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	logger.Logger.Info().Int("workerType", w.WorkerType).Msg("worker started!")

	switch w.WorkerType {
	case CreateOneBookWorker:
		w.startCreateOneBookWorker()
	case CreateBulkBooksWorker:
		w.startCreateBulkBookWorker()
	case EditOneBookWorker:
		w.startUpdateOneBookWorker()
	case EditBulkBooksWorker:
		w.startUpdateBulkBookWorker()
	case DeleteOneBookWorker:
		w.startDeleteOneBookWorker()
	case DeleteBulkBooksWorker:
		w.startDeleteBulkBookWorker()

	default:
		logger.Logger.Warn().Msg("unsupported worker type")
		close(w.stopChan)
	}

	<-w.stopChan

	logger.Logger.Info().Msg("worker stoped!")
}

func (w *Worker) Stop() {
	close(w.stopChan)
}
