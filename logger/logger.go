package logger

import (
	"os"
	"tsProject/models"

	"github.com/sirupsen/logrus"
)

// Logger yapısı, Logrus logger'ı sarmalar
type Logger struct {
	logger *logrus.Logger
}

// NewLogger, yeni bir Logger örneği oluşturur ve JSON formatında loglar
func NewLogger(logFile string) *Logger {
	// Logrus logger'ı oluştur
	log := logrus.New()

	// Dosyaya log yazmak için dosya aç
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Logger'ı yapılandır (JSON format)
	log.SetOutput(file)
	log.SetFormatter(&logrus.JSONFormatter{})

	// Logger seviyesini ayarla (opsiyonel)
	log.SetLevel(logrus.InfoLevel)

	return &Logger{
		logger: log,
	}
}

// Info, InfoLog struct kullanarak loglama yapar
func (l *Logger) Info(log models.InfoLog) {
	l.logger.WithFields(logrus.Fields{
		"username":    log.UserName,
		"user_id":      log.UserID,
		"event":        log.Event,
		"product_id":   log.ProductID,
		"product_name": log.ProductName,
	}).Info(log.Message)
}

// Warn, WarnLog struct kullanarak loglama yapar
func (l *Logger) Warn(log models.WarnLog) {
	l.logger.WithFields(logrus.Fields{
		"warning": log.Warning,
		"context": log.Context,
	}).Warn(log.Message)
}

// Error, ErrorLog struct kullanarak loglama yapar
func (l *Logger) Error(log models.ErrorLog) {
	l.logger.WithFields(logrus.Fields{
		"error":   log.Error,
		"context": log.Context,
	}).Error(log.Message)
}

func (l *Logger) CycleInfoLog(log models.CycleInfoLog) {
	l.logger.WithFields(logrus.Fields{
		"event":      log.Event,
		"product_id": log.ProductID,
		"year":       log.Year,
		"month":      log.Month,
		"cycle_count":   log.CycleCount,
	}).Info(log.Message)
}

func (l *Logger) InstitutionInfoLog(log models.InstitutionInfoLog) {
	l.logger.WithFields(logrus.Fields{
		"event":            log.Event,
		"institution_name": log.InstitutionName,
		"city":             log.City,
	}).Info(log.Message)
}

func (l *Logger) ProductInfoLog(log models.ProductInfoLog) {
	l.logger.WithFields(logrus.Fields{
		"event":          log.Event,
		"product_id":     log.ProductID,
		"product_name":   log.ProductName,
		"product_serial": log.ProductSerial,
		"institution_id": log.InstitutionID,
		"owner":          log.OwnerID,
		"status":         log.Status,
		"responsible":    log.ResponsibleID,
	}).Info(log.Message)
}

func (l *Logger) UserInfoLog(log models.UserInfoLog) {
	l.logger.WithFields(logrus.Fields{
		"event":     log.Event,
		"user_id":   log.UserID,
		"username": log.UserName,
		"role":      log.Role,
	}).Info(log.Message)
}
