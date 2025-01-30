package service

import (
	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/dto"
	log "github.com/fydhfzh/ecommerce-go-application/src/logger-service/model"
	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/repository"
)

type logService struct {
	logRepository repository.LogRepository
}

type LogService interface {
	Save(logRequest dto.LogRequest) (*dto.LogResponse, error)
	GetAll() ([]dto.LogResponse, error)
}

func NewLogService(logRepository repository.LogRepository) LogService {
	return &logService{
		logRepository: logRepository,
	}
}

func (l *logService) Save(logRequest dto.LogRequest) (*dto.LogResponse, error) {
	log := log.NewLog(logRequest.Content)

	newLog, err := l.logRepository.Save(log)
	if err != nil {
		return nil, err
	}

	logResponse := &dto.LogResponse{
		ID:        newLog.ID,
		Content:   newLog.Content,
		CreatedAt: newLog.CreatedAt.String(),
		UpdatedAt: log.UpdatedAt.String(),
	}

	return logResponse, nil
}

func (l *logService) GetAll() ([]dto.LogResponse, error) {
	logs, err := l.logRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var logsResponse []dto.LogResponse

	for _, log := range logs {
		logResponse := dto.LogResponse{
			ID:        log.ID,
			Content:   log.Content,
			CreatedAt: log.CreatedAt.String(),
			UpdatedAt: log.UpdatedAt.String(),
		}

		logsResponse = append(logsResponse, logResponse)
	}

	return logsResponse, nil
}
