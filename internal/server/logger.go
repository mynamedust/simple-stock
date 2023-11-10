package server

import "go.uber.org/zap"

func (s *Server) initLogger() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	s.logger = logger.Sugar()
	return nil
}
