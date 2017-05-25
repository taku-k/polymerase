package tempbackup

import (
	"io"

	log "github.com/sirupsen/logrus"
	pb "github.com/taku-k/polymerase/pkg/tempbackup/proto"
)

type TempBackupTransferService struct {
	manager *TempBackupManager
}

func NewBackupTransferService(m *TempBackupManager) *TempBackupTransferService {
	return &TempBackupTransferService{
		manager: m,
	}
}

func (s *TempBackupTransferService) TransferFullBackup(stream pb.BackupTransferService_TransferFullBackupServer) error {
	var state *TempBackupState
	for {
		content, err := stream.Recv()
		if err == io.EOF {
			if err := state.Close(); err != nil {
				return err
			}
			return stream.SendAndClose(&pb.BackupReply{
				Message: "success",
			})
		}
		if err != nil {
			return err
		}
		if state == nil {
			state, err = s.manager.OpenFullBackup(content.Db)
			log.WithField("db", content.Db).Info("Connected from db")
			log.WithField("path", state.tempDir).Info("Create temporary directory")
			if err != nil {
				return err
			}
		}
		if err := state.Append(content.Content); err != nil {
			return err
		}
	}
}

func (s *TempBackupTransferService) TransferIncBackup(stream pb.BackupTransferService_TransferIncBackupServer) error {
	var state *TempBackupState
	var content *pb.IncBackupContentStream
	var err error
	for {
		content, err = stream.Recv()
		if err == io.EOF {
			if err := state.Close(); err != nil {
				return err
			}
			return stream.SendAndClose(&pb.BackupReply{
				Message: "success",
			})
		}
		if err != nil {
			return err
		}
		if state == nil {
			state, err = s.manager.OpenIncBackup(content.Db, content.Lsn)
			if err != nil {
				return err
			}
		}
		if err := state.Append(content.Content); err != nil {
			return err
		}
	}
}
