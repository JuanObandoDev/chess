package games

import (
	"github.com/google/uuid"
	"github.com/olahol/melody"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Service struct {
	Websocket *melody.Melody
}

func New() *Service {
	websocket := melody.New()
	return &Service{Websocket: websocket}
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(func(s *Service) {
		s.Websocket.HandleConnect(func(session *melody.Session) {
			id := uuid.NewString()
			session.Set("id", id)
			zap.L().Debug("Websocket " + id + " connected")
		})

		s.Websocket.HandleDisconnect(func(session *melody.Session) {
			id, exists := session.Get("id")
			if !exists {
				return
			}
			zap.L().Debug("Websocket " + id.(string) + " disconnected")
		})

		s.Websocket.HandleMessage(func(session *melody.Session, msg []byte) {
			id, exists := session.Get("id")
			if !exists {
				return
			}
			zap.L().Debug("Websocket "+id.(string)+" message", zap.ByteString("msg", msg))

			s.Websocket.Broadcast(msg)
		})
	}),
)
