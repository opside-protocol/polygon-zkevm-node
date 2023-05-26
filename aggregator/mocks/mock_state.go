// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	pgx "github.com/jackc/pgx/v4"
	mock "github.com/stretchr/testify/mock"

	state "github.com/0xPolygonHermez/zkevm-node/state"
)

// StateMock is an autogenerated mock type for the stateInterface type
type StateMock struct {
	mock.Mock
}

// AddGeneratedProof provides a mock function with given fields: ctx, proof, dbTx
func (_m *StateMock) AddGeneratedProof(ctx context.Context, proof *state.Proof, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, proof, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Proof, pgx.Tx) error); ok {
		r0 = rf(ctx, proof, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *StateMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckProofContainsCompleteSequences provides a mock function with given fields: ctx, proof, dbTx
func (_m *StateMock) CheckProofContainsCompleteSequences(ctx context.Context, proof *state.Proof, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, proof, dbTx)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Proof, pgx.Tx) (bool, error)); ok {
		return rf(ctx, proof, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Proof, pgx.Tx) bool); ok {
		r0 = rf(ctx, proof, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Proof, pgx.Tx) error); ok {
		r1 = rf(ctx, proof, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanupGeneratedProofs provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) CleanupGeneratedProofs(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CleanupLockedProofs provides a mock function with given fields: ctx, duration, dbTx
func (_m *StateMock) CleanupLockedProofs(ctx context.Context, duration string, dbTx pgx.Tx) (int64, error) {
	ret := _m.Called(ctx, duration, dbTx)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, pgx.Tx) (int64, error)); ok {
		return rf(ctx, duration, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, pgx.Tx) int64); ok {
		r0 = rf(ctx, duration, dbTx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, pgx.Tx) error); ok {
		r1 = rf(ctx, duration, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGeneratedProofs provides a mock function with given fields: ctx, batchNumber, batchNumberFinal, dbTx
func (_m *StateMock) DeleteGeneratedProofs(ctx context.Context, batchNumber uint64, batchNumberFinal uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, batchNumberFinal, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, batchNumberFinal, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUngeneratedProofs provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) DeleteUngeneratedProofs(ctx context.Context, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) error); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *StateMock) GetEarlyProofHashByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVerifiedBatch provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastVerifiedBatch(ctx context.Context, dbTx pgx.Tx) (*state.VerifiedBatch, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.VerifiedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.VerifiedBatch, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.VerifiedBatch); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.VerifiedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProofReadyToVerify provides a mock function with given fields: ctx, lastVerfiedBatchNumber, dbTx
func (_m *StateMock) GetProofReadyToVerify(ctx context.Context, lastVerfiedBatchNumber uint64, dbTx pgx.Tx) (*state.Proof, error) {
	ret := _m.Called(ctx, lastVerfiedBatchNumber, dbTx)

	var r0 *state.Proof
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Proof, error)); ok {
		return rf(ctx, lastVerfiedBatchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Proof); ok {
		r0 = rf(ctx, lastVerfiedBatchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Proof)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, lastVerfiedBatchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *StateMock)GetLastBlock(ctx context.Context, dbTx pgx.Tx) (*state.Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProofsToAggregate provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetProofsToAggregate(ctx context.Context, dbTx pgx.Tx) (*state.Proof, *state.Proof, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.Proof
	var r1 *state.Proof
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.Proof, *state.Proof, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.Proof); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Proof)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) *state.Proof); ok {
		r1 = rf(ctx, dbTx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*state.Proof)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, pgx.Tx) error); ok {
		r2 = rf(ctx, dbTx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

func (_m *StateMock) GetProofHashBySender(ctx context.Context, sender string, batchNumber, minCommit, lastBlockNumber uint64, dbTx pgx.Tx) (string, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (string, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) string); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *StateMock) AddProverProof(ctx context.Context, proverProof *state.ProverProof, dbTx pgx.Tx) error {
	return nil
}

func (_m *StateMock) GetProverProofByHash(ctx context.Context, hash string, batchNumberFinal uint64, dbTx pgx.Tx) (*state.ProverProof, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.ProverProof
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.ProverProof, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.ProverProof); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ProverProof)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetVirtualBatchToProve provides a mock function with given fields: ctx, lastVerfiedBatchNumber, dbTx
func (_m *StateMock) GetVirtualBatchToProve(ctx context.Context, lastVerfiedBatchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, lastVerfiedBatchNumber, dbTx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, lastVerfiedBatchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, lastVerfiedBatchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, lastVerfiedBatchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGeneratedProof provides a mock function with given fields: ctx, proof, dbTx
func (_m *StateMock) UpdateGeneratedProof(ctx context.Context, proof *state.Proof, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, proof, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Proof, pgx.Tx) error); ok {
		r0 = rf(ctx, proof, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}


func (p *StateMock) AddFinalProof(ctx context.Context, finalProof *state.FinalProof, dbTx pgx.Tx) error {
	return nil
}

func (p *StateMock) GetFinalProofByMonitoredId(ctx context.Context, monitoredId string, dbTx pgx.Tx) (*state.FinalProof, error) {
	return nil,nil
}

func (p *StateMock) GetSequence(ctx context.Context, lastVerifiedBatchNumber uint64, dbTx pgx.Tx) (state.Sequence, error) {
	return state.Sequence{},nil
}

func (p *StateMock) GetStatusDoneBlockNum(ctx context.Context, id string, dbTx pgx.Tx) (uint64, error)  {
	return 0, nil
}

func (p *StateMock) HaveProverProofByBatchNum(ctx context.Context, batchNumberFinal uint64, dbTx pgx.Tx) (bool, error) {
	return true, nil
}

type mockConstructorTestingTNewStateMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewStateMock creates a new instance of StateMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStateMock(t mockConstructorTestingTNewStateMock) *StateMock {
	mock := &StateMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
