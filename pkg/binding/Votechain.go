// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VotechainCandidate is an auto generated low-level Go binding around an user-defined struct.
type VotechainCandidate struct {
	Name      string
	VoteCount *big.Int
	IsActive  bool
}

// VotechainKPUBranch is an auto generated low-level Go binding around an user-defined struct.
type VotechainKPUBranch struct {
	Name          string
	BranchAddress common.Address
	IsActive      bool
	Region        string
}

// VotechainVoter is an auto generated low-level Go binding around an user-defined struct.
type VotechainVoter struct {
	Nik          string
	VoterAddress common.Address
	HasVoted     bool
	Region       string
	IsRegistered bool
}

// VotechainMetaData contains all meta data concerning the Votechain contract.
var VotechainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BranchAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BranchNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCandidate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyKpuAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyKpuBranch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedVoter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VoterAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VoterNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotingNotActive\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CandidateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"CandidateStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"}],\"name\":\"KPUBranchDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"KPUBranchRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"VoteCasted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"VoterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"VotingStatusChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"}],\"name\":\"deactivateKPUBranch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structVotechain.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllKPUBranches\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"internalType\":\"structVotechain.KPUBranch[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllVoter\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"internalType\":\"structVotechain.Voter[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"}],\"name\":\"getBranchByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"internalType\":\"structVotechain.KPUBranch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"getCandidate\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structVotechain.Candidate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"getVoterByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"internalType\":\"structVotechain.Voter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"}],\"name\":\"getVoterByNIK\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"internalType\":\"structVotechain.Voter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"getVoterByRegion\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"internalType\":\"structVotechain.Voter[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kpuAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kpuBranchAddresses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"kpuBranches\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"branchAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"}],\"name\":\"registerKPUBranch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setKpuAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setVotingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"toggleCandidateActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voterAddresses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voterNIKByAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotechainABI is the input ABI used to generate the binding from.
// Deprecated: Use VotechainMetaData.ABI instead.
var VotechainABI = VotechainMetaData.ABI

// Votechain is an auto generated Go binding around an Ethereum contract.
type Votechain struct {
	VotechainCaller     // Read-only binding to the contract
	VotechainTransactor // Write-only binding to the contract
	VotechainFilterer   // Log filterer for contract events
}

// VotechainCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotechainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotechainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotechainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotechainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotechainSession struct {
	Contract     *Votechain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotechainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotechainCallerSession struct {
	Contract *VotechainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VotechainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotechainTransactorSession struct {
	Contract     *VotechainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VotechainRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotechainRaw struct {
	Contract *Votechain // Generic contract binding to access the raw methods on
}

// VotechainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotechainCallerRaw struct {
	Contract *VotechainCaller // Generic read-only contract binding to access the raw methods on
}

// VotechainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotechainTransactorRaw struct {
	Contract *VotechainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotechain creates a new instance of Votechain, bound to a specific deployed contract.
func NewVotechain(address common.Address, backend bind.ContractBackend) (*Votechain, error) {
	contract, err := bindVotechain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Votechain{VotechainCaller: VotechainCaller{contract: contract}, VotechainTransactor: VotechainTransactor{contract: contract}, VotechainFilterer: VotechainFilterer{contract: contract}}, nil
}

// NewVotechainCaller creates a new read-only instance of Votechain, bound to a specific deployed contract.
func NewVotechainCaller(address common.Address, caller bind.ContractCaller) (*VotechainCaller, error) {
	contract, err := bindVotechain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotechainCaller{contract: contract}, nil
}

// NewVotechainTransactor creates a new write-only instance of Votechain, bound to a specific deployed contract.
func NewVotechainTransactor(address common.Address, transactor bind.ContractTransactor) (*VotechainTransactor, error) {
	contract, err := bindVotechain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotechainTransactor{contract: contract}, nil
}

// NewVotechainFilterer creates a new log filterer instance of Votechain, bound to a specific deployed contract.
func NewVotechainFilterer(address common.Address, filterer bind.ContractFilterer) (*VotechainFilterer, error) {
	contract, err := bindVotechain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotechainFilterer{contract: contract}, nil
}

// bindVotechain binds a generic wrapper to an already deployed contract.
func bindVotechain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotechainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Votechain *VotechainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Votechain.Contract.VotechainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Votechain *VotechainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votechain.Contract.VotechainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Votechain *VotechainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Votechain.Contract.VotechainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Votechain *VotechainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Votechain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Votechain *VotechainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votechain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Votechain *VotechainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Votechain.Contract.contract.Transact(opts, method, params...)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Votechain *VotechainCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "candidateCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Votechain *VotechainSession) CandidateCount() (*big.Int, error) {
	return _Votechain.Contract.CandidateCount(&_Votechain.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() view returns(uint256)
func (_Votechain *VotechainCallerSession) CandidateCount() (*big.Int, error) {
	return _Votechain.Contract.CandidateCount(&_Votechain.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount, bool isActive)
func (_Votechain *VotechainCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
	IsActive  bool
}, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Name      string
		VoteCount *big.Int
		IsActive  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount, bool isActive)
func (_Votechain *VotechainSession) Candidates(arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
	IsActive  bool
}, error) {
	return _Votechain.Contract.Candidates(&_Votechain.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount, bool isActive)
func (_Votechain *VotechainCallerSession) Candidates(arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
	IsActive  bool
}, error) {
	return _Votechain.Contract.Candidates(&_Votechain.CallOpts, arg0)
}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((string,uint256,bool)[])
func (_Votechain *VotechainCaller) GetAllCandidates(opts *bind.CallOpts) ([]VotechainCandidate, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getAllCandidates")

	if err != nil {
		return *new([]VotechainCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotechainCandidate)).(*[]VotechainCandidate)

	return out0, err

}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((string,uint256,bool)[])
func (_Votechain *VotechainSession) GetAllCandidates() ([]VotechainCandidate, error) {
	return _Votechain.Contract.GetAllCandidates(&_Votechain.CallOpts)
}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((string,uint256,bool)[])
func (_Votechain *VotechainCallerSession) GetAllCandidates() ([]VotechainCandidate, error) {
	return _Votechain.Contract.GetAllCandidates(&_Votechain.CallOpts)
}

// GetAllKPUBranches is a free data retrieval call binding the contract method 0xb0c9b9a0.
//
// Solidity: function getAllKPUBranches() view returns((string,address,bool,string)[])
func (_Votechain *VotechainCaller) GetAllKPUBranches(opts *bind.CallOpts) ([]VotechainKPUBranch, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getAllKPUBranches")

	if err != nil {
		return *new([]VotechainKPUBranch), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotechainKPUBranch)).(*[]VotechainKPUBranch)

	return out0, err

}

// GetAllKPUBranches is a free data retrieval call binding the contract method 0xb0c9b9a0.
//
// Solidity: function getAllKPUBranches() view returns((string,address,bool,string)[])
func (_Votechain *VotechainSession) GetAllKPUBranches() ([]VotechainKPUBranch, error) {
	return _Votechain.Contract.GetAllKPUBranches(&_Votechain.CallOpts)
}

// GetAllKPUBranches is a free data retrieval call binding the contract method 0xb0c9b9a0.
//
// Solidity: function getAllKPUBranches() view returns((string,address,bool,string)[])
func (_Votechain *VotechainCallerSession) GetAllKPUBranches() ([]VotechainKPUBranch, error) {
	return _Votechain.Contract.GetAllKPUBranches(&_Votechain.CallOpts)
}

// GetAllVoter is a free data retrieval call binding the contract method 0xf44f4e14.
//
// Solidity: function getAllVoter() view returns((string,address,bool,string,bool)[])
func (_Votechain *VotechainCaller) GetAllVoter(opts *bind.CallOpts) ([]VotechainVoter, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getAllVoter")

	if err != nil {
		return *new([]VotechainVoter), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotechainVoter)).(*[]VotechainVoter)

	return out0, err

}

// GetAllVoter is a free data retrieval call binding the contract method 0xf44f4e14.
//
// Solidity: function getAllVoter() view returns((string,address,bool,string,bool)[])
func (_Votechain *VotechainSession) GetAllVoter() ([]VotechainVoter, error) {
	return _Votechain.Contract.GetAllVoter(&_Votechain.CallOpts)
}

// GetAllVoter is a free data retrieval call binding the contract method 0xf44f4e14.
//
// Solidity: function getAllVoter() view returns((string,address,bool,string,bool)[])
func (_Votechain *VotechainCallerSession) GetAllVoter() ([]VotechainVoter, error) {
	return _Votechain.Contract.GetAllVoter(&_Votechain.CallOpts)
}

// GetBranchByAddress is a free data retrieval call binding the contract method 0xefe3b01b.
//
// Solidity: function getBranchByAddress(address branchAddress) view returns((string,address,bool,string))
func (_Votechain *VotechainCaller) GetBranchByAddress(opts *bind.CallOpts, branchAddress common.Address) (VotechainKPUBranch, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getBranchByAddress", branchAddress)

	if err != nil {
		return *new(VotechainKPUBranch), err
	}

	out0 := *abi.ConvertType(out[0], new(VotechainKPUBranch)).(*VotechainKPUBranch)

	return out0, err

}

// GetBranchByAddress is a free data retrieval call binding the contract method 0xefe3b01b.
//
// Solidity: function getBranchByAddress(address branchAddress) view returns((string,address,bool,string))
func (_Votechain *VotechainSession) GetBranchByAddress(branchAddress common.Address) (VotechainKPUBranch, error) {
	return _Votechain.Contract.GetBranchByAddress(&_Votechain.CallOpts, branchAddress)
}

// GetBranchByAddress is a free data retrieval call binding the contract method 0xefe3b01b.
//
// Solidity: function getBranchByAddress(address branchAddress) view returns((string,address,bool,string))
func (_Votechain *VotechainCallerSession) GetBranchByAddress(branchAddress common.Address) (VotechainKPUBranch, error) {
	return _Votechain.Contract.GetBranchByAddress(&_Votechain.CallOpts, branchAddress)
}

// GetCandidate is a free data retrieval call binding the contract method 0x35b8e820.
//
// Solidity: function getCandidate(uint256 candidateId) view returns((string,uint256,bool))
func (_Votechain *VotechainCaller) GetCandidate(opts *bind.CallOpts, candidateId *big.Int) (VotechainCandidate, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getCandidate", candidateId)

	if err != nil {
		return *new(VotechainCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new(VotechainCandidate)).(*VotechainCandidate)

	return out0, err

}

// GetCandidate is a free data retrieval call binding the contract method 0x35b8e820.
//
// Solidity: function getCandidate(uint256 candidateId) view returns((string,uint256,bool))
func (_Votechain *VotechainSession) GetCandidate(candidateId *big.Int) (VotechainCandidate, error) {
	return _Votechain.Contract.GetCandidate(&_Votechain.CallOpts, candidateId)
}

// GetCandidate is a free data retrieval call binding the contract method 0x35b8e820.
//
// Solidity: function getCandidate(uint256 candidateId) view returns((string,uint256,bool))
func (_Votechain *VotechainCallerSession) GetCandidate(candidateId *big.Int) (VotechainCandidate, error) {
	return _Votechain.Contract.GetCandidate(&_Votechain.CallOpts, candidateId)
}

// GetVoterByAddress is a free data retrieval call binding the contract method 0x4bdd7585.
//
// Solidity: function getVoterByAddress(address voterAddress) view returns((string,address,bool,string,bool))
func (_Votechain *VotechainCaller) GetVoterByAddress(opts *bind.CallOpts, voterAddress common.Address) (VotechainVoter, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getVoterByAddress", voterAddress)

	if err != nil {
		return *new(VotechainVoter), err
	}

	out0 := *abi.ConvertType(out[0], new(VotechainVoter)).(*VotechainVoter)

	return out0, err

}

// GetVoterByAddress is a free data retrieval call binding the contract method 0x4bdd7585.
//
// Solidity: function getVoterByAddress(address voterAddress) view returns((string,address,bool,string,bool))
func (_Votechain *VotechainSession) GetVoterByAddress(voterAddress common.Address) (VotechainVoter, error) {
	return _Votechain.Contract.GetVoterByAddress(&_Votechain.CallOpts, voterAddress)
}

// GetVoterByAddress is a free data retrieval call binding the contract method 0x4bdd7585.
//
// Solidity: function getVoterByAddress(address voterAddress) view returns((string,address,bool,string,bool))
func (_Votechain *VotechainCallerSession) GetVoterByAddress(voterAddress common.Address) (VotechainVoter, error) {
	return _Votechain.Contract.GetVoterByAddress(&_Votechain.CallOpts, voterAddress)
}

// GetVoterByNIK is a free data retrieval call binding the contract method 0x2a70ecca.
//
// Solidity: function getVoterByNIK(string nik) view returns((string,address,bool,string,bool))
func (_Votechain *VotechainCaller) GetVoterByNIK(opts *bind.CallOpts, nik string) (VotechainVoter, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getVoterByNIK", nik)

	if err != nil {
		return *new(VotechainVoter), err
	}

	out0 := *abi.ConvertType(out[0], new(VotechainVoter)).(*VotechainVoter)

	return out0, err

}

// GetVoterByNIK is a free data retrieval call binding the contract method 0x2a70ecca.
//
// Solidity: function getVoterByNIK(string nik) view returns((string,address,bool,string,bool))
func (_Votechain *VotechainSession) GetVoterByNIK(nik string) (VotechainVoter, error) {
	return _Votechain.Contract.GetVoterByNIK(&_Votechain.CallOpts, nik)
}

// GetVoterByNIK is a free data retrieval call binding the contract method 0x2a70ecca.
//
// Solidity: function getVoterByNIK(string nik) view returns((string,address,bool,string,bool))
func (_Votechain *VotechainCallerSession) GetVoterByNIK(nik string) (VotechainVoter, error) {
	return _Votechain.Contract.GetVoterByNIK(&_Votechain.CallOpts, nik)
}

// GetVoterByRegion is a free data retrieval call binding the contract method 0xe0d5343b.
//
// Solidity: function getVoterByRegion(string region) view returns((string,address,bool,string,bool)[])
func (_Votechain *VotechainCaller) GetVoterByRegion(opts *bind.CallOpts, region string) ([]VotechainVoter, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "getVoterByRegion", region)

	if err != nil {
		return *new([]VotechainVoter), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotechainVoter)).(*[]VotechainVoter)

	return out0, err

}

// GetVoterByRegion is a free data retrieval call binding the contract method 0xe0d5343b.
//
// Solidity: function getVoterByRegion(string region) view returns((string,address,bool,string,bool)[])
func (_Votechain *VotechainSession) GetVoterByRegion(region string) ([]VotechainVoter, error) {
	return _Votechain.Contract.GetVoterByRegion(&_Votechain.CallOpts, region)
}

// GetVoterByRegion is a free data retrieval call binding the contract method 0xe0d5343b.
//
// Solidity: function getVoterByRegion(string region) view returns((string,address,bool,string,bool)[])
func (_Votechain *VotechainCallerSession) GetVoterByRegion(region string) ([]VotechainVoter, error) {
	return _Votechain.Contract.GetVoterByRegion(&_Votechain.CallOpts, region)
}

// KpuAdmin is a free data retrieval call binding the contract method 0xfb4ab164.
//
// Solidity: function kpuAdmin() view returns(address)
func (_Votechain *VotechainCaller) KpuAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "kpuAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KpuAdmin is a free data retrieval call binding the contract method 0xfb4ab164.
//
// Solidity: function kpuAdmin() view returns(address)
func (_Votechain *VotechainSession) KpuAdmin() (common.Address, error) {
	return _Votechain.Contract.KpuAdmin(&_Votechain.CallOpts)
}

// KpuAdmin is a free data retrieval call binding the contract method 0xfb4ab164.
//
// Solidity: function kpuAdmin() view returns(address)
func (_Votechain *VotechainCallerSession) KpuAdmin() (common.Address, error) {
	return _Votechain.Contract.KpuAdmin(&_Votechain.CallOpts)
}

// KpuBranchAddresses is a free data retrieval call binding the contract method 0x04abffb5.
//
// Solidity: function kpuBranchAddresses(uint256 ) view returns(string name, address branchAddress, bool isActive, string region)
func (_Votechain *VotechainCaller) KpuBranchAddresses(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name          string
	BranchAddress common.Address
	IsActive      bool
	Region        string
}, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "kpuBranchAddresses", arg0)

	outstruct := new(struct {
		Name          string
		BranchAddress common.Address
		IsActive      bool
		Region        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.BranchAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Region = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// KpuBranchAddresses is a free data retrieval call binding the contract method 0x04abffb5.
//
// Solidity: function kpuBranchAddresses(uint256 ) view returns(string name, address branchAddress, bool isActive, string region)
func (_Votechain *VotechainSession) KpuBranchAddresses(arg0 *big.Int) (struct {
	Name          string
	BranchAddress common.Address
	IsActive      bool
	Region        string
}, error) {
	return _Votechain.Contract.KpuBranchAddresses(&_Votechain.CallOpts, arg0)
}

// KpuBranchAddresses is a free data retrieval call binding the contract method 0x04abffb5.
//
// Solidity: function kpuBranchAddresses(uint256 ) view returns(string name, address branchAddress, bool isActive, string region)
func (_Votechain *VotechainCallerSession) KpuBranchAddresses(arg0 *big.Int) (struct {
	Name          string
	BranchAddress common.Address
	IsActive      bool
	Region        string
}, error) {
	return _Votechain.Contract.KpuBranchAddresses(&_Votechain.CallOpts, arg0)
}

// KpuBranches is a free data retrieval call binding the contract method 0xab165a3e.
//
// Solidity: function kpuBranches(address ) view returns(string name, address branchAddress, bool isActive, string region)
func (_Votechain *VotechainCaller) KpuBranches(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name          string
	BranchAddress common.Address
	IsActive      bool
	Region        string
}, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "kpuBranches", arg0)

	outstruct := new(struct {
		Name          string
		BranchAddress common.Address
		IsActive      bool
		Region        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.BranchAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Region = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// KpuBranches is a free data retrieval call binding the contract method 0xab165a3e.
//
// Solidity: function kpuBranches(address ) view returns(string name, address branchAddress, bool isActive, string region)
func (_Votechain *VotechainSession) KpuBranches(arg0 common.Address) (struct {
	Name          string
	BranchAddress common.Address
	IsActive      bool
	Region        string
}, error) {
	return _Votechain.Contract.KpuBranches(&_Votechain.CallOpts, arg0)
}

// KpuBranches is a free data retrieval call binding the contract method 0xab165a3e.
//
// Solidity: function kpuBranches(address ) view returns(string name, address branchAddress, bool isActive, string region)
func (_Votechain *VotechainCallerSession) KpuBranches(arg0 common.Address) (struct {
	Name          string
	BranchAddress common.Address
	IsActive      bool
	Region        string
}, error) {
	return _Votechain.Contract.KpuBranches(&_Votechain.CallOpts, arg0)
}

// VoterAddresses is a free data retrieval call binding the contract method 0xdd0e2373.
//
// Solidity: function voterAddresses(uint256 ) view returns(string nik, address voterAddress, bool hasVoted, string region, bool isRegistered)
func (_Votechain *VotechainCaller) VoterAddresses(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Nik          string
	VoterAddress common.Address
	HasVoted     bool
	Region       string
	IsRegistered bool
}, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "voterAddresses", arg0)

	outstruct := new(struct {
		Nik          string
		VoterAddress common.Address
		HasVoted     bool
		Region       string
		IsRegistered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nik = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoterAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.HasVoted = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Region = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsRegistered = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// VoterAddresses is a free data retrieval call binding the contract method 0xdd0e2373.
//
// Solidity: function voterAddresses(uint256 ) view returns(string nik, address voterAddress, bool hasVoted, string region, bool isRegistered)
func (_Votechain *VotechainSession) VoterAddresses(arg0 *big.Int) (struct {
	Nik          string
	VoterAddress common.Address
	HasVoted     bool
	Region       string
	IsRegistered bool
}, error) {
	return _Votechain.Contract.VoterAddresses(&_Votechain.CallOpts, arg0)
}

// VoterAddresses is a free data retrieval call binding the contract method 0xdd0e2373.
//
// Solidity: function voterAddresses(uint256 ) view returns(string nik, address voterAddress, bool hasVoted, string region, bool isRegistered)
func (_Votechain *VotechainCallerSession) VoterAddresses(arg0 *big.Int) (struct {
	Nik          string
	VoterAddress common.Address
	HasVoted     bool
	Region       string
	IsRegistered bool
}, error) {
	return _Votechain.Contract.VoterAddresses(&_Votechain.CallOpts, arg0)
}

// VoterNIKByAddress is a free data retrieval call binding the contract method 0xf0416e5f.
//
// Solidity: function voterNIKByAddress(address ) view returns(string)
func (_Votechain *VotechainCaller) VoterNIKByAddress(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "voterNIKByAddress", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VoterNIKByAddress is a free data retrieval call binding the contract method 0xf0416e5f.
//
// Solidity: function voterNIKByAddress(address ) view returns(string)
func (_Votechain *VotechainSession) VoterNIKByAddress(arg0 common.Address) (string, error) {
	return _Votechain.Contract.VoterNIKByAddress(&_Votechain.CallOpts, arg0)
}

// VoterNIKByAddress is a free data retrieval call binding the contract method 0xf0416e5f.
//
// Solidity: function voterNIKByAddress(address ) view returns(string)
func (_Votechain *VotechainCallerSession) VoterNIKByAddress(arg0 common.Address) (string, error) {
	return _Votechain.Contract.VoterNIKByAddress(&_Votechain.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0x53fa2e64.
//
// Solidity: function voters(string ) view returns(string nik, address voterAddress, bool hasVoted, string region, bool isRegistered)
func (_Votechain *VotechainCaller) Voters(opts *bind.CallOpts, arg0 string) (struct {
	Nik          string
	VoterAddress common.Address
	HasVoted     bool
	Region       string
	IsRegistered bool
}, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Nik          string
		VoterAddress common.Address
		HasVoted     bool
		Region       string
		IsRegistered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nik = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoterAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.HasVoted = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Region = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsRegistered = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0x53fa2e64.
//
// Solidity: function voters(string ) view returns(string nik, address voterAddress, bool hasVoted, string region, bool isRegistered)
func (_Votechain *VotechainSession) Voters(arg0 string) (struct {
	Nik          string
	VoterAddress common.Address
	HasVoted     bool
	Region       string
	IsRegistered bool
}, error) {
	return _Votechain.Contract.Voters(&_Votechain.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0x53fa2e64.
//
// Solidity: function voters(string ) view returns(string nik, address voterAddress, bool hasVoted, string region, bool isRegistered)
func (_Votechain *VotechainCallerSession) Voters(arg0 string) (struct {
	Nik          string
	VoterAddress common.Address
	HasVoted     bool
	Region       string
	IsRegistered bool
}, error) {
	return _Votechain.Contract.Voters(&_Votechain.CallOpts, arg0)
}

// VotingActive is a free data retrieval call binding the contract method 0x408e2727.
//
// Solidity: function votingActive() view returns(bool)
func (_Votechain *VotechainCaller) VotingActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Votechain.contract.Call(opts, &out, "votingActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VotingActive is a free data retrieval call binding the contract method 0x408e2727.
//
// Solidity: function votingActive() view returns(bool)
func (_Votechain *VotechainSession) VotingActive() (bool, error) {
	return _Votechain.Contract.VotingActive(&_Votechain.CallOpts)
}

// VotingActive is a free data retrieval call binding the contract method 0x408e2727.
//
// Solidity: function votingActive() view returns(bool)
func (_Votechain *VotechainCallerSession) VotingActive() (bool, error) {
	return _Votechain.Contract.VotingActive(&_Votechain.CallOpts)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string name) returns()
func (_Votechain *VotechainTransactor) AddCandidate(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "addCandidate", name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string name) returns()
func (_Votechain *VotechainSession) AddCandidate(name string) (*types.Transaction, error) {
	return _Votechain.Contract.AddCandidate(&_Votechain.TransactOpts, name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string name) returns()
func (_Votechain *VotechainTransactorSession) AddCandidate(name string) (*types.Transaction, error) {
	return _Votechain.Contract.AddCandidate(&_Votechain.TransactOpts, name)
}

// DeactivateKPUBranch is a paid mutator transaction binding the contract method 0x4c361435.
//
// Solidity: function deactivateKPUBranch(address branchAddress) returns()
func (_Votechain *VotechainTransactor) DeactivateKPUBranch(opts *bind.TransactOpts, branchAddress common.Address) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "deactivateKPUBranch", branchAddress)
}

// DeactivateKPUBranch is a paid mutator transaction binding the contract method 0x4c361435.
//
// Solidity: function deactivateKPUBranch(address branchAddress) returns()
func (_Votechain *VotechainSession) DeactivateKPUBranch(branchAddress common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.DeactivateKPUBranch(&_Votechain.TransactOpts, branchAddress)
}

// DeactivateKPUBranch is a paid mutator transaction binding the contract method 0x4c361435.
//
// Solidity: function deactivateKPUBranch(address branchAddress) returns()
func (_Votechain *VotechainTransactorSession) DeactivateKPUBranch(branchAddress common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.DeactivateKPUBranch(&_Votechain.TransactOpts, branchAddress)
}

// RegisterKPUBranch is a paid mutator transaction binding the contract method 0x027d8514.
//
// Solidity: function registerKPUBranch(address branchAddress, string name, string region) returns()
func (_Votechain *VotechainTransactor) RegisterKPUBranch(opts *bind.TransactOpts, branchAddress common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "registerKPUBranch", branchAddress, name, region)
}

// RegisterKPUBranch is a paid mutator transaction binding the contract method 0x027d8514.
//
// Solidity: function registerKPUBranch(address branchAddress, string name, string region) returns()
func (_Votechain *VotechainSession) RegisterKPUBranch(branchAddress common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterKPUBranch(&_Votechain.TransactOpts, branchAddress, name, region)
}

// RegisterKPUBranch is a paid mutator transaction binding the contract method 0x027d8514.
//
// Solidity: function registerKPUBranch(address branchAddress, string name, string region) returns()
func (_Votechain *VotechainTransactorSession) RegisterKPUBranch(branchAddress common.Address, name string, region string) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterKPUBranch(&_Votechain.TransactOpts, branchAddress, name, region)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x4a075de2.
//
// Solidity: function registerVoter(string nik, address voterAddress) returns()
func (_Votechain *VotechainTransactor) RegisterVoter(opts *bind.TransactOpts, nik string, voterAddress common.Address) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "registerVoter", nik, voterAddress)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x4a075de2.
//
// Solidity: function registerVoter(string nik, address voterAddress) returns()
func (_Votechain *VotechainSession) RegisterVoter(nik string, voterAddress common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterVoter(&_Votechain.TransactOpts, nik, voterAddress)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x4a075de2.
//
// Solidity: function registerVoter(string nik, address voterAddress) returns()
func (_Votechain *VotechainTransactorSession) RegisterVoter(nik string, voterAddress common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.RegisterVoter(&_Votechain.TransactOpts, nik, voterAddress)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_Votechain *VotechainTransactor) SetKpuAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "setKpuAdmin", newAdmin)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_Votechain *VotechainSession) SetKpuAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.SetKpuAdmin(&_Votechain.TransactOpts, newAdmin)
}

// SetKpuAdmin is a paid mutator transaction binding the contract method 0x9df86dc1.
//
// Solidity: function setKpuAdmin(address newAdmin) returns()
func (_Votechain *VotechainTransactorSession) SetKpuAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Votechain.Contract.SetKpuAdmin(&_Votechain.TransactOpts, newAdmin)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_Votechain *VotechainTransactor) SetVotingStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "setVotingStatus", status)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_Votechain *VotechainSession) SetVotingStatus(status bool) (*types.Transaction, error) {
	return _Votechain.Contract.SetVotingStatus(&_Votechain.TransactOpts, status)
}

// SetVotingStatus is a paid mutator transaction binding the contract method 0x7478c9fe.
//
// Solidity: function setVotingStatus(bool status) returns()
func (_Votechain *VotechainTransactorSession) SetVotingStatus(status bool) (*types.Transaction, error) {
	return _Votechain.Contract.SetVotingStatus(&_Votechain.TransactOpts, status)
}

// ToggleCandidateActive is a paid mutator transaction binding the contract method 0xb07bab99.
//
// Solidity: function toggleCandidateActive(uint256 candidateId) returns()
func (_Votechain *VotechainTransactor) ToggleCandidateActive(opts *bind.TransactOpts, candidateId *big.Int) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "toggleCandidateActive", candidateId)
}

// ToggleCandidateActive is a paid mutator transaction binding the contract method 0xb07bab99.
//
// Solidity: function toggleCandidateActive(uint256 candidateId) returns()
func (_Votechain *VotechainSession) ToggleCandidateActive(candidateId *big.Int) (*types.Transaction, error) {
	return _Votechain.Contract.ToggleCandidateActive(&_Votechain.TransactOpts, candidateId)
}

// ToggleCandidateActive is a paid mutator transaction binding the contract method 0xb07bab99.
//
// Solidity: function toggleCandidateActive(uint256 candidateId) returns()
func (_Votechain *VotechainTransactorSession) ToggleCandidateActive(candidateId *big.Int) (*types.Transaction, error) {
	return _Votechain.Contract.ToggleCandidateActive(&_Votechain.TransactOpts, candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 candidateId) returns()
func (_Votechain *VotechainTransactor) Vote(opts *bind.TransactOpts, candidateId *big.Int) (*types.Transaction, error) {
	return _Votechain.contract.Transact(opts, "vote", candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 candidateId) returns()
func (_Votechain *VotechainSession) Vote(candidateId *big.Int) (*types.Transaction, error) {
	return _Votechain.Contract.Vote(&_Votechain.TransactOpts, candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 candidateId) returns()
func (_Votechain *VotechainTransactorSession) Vote(candidateId *big.Int) (*types.Transaction, error) {
	return _Votechain.Contract.Vote(&_Votechain.TransactOpts, candidateId)
}

// VotechainCandidateAddedIterator is returned from FilterCandidateAdded and is used to iterate over the raw logs and unpacked data for CandidateAdded events raised by the Votechain contract.
type VotechainCandidateAddedIterator struct {
	Event *VotechainCandidateAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotechainCandidateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainCandidateAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotechainCandidateAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotechainCandidateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainCandidateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainCandidateAdded represents a CandidateAdded event raised by the Votechain contract.
type VotechainCandidateAdded struct {
	CandidateId *big.Int
	Name        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 indexed candidateId, string name)
func (_Votechain *VotechainFilterer) FilterCandidateAdded(opts *bind.FilterOpts, candidateId []*big.Int) (*VotechainCandidateAddedIterator, error) {

	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "CandidateAdded", candidateIdRule)
	if err != nil {
		return nil, err
	}
	return &VotechainCandidateAddedIterator{contract: _Votechain.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 indexed candidateId, string name)
func (_Votechain *VotechainFilterer) WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *VotechainCandidateAdded, candidateId []*big.Int) (event.Subscription, error) {

	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "CandidateAdded", candidateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainCandidateAdded)
				if err := _Votechain.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCandidateAdded is a log parse operation binding the contract event 0xe83b2a43e7e82d975c8a0a6d2f045153c869e111136a34d1889ab7b598e396a3.
//
// Solidity: event CandidateAdded(uint256 indexed candidateId, string name)
func (_Votechain *VotechainFilterer) ParseCandidateAdded(log types.Log) (*VotechainCandidateAdded, error) {
	event := new(VotechainCandidateAdded)
	if err := _Votechain.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainCandidateStatusChangedIterator is returned from FilterCandidateStatusChanged and is used to iterate over the raw logs and unpacked data for CandidateStatusChanged events raised by the Votechain contract.
type VotechainCandidateStatusChangedIterator struct {
	Event *VotechainCandidateStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotechainCandidateStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainCandidateStatusChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotechainCandidateStatusChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotechainCandidateStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainCandidateStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainCandidateStatusChanged represents a CandidateStatusChanged event raised by the Votechain contract.
type VotechainCandidateStatusChanged struct {
	CandidateId *big.Int
	IsActive    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCandidateStatusChanged is a free log retrieval operation binding the contract event 0x07a9fd7fc1368910c1b62e58319c10b4ba034e9e4d58aa2ef36e9283fc2c5f97.
//
// Solidity: event CandidateStatusChanged(uint256 indexed candidateId, bool isActive)
func (_Votechain *VotechainFilterer) FilterCandidateStatusChanged(opts *bind.FilterOpts, candidateId []*big.Int) (*VotechainCandidateStatusChangedIterator, error) {

	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "CandidateStatusChanged", candidateIdRule)
	if err != nil {
		return nil, err
	}
	return &VotechainCandidateStatusChangedIterator{contract: _Votechain.contract, event: "CandidateStatusChanged", logs: logs, sub: sub}, nil
}

// WatchCandidateStatusChanged is a free log subscription operation binding the contract event 0x07a9fd7fc1368910c1b62e58319c10b4ba034e9e4d58aa2ef36e9283fc2c5f97.
//
// Solidity: event CandidateStatusChanged(uint256 indexed candidateId, bool isActive)
func (_Votechain *VotechainFilterer) WatchCandidateStatusChanged(opts *bind.WatchOpts, sink chan<- *VotechainCandidateStatusChanged, candidateId []*big.Int) (event.Subscription, error) {

	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "CandidateStatusChanged", candidateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainCandidateStatusChanged)
				if err := _Votechain.contract.UnpackLog(event, "CandidateStatusChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCandidateStatusChanged is a log parse operation binding the contract event 0x07a9fd7fc1368910c1b62e58319c10b4ba034e9e4d58aa2ef36e9283fc2c5f97.
//
// Solidity: event CandidateStatusChanged(uint256 indexed candidateId, bool isActive)
func (_Votechain *VotechainFilterer) ParseCandidateStatusChanged(log types.Log) (*VotechainCandidateStatusChanged, error) {
	event := new(VotechainCandidateStatusChanged)
	if err := _Votechain.contract.UnpackLog(event, "CandidateStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainKPUBranchDeactivatedIterator is returned from FilterKPUBranchDeactivated and is used to iterate over the raw logs and unpacked data for KPUBranchDeactivated events raised by the Votechain contract.
type VotechainKPUBranchDeactivatedIterator struct {
	Event *VotechainKPUBranchDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotechainKPUBranchDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainKPUBranchDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotechainKPUBranchDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotechainKPUBranchDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainKPUBranchDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainKPUBranchDeactivated represents a KPUBranchDeactivated event raised by the Votechain contract.
type VotechainKPUBranchDeactivated struct {
	BranchAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterKPUBranchDeactivated is a free log retrieval operation binding the contract event 0x3954ed3404fab00c2419751b1587a1f6655215f0dc812ab5ed515d4c2677b2d3.
//
// Solidity: event KPUBranchDeactivated(address indexed branchAddress)
func (_Votechain *VotechainFilterer) FilterKPUBranchDeactivated(opts *bind.FilterOpts, branchAddress []common.Address) (*VotechainKPUBranchDeactivatedIterator, error) {

	var branchAddressRule []interface{}
	for _, branchAddressItem := range branchAddress {
		branchAddressRule = append(branchAddressRule, branchAddressItem)
	}

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "KPUBranchDeactivated", branchAddressRule)
	if err != nil {
		return nil, err
	}
	return &VotechainKPUBranchDeactivatedIterator{contract: _Votechain.contract, event: "KPUBranchDeactivated", logs: logs, sub: sub}, nil
}

// WatchKPUBranchDeactivated is a free log subscription operation binding the contract event 0x3954ed3404fab00c2419751b1587a1f6655215f0dc812ab5ed515d4c2677b2d3.
//
// Solidity: event KPUBranchDeactivated(address indexed branchAddress)
func (_Votechain *VotechainFilterer) WatchKPUBranchDeactivated(opts *bind.WatchOpts, sink chan<- *VotechainKPUBranchDeactivated, branchAddress []common.Address) (event.Subscription, error) {

	var branchAddressRule []interface{}
	for _, branchAddressItem := range branchAddress {
		branchAddressRule = append(branchAddressRule, branchAddressItem)
	}

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "KPUBranchDeactivated", branchAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainKPUBranchDeactivated)
				if err := _Votechain.contract.UnpackLog(event, "KPUBranchDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKPUBranchDeactivated is a log parse operation binding the contract event 0x3954ed3404fab00c2419751b1587a1f6655215f0dc812ab5ed515d4c2677b2d3.
//
// Solidity: event KPUBranchDeactivated(address indexed branchAddress)
func (_Votechain *VotechainFilterer) ParseKPUBranchDeactivated(log types.Log) (*VotechainKPUBranchDeactivated, error) {
	event := new(VotechainKPUBranchDeactivated)
	if err := _Votechain.contract.UnpackLog(event, "KPUBranchDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainKPUBranchRegisteredIterator is returned from FilterKPUBranchRegistered and is used to iterate over the raw logs and unpacked data for KPUBranchRegistered events raised by the Votechain contract.
type VotechainKPUBranchRegisteredIterator struct {
	Event *VotechainKPUBranchRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotechainKPUBranchRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainKPUBranchRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotechainKPUBranchRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotechainKPUBranchRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainKPUBranchRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainKPUBranchRegistered represents a KPUBranchRegistered event raised by the Votechain contract.
type VotechainKPUBranchRegistered struct {
	BranchAddress common.Address
	Name          string
	Region        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterKPUBranchRegistered is a free log retrieval operation binding the contract event 0xc4e2c95246cc050bddc27763a59824d38df0df18e23f19099623d7e1618790f6.
//
// Solidity: event KPUBranchRegistered(address indexed branchAddress, string name, string region)
func (_Votechain *VotechainFilterer) FilterKPUBranchRegistered(opts *bind.FilterOpts, branchAddress []common.Address) (*VotechainKPUBranchRegisteredIterator, error) {

	var branchAddressRule []interface{}
	for _, branchAddressItem := range branchAddress {
		branchAddressRule = append(branchAddressRule, branchAddressItem)
	}

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "KPUBranchRegistered", branchAddressRule)
	if err != nil {
		return nil, err
	}
	return &VotechainKPUBranchRegisteredIterator{contract: _Votechain.contract, event: "KPUBranchRegistered", logs: logs, sub: sub}, nil
}

// WatchKPUBranchRegistered is a free log subscription operation binding the contract event 0xc4e2c95246cc050bddc27763a59824d38df0df18e23f19099623d7e1618790f6.
//
// Solidity: event KPUBranchRegistered(address indexed branchAddress, string name, string region)
func (_Votechain *VotechainFilterer) WatchKPUBranchRegistered(opts *bind.WatchOpts, sink chan<- *VotechainKPUBranchRegistered, branchAddress []common.Address) (event.Subscription, error) {

	var branchAddressRule []interface{}
	for _, branchAddressItem := range branchAddress {
		branchAddressRule = append(branchAddressRule, branchAddressItem)
	}

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "KPUBranchRegistered", branchAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainKPUBranchRegistered)
				if err := _Votechain.contract.UnpackLog(event, "KPUBranchRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKPUBranchRegistered is a log parse operation binding the contract event 0xc4e2c95246cc050bddc27763a59824d38df0df18e23f19099623d7e1618790f6.
//
// Solidity: event KPUBranchRegistered(address indexed branchAddress, string name, string region)
func (_Votechain *VotechainFilterer) ParseKPUBranchRegistered(log types.Log) (*VotechainKPUBranchRegistered, error) {
	event := new(VotechainKPUBranchRegistered)
	if err := _Votechain.contract.UnpackLog(event, "KPUBranchRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainVoteCastedIterator is returned from FilterVoteCasted and is used to iterate over the raw logs and unpacked data for VoteCasted events raised by the Votechain contract.
type VotechainVoteCastedIterator struct {
	Event *VotechainVoteCasted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotechainVoteCastedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainVoteCasted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotechainVoteCasted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotechainVoteCastedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainVoteCastedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainVoteCasted represents a VoteCasted event raised by the Votechain contract.
type VotechainVoteCasted struct {
	Nik         common.Hash
	CandidateId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVoteCasted is a free log retrieval operation binding the contract event 0xfef5593509c07a56349ed6186376be2d0864f40445c320bd3ea7f81c5e1de00d.
//
// Solidity: event VoteCasted(string indexed nik, uint256 indexed candidateId)
func (_Votechain *VotechainFilterer) FilterVoteCasted(opts *bind.FilterOpts, nik []string, candidateId []*big.Int) (*VotechainVoteCastedIterator, error) {

	var nikRule []interface{}
	for _, nikItem := range nik {
		nikRule = append(nikRule, nikItem)
	}
	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "VoteCasted", nikRule, candidateIdRule)
	if err != nil {
		return nil, err
	}
	return &VotechainVoteCastedIterator{contract: _Votechain.contract, event: "VoteCasted", logs: logs, sub: sub}, nil
}

// WatchVoteCasted is a free log subscription operation binding the contract event 0xfef5593509c07a56349ed6186376be2d0864f40445c320bd3ea7f81c5e1de00d.
//
// Solidity: event VoteCasted(string indexed nik, uint256 indexed candidateId)
func (_Votechain *VotechainFilterer) WatchVoteCasted(opts *bind.WatchOpts, sink chan<- *VotechainVoteCasted, nik []string, candidateId []*big.Int) (event.Subscription, error) {

	var nikRule []interface{}
	for _, nikItem := range nik {
		nikRule = append(nikRule, nikItem)
	}
	var candidateIdRule []interface{}
	for _, candidateIdItem := range candidateId {
		candidateIdRule = append(candidateIdRule, candidateIdItem)
	}

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "VoteCasted", nikRule, candidateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainVoteCasted)
				if err := _Votechain.contract.UnpackLog(event, "VoteCasted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCasted is a log parse operation binding the contract event 0xfef5593509c07a56349ed6186376be2d0864f40445c320bd3ea7f81c5e1de00d.
//
// Solidity: event VoteCasted(string indexed nik, uint256 indexed candidateId)
func (_Votechain *VotechainFilterer) ParseVoteCasted(log types.Log) (*VotechainVoteCasted, error) {
	event := new(VotechainVoteCasted)
	if err := _Votechain.contract.UnpackLog(event, "VoteCasted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainVoterRegisteredIterator is returned from FilterVoterRegistered and is used to iterate over the raw logs and unpacked data for VoterRegistered events raised by the Votechain contract.
type VotechainVoterRegisteredIterator struct {
	Event *VotechainVoterRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotechainVoterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainVoterRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotechainVoterRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotechainVoterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainVoterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainVoterRegistered represents a VoterRegistered event raised by the Votechain contract.
type VotechainVoterRegistered struct {
	Nik          common.Hash
	VoterAddress common.Address
	Region       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoterRegistered is a free log retrieval operation binding the contract event 0xe8bf381bec3899d7c4d98d7e52cfd45dfe7254b2ceafbb4d6dca1235ed10624d.
//
// Solidity: event VoterRegistered(string indexed nik, address indexed voterAddress, string region)
func (_Votechain *VotechainFilterer) FilterVoterRegistered(opts *bind.FilterOpts, nik []string, voterAddress []common.Address) (*VotechainVoterRegisteredIterator, error) {

	var nikRule []interface{}
	for _, nikItem := range nik {
		nikRule = append(nikRule, nikItem)
	}
	var voterAddressRule []interface{}
	for _, voterAddressItem := range voterAddress {
		voterAddressRule = append(voterAddressRule, voterAddressItem)
	}

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "VoterRegistered", nikRule, voterAddressRule)
	if err != nil {
		return nil, err
	}
	return &VotechainVoterRegisteredIterator{contract: _Votechain.contract, event: "VoterRegistered", logs: logs, sub: sub}, nil
}

// WatchVoterRegistered is a free log subscription operation binding the contract event 0xe8bf381bec3899d7c4d98d7e52cfd45dfe7254b2ceafbb4d6dca1235ed10624d.
//
// Solidity: event VoterRegistered(string indexed nik, address indexed voterAddress, string region)
func (_Votechain *VotechainFilterer) WatchVoterRegistered(opts *bind.WatchOpts, sink chan<- *VotechainVoterRegistered, nik []string, voterAddress []common.Address) (event.Subscription, error) {

	var nikRule []interface{}
	for _, nikItem := range nik {
		nikRule = append(nikRule, nikItem)
	}
	var voterAddressRule []interface{}
	for _, voterAddressItem := range voterAddress {
		voterAddressRule = append(voterAddressRule, voterAddressItem)
	}

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "VoterRegistered", nikRule, voterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainVoterRegistered)
				if err := _Votechain.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoterRegistered is a log parse operation binding the contract event 0xe8bf381bec3899d7c4d98d7e52cfd45dfe7254b2ceafbb4d6dca1235ed10624d.
//
// Solidity: event VoterRegistered(string indexed nik, address indexed voterAddress, string region)
func (_Votechain *VotechainFilterer) ParseVoterRegistered(log types.Log) (*VotechainVoterRegistered, error) {
	event := new(VotechainVoterRegistered)
	if err := _Votechain.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotechainVotingStatusChangedIterator is returned from FilterVotingStatusChanged and is used to iterate over the raw logs and unpacked data for VotingStatusChanged events raised by the Votechain contract.
type VotechainVotingStatusChangedIterator struct {
	Event *VotechainVotingStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotechainVotingStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotechainVotingStatusChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotechainVotingStatusChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotechainVotingStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotechainVotingStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotechainVotingStatusChanged represents a VotingStatusChanged event raised by the Votechain contract.
type VotechainVotingStatusChanged struct {
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVotingStatusChanged is a free log retrieval operation binding the contract event 0x9069a1a16ace751e8690f383e12f87b01e8488ba387e626810bd113fef0417f9.
//
// Solidity: event VotingStatusChanged(bool isActive)
func (_Votechain *VotechainFilterer) FilterVotingStatusChanged(opts *bind.FilterOpts) (*VotechainVotingStatusChangedIterator, error) {

	logs, sub, err := _Votechain.contract.FilterLogs(opts, "VotingStatusChanged")
	if err != nil {
		return nil, err
	}
	return &VotechainVotingStatusChangedIterator{contract: _Votechain.contract, event: "VotingStatusChanged", logs: logs, sub: sub}, nil
}

// WatchVotingStatusChanged is a free log subscription operation binding the contract event 0x9069a1a16ace751e8690f383e12f87b01e8488ba387e626810bd113fef0417f9.
//
// Solidity: event VotingStatusChanged(bool isActive)
func (_Votechain *VotechainFilterer) WatchVotingStatusChanged(opts *bind.WatchOpts, sink chan<- *VotechainVotingStatusChanged) (event.Subscription, error) {

	logs, sub, err := _Votechain.contract.WatchLogs(opts, "VotingStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotechainVotingStatusChanged)
				if err := _Votechain.contract.UnpackLog(event, "VotingStatusChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotingStatusChanged is a log parse operation binding the contract event 0x9069a1a16ace751e8690f383e12f87b01e8488ba387e626810bd113fef0417f9.
//
// Solidity: event VotingStatusChanged(bool isActive)
func (_Votechain *VotechainFilterer) ParseVotingStatusChanged(log types.Log) (*VotechainVotingStatusChanged, error) {
	event := new(VotechainVotingStatusChanged)
	if err := _Votechain.contract.UnpackLog(event, "VotingStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
