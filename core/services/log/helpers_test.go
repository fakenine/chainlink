package log_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/store"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

type simpleLogListener struct {
	handler    func(lb log.Broadcast, err error)
	consumerID models.JobID
}

func (listener simpleLogListener) HandleLog(lb log.Broadcast, err error) {
	listener.handler(lb, err)
}
func (listener simpleLogListener) OnConnect()    {}
func (listener simpleLogListener) OnDisconnect() {}
func (listener simpleLogListener) JobID() models.JobID {
	return listener.consumerID
}
func (listener simpleLogListener) IsV2Job() bool {
	return false
}
func (listener simpleLogListener) JobIDV2() int32 {
	return 0
}

type logBroadcastRow struct {
	BlockHash   common.Hash
	BlockNumber uint64
	LogIndex    uint
	JobID       models.JobID
	JobIDV2     int32
	Consumed    bool
}

type mockListener struct {
	jobID   models.JobID
	jobIDV2 int32
}

func (l *mockListener) JobID() models.JobID            { return l.jobID }
func (l *mockListener) JobIDV2() int32                 { return l.jobIDV2 }
func (l *mockListener) IsV2Job() bool                  { return l.jobID.IsZero() }
func (l *mockListener) OnConnect()                     {}
func (l *mockListener) OnDisconnect()                  {}
func (l *mockListener) HandleLog(log.Broadcast, error) {}

func createJob(t *testing.T, store *store.Store) models.JobSpec {
	t.Helper()

	job := cltest.NewJob()
	err := store.ORM.CreateJob(&job)
	require.NoError(t, err)
	return job
}
