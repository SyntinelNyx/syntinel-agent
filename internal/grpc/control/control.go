package control

import (
	"io"

	"github.com/SyntinelNyx/syntinel-agent/internal/commands"
	"github.com/SyntinelNyx/syntinel-agent/internal/data"
	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	controlpb "github.com/SyntinelNyx/syntinel-agent/internal/proto/controlpb"
	"github.com/SyntinelNyx/syntinel-agent/internal/sysinfo"
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
			logger.Error("Receive error: %v", err)
			return err
		}

		logger.Info("Received: command=%s payload=%s misc=%t", msg.Command, msg.Payload, len(msg.Misc) != 0)

		var result string
		switch msg.Command {
		case "heartbeat":
			result = "<3"
		case "exec":
			result, err = commands.Exec(msg.Payload)
		case "download":
			result, err = commands.DownloadFile(msg.Payload, msg.Misc)
		case "hostinfo":
			result = sysinfo.CombinedInfo()
		case "sysinfo":
			result = sysinfo.SysInfo()
		default:
			result = "unknown command"
		}

		if err != nil {
			logger.Error("Error executing command: %v", err)
			return err
		}

		err = stream.Send(&controlpb.ControlResponse{
			Uuid:   string(data.ID),
			Result: result,
			Status: "ok",
		})
		if err != nil {
			logger.Error("Send error: %v", err)
			return err
		}
	}
}
