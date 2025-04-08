package control

import (
	"io"
	"log"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	controlpb "github.com/SyntinelNyx/syntinel-agent/internal/proto/controlpb"
)

type Agent struct {
	controlpb.UnimplementedAgentServiceServer
}

func (a *Agent) Control(stream controlpb.AgentService_ControlServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Recv error: %v", err)
			return err
		}

		logger.Info("Received: command=%s payload=%s misc=%t", msg.Command, msg.Payload, len(msg.Misc) != 0)

		var result string
		switch msg.Command {
		case "heartbeat":
			result = "<3"
		case "exec":

		case "download":

		default:
			result = "unknown command"
		}

		err = stream.Send(&controlpb.ControlResponse{
			Result: result,
			Status: "ok",
		})
		if err != nil {
			log.Printf("Send error: %v", err)
			return err
		}
	}
}
