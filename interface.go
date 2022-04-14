package bony_api

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api/v0api"
	lotusMiner "github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/node/modules/dtypes"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var (
	_ = xerrors.Errorf
)

type Net = v0api.Net
type NetStruct = v0api.NetStruct
type NetStub = v0api.NetStub

type API interface {
	ChainAPI
	StateAPI
	LotusServiceAPI
	SyncState(context.Context) (*api.SyncState, error) //perm:read
}

type ChainAPI interface {
	ChainNotify(context.Context) (<-chan []*api.HeadChange, error)                                     //perm:read
	ChainHead(context.Context) (*types.TipSet, error)                                                  //perm:read
	ChainHasObj(ctx context.Context, obj cid.Cid) (bool, error)                                        //perm:read
	ChainReadObj(ctx context.Context, obj cid.Cid) ([]byte, error)                                     //perm:read
	ChainGetGenesis(ctx context.Context) (*types.TipSet, error)                                        //perm:read
	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)                            //perm:read
	ChainGetTipSetByHeight(context.Context, abi.ChainEpoch, types.TipSetKey) (*types.TipSet, error)    //perm:read
	ChainGetTipSetAfterHeight(context.Context, abi.ChainEpoch, types.TipSetKey) (*types.TipSet, error) //perm:read
	ChainGetBlock(context.Context, cid.Cid) (*types.BlockHeader, error)                                //perm:read
	ChainGetBlockMessages(ctx context.Context, msg cid.Cid) (*api.BlockMessages, error)                //perm:read
	ChainGetParentMessages(ctx context.Context, blockCid cid.Cid) ([]api.Message, error)               //perm:read
	ChainGetParentReceipts(ctx context.Context, blockCid cid.Cid) ([]*types.MessageReceipt, error)     //perm:read
	ChainSetHead(context.Context, types.TipSetKey) error                                               //perm:admin
	ChainStatObj(ctx context.Context, obj cid.Cid, base cid.Cid) (api.ObjStat, error)                  //perm:read
}

type StateAPI interface {
	StateGetActor(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*types.Actor, error)                                                      //perm:read
	StateListActors(context.Context, types.TipSetKey) ([]address.Address, error)                                                                             //perm:read
	StateChangedActors(context.Context, cid.Cid, cid.Cid) (map[string]types.Actor, error)                                                                    //perm:read
	StateMinerPower(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*api.MinerPower, error)                                                 //perm:read
	StateMarketDeals(context.Context, types.TipSetKey) (map[string]api.MarketDeal, error)                                                                    //perm:read
	StateReadState(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*api.ActorState, error)                                                  //perm:read
	StateGetReceipt(ctx context.Context, bcid cid.Cid, tsk types.TipSetKey) (*types.MessageReceipt, error)                                                   //perm:read
	StateVMCirculatingSupplyInternal(context.Context, types.TipSetKey) (api.CirculatingSupply, error)                                                        //perm:read
	StateNetworkName(context.Context) (dtypes.NetworkName, error)                                                                                            //perm:read
	StateMinerSectorCount(context.Context, address.Address, types.TipSetKey) (api.MinerSectors, error)                                                       //perm:read
	StateMinerPartitions(ctx context.Context, m address.Address, dlIdx uint64, tsk types.TipSetKey) ([]api.Partition, error)                                 //perm:read
	StateMinerInfo(context.Context, address.Address, types.TipSetKey) (lotusMiner.MinerInfo, error)                                                          //perm:read
	StateMinerSectorAllocated(context.Context, address.Address, abi.SectorNumber, types.TipSetKey) (bool, error)                                             //perm:read
	StateSectorPreCommitInfo(context.Context, address.Address, abi.SectorNumber, types.TipSetKey) (lotusMiner.SectorPreCommitOnChainInfo, error)             //perm:read
	StateSectorGetInfo(context.Context, address.Address, abi.SectorNumber, types.TipSetKey) (*lotusMiner.SectorOnChainInfo, error)                           //perm:read
	StateSectorExpiration(context.Context, address.Address, abi.SectorNumber, types.TipSetKey) (*lotusMiner.SectorExpiration, error)                         //perm:read
	StateSectorPartition(ctx context.Context, maddr address.Address, sectorNumber abi.SectorNumber, tsk types.TipSetKey) (*lotusMiner.SectorLocation, error) //perm:read
	StateMinerDeadlines(context.Context, address.Address, types.TipSetKey) ([]api.Deadline, error)                                                           //perm:read
}

type LotusServiceAPI interface {
	MinerSectorChanges(ctx context.Context, addr address.Address, from, to abi.ChainEpoch) (*lotusMiner.SectorChanges, error) //perm:read
	MinerVestingFunds(ctx context.Context, addr address.Address, target abi.ChainEpoch) (abi.TokenAmount, error)              //perm:read
}

type AuthAPI interface {
	AuthVerify(ctx context.Context, token string) ([]auth.Permission, error) //perm:read
}
