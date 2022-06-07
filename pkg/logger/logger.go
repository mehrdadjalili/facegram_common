package logger

import (
	"context"
	"github.com/mehrdadjalili/facegram_common/pkg/logger/models"
	"github.com/mehrdadjalili/facegram_common/pkg/logger/pb/pd_logger_client"
)

type (
	grpcClient struct {
		url         string
		token       string
		certificate string
		domain      string
	}
	logger struct {
		service    models.Service
		grpcClient grpcClient
	}
	LogDataInfo struct {
		Title       string
		Body        string
		IsEncrypted bool
	}
	LogData struct {
		Section     string
		Function    string
		Message     string
		Information []LogDataInfo
	}
	Logger interface {
		Info(data *LogData)
		Warning(data *LogData)
		Error(data *LogData)
	}
)

func New(url, token, serviceName, serviceIp, certificate, domain string) (Logger, error) {
	return &logger{
		service: models.Service{
			Ip:   serviceIp,
			Name: serviceName,
		},
		grpcClient: grpcClient{
			url:         url,
			token:       token,
			certificate: certificate,
			domain:      domain,
		},
	}, nil
}

func (l *logger) Info(data *LogData) {
	l.submit(pd_logger_client.LogType_INFO, data)
}

func (l *logger) Warning(data *LogData) {
	l.submit(pd_logger_client.LogType_WARNING, data)
}

func (l *logger) Error(data *LogData) {
	l.submit(pd_logger_client.LogType_ERROR, data)
}

func (l *logger) submit(Type pd_logger_client.LogType, data *LogData) {
	conn, client, err := connectToLoggerServer(
		l.grpcClient.url,
		l.grpcClient.token,
		l.grpcClient.certificate,
		l.grpcClient.domain,
	)
	if err != nil {
		return
	}
	defer conn.Close()
	var infor []*pd_logger_client.Info
	for _, item := range data.Information {
		infor = append(infor, &pd_logger_client.Info{
			Title:       item.Title,
			Body:        item.Body,
			IsEncrypted: item.IsEncrypted,
		})
	}
	_, _ = client.Submit(context.Background(), &pd_logger_client.SubmitRequest{
		LogType:     Type,
		Section:     data.Section,
		Function:    data.Function,
		Message:     data.Message,
		Information: infor,
		Service: &pd_logger_client.Service{
			Name: l.service.Name,
			Ip:   l.service.Ip,
		},
	})
}
