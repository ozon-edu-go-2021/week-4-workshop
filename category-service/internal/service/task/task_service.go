package task

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Service struct {
	repository RepositoryInterface

	mutex         *sync.Mutex
	uniqueTaskIDs map[uint64]struct{}
}

type RepositoryInterface interface {
	FindNonStartedTask(context.Context) (*Task, error)
}

func New(repository RepositoryInterface) Service {
	return Service{
		repository: repository,

		mutex:         new(sync.Mutex),
		uniqueTaskIDs: make(map[uint64]struct{}),
	}
}

func (s Service) ExecTask(ctx context.Context) error {
	task, err := s.repository.FindNonStartedTask(ctx)
	if err != nil {
		return errors.Wrap(err, "repository.FindNonStartedTask()")
	}

	s.mutex.Lock()
	if _, ok := s.uniqueTaskIDs[task.ID]; ok {
		log.Fatal("Дубликат")
	}
	s.uniqueTaskIDs[task.ID] = struct{}{}
	s.mutex.Unlock()

	time.Sleep(1 * time.Second)

	return nil
}
