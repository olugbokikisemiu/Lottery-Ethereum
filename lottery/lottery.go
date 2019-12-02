// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lottery

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LotteryABI is the input ABI used to generate the binding from.
const LotteryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPlayer\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enter\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"players\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"selectWinner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LotteryBin is the compiled bytecode used for deploying new contracts.
var LotteryBin = "0x6080604052604051806060016040528060278152602001610b146027913960039080519060200190610032929190610100565b506040518060400160405280601e81526020017f4f6e6c79206d616e616765722063616e2073656c6563742077696e6e657200008152506004908051906020019061007e929190610100565b50604051806060016040528060218152602001610b3b60219139600590805190602001906100ad929190610100565b503480156100ba57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101a5565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014157805160ff191683800117855561016f565b8280016001018555821561016f579182015b8281111561016e578251825591602001919060010190610153565b5b50905061017c9190610180565b5090565b6101a291905b8082111561019e576000816000905550600101610186565b5090565b90565b610960806101b46000396000f3fe60806040526004361061004a5760003560e01c806333a99e041461004f578063481c6a75146100a6578063aa295dce146100fd578063e97dcb6214610169578063f71d96cb14610173575b600080fd5b34801561005b57600080fd5b506100646101ee565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100b257600080fd5b506100bb6103d2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561010957600080fd5b506101126103f7565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561015557808201518184015260208101905061013a565b505050509050019250505060405180910390f35b610171610485565b005b34801561017f57600080fd5b506101ac6004803603602081101561019657600080fd5b8101908080359060200190929190505050610689565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614600490610305576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156102f65780601f106102cb576101008083540402835291602001916102f6565b820191906000526020600020905b8154815290600101906020018083116102d957829003601f168201915b50509250505060405180910390fd5b5060006002805490506103166106c5565b8161031d57fe5b06905060006002828154811061032f57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f193505050501580156103b9573d6000803e3d6000fd5b5060006002816103c991906108da565b50809250505090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600280548060200260200160405190810160405280929190818152602001828054801561047b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610431575b5050505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141560039061059b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382528381815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561058c5780601f106105615761010080835404028352916020019161058c565b820191906000526020600020905b81548152906001019060200180831161056f57829003601f168201915b50509250505060405180910390fd5b5067016345785d8a00003411610619576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f496e73756666696369656e7420616d6f756e742073656e64200000000000000081525060200191505060405180910390fd5b610621610767565b60023390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6002818154811061069657fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60004442600260405160200180848152602001838152602001828054801561074257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116106f8575b505093505050506040516020818303038152906040528051906020012060001c905090565b60001515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514600590610880576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156108715780601f1061084657610100808354040283529160200191610871565b820191906000526020600020905b81548152906001019060200180831161085457829003601f168201915b50509250505060405180910390fd5b5060018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550565b815481835581811115610901578183600052602060002091820191016109009190610906565b5b505050565b61092891905b8082111561092457600081600090555060010161090c565b5090565b9056fea265627a7a723158204277f8a411fb937a64dd83cd36bf348cdb4599e02ce656cadc93e783f967fad464736f6c634300050c00324d616e616765727320617265206e6f7420616c6c6f77656420696e20636f6d7065746974696f6e506c6179657220616c7265616479206a6f696e656420746865206c6f7474657279"

// DeployLottery deploys a new Ethereum contract, binding an instance of Lottery to it.
func DeployLottery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lottery, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LotteryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// Lottery is an auto generated Go binding around an Ethereum contract.
type Lottery struct {
	LotteryCaller     // Read-only binding to the contract
	LotteryTransactor // Write-only binding to the contract
	LotteryFilterer   // Log filterer for contract events
}

// LotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotterySession struct {
	Contract     *Lottery          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryCallerSession struct {
	Contract *LotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryTransactorSession struct {
	Contract     *LotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryRaw struct {
	Contract *Lottery // Generic contract binding to access the raw methods on
}

// LotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryCallerRaw struct {
	Contract *LotteryCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryTransactorRaw struct {
	Contract *LotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLottery creates a new instance of Lottery, bound to a specific deployed contract.
func NewLottery(address common.Address, backend bind.ContractBackend) (*Lottery, error) {
	contract, err := bindLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// NewLotteryCaller creates a new read-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryCaller(address common.Address, caller bind.ContractCaller) (*LotteryCaller, error) {
	contract, err := bindLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryCaller{contract: contract}, nil
}

// NewLotteryTransactor creates a new write-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryTransactor, error) {
	contract, err := bindLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryTransactor{contract: contract}, nil
}

// NewLotteryFilterer creates a new log filterer instance of Lottery, bound to a specific deployed contract.
func NewLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryFilterer, error) {
	contract, err := bindLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryFilterer{contract: contract}, nil
}

// bindLottery binds a generic wrapper to an already deployed contract.
func bindLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.LotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transact(opts, method, params...)
}

// AllPlayer is a free data retrieval call binding the contract method 0xaa295dce.
//
// Solidity: function allPlayer() constant returns(address[])
func (_Lottery *LotteryCaller) AllPlayer(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Lottery.contract.Call(opts, out, "allPlayer")
	return *ret0, err
}

// AllPlayer is a free data retrieval call binding the contract method 0xaa295dce.
//
// Solidity: function allPlayer() constant returns(address[])
func (_Lottery *LotterySession) AllPlayer() ([]common.Address, error) {
	return _Lottery.Contract.AllPlayer(&_Lottery.CallOpts)
}

// AllPlayer is a free data retrieval call binding the contract method 0xaa295dce.
//
// Solidity: function allPlayer() constant returns(address[])
func (_Lottery *LotteryCallerSession) AllPlayer() ([]common.Address, error) {
	return _Lottery.Contract.AllPlayer(&_Lottery.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Lottery *LotteryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lottery.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Lottery *LotterySession) Manager() (common.Address, error) {
	return _Lottery.Contract.Manager(&_Lottery.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Lottery *LotteryCallerSession) Manager() (common.Address, error) {
	return _Lottery.Contract.Manager(&_Lottery.CallOpts)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) constant returns(address)
func (_Lottery *LotteryCaller) Players(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lottery.contract.Call(opts, out, "players", arg0)
	return *ret0, err
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) constant returns(address)
func (_Lottery *LotterySession) Players(arg0 *big.Int) (common.Address, error) {
	return _Lottery.Contract.Players(&_Lottery.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) constant returns(address)
func (_Lottery *LotteryCallerSession) Players(arg0 *big.Int) (common.Address, error) {
	return _Lottery.Contract.Players(&_Lottery.CallOpts, arg0)
}

// Enter is a paid mutator transaction binding the contract method 0xe97dcb62.
//
// Solidity: function enter() returns()
func (_Lottery *LotteryTransactor) Enter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "enter")
}

// Enter is a paid mutator transaction binding the contract method 0xe97dcb62.
//
// Solidity: function enter() returns()
func (_Lottery *LotterySession) Enter() (*types.Transaction, error) {
	return _Lottery.Contract.Enter(&_Lottery.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0xe97dcb62.
//
// Solidity: function enter() returns()
func (_Lottery *LotteryTransactorSession) Enter() (*types.Transaction, error) {
	return _Lottery.Contract.Enter(&_Lottery.TransactOpts)
}

// SelectWinner is a paid mutator transaction binding the contract method 0x33a99e04.
//
// Solidity: function selectWinner() returns(address)
func (_Lottery *LotteryTransactor) SelectWinner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "selectWinner")
}

// SelectWinner is a paid mutator transaction binding the contract method 0x33a99e04.
//
// Solidity: function selectWinner() returns(address)
func (_Lottery *LotterySession) SelectWinner() (*types.Transaction, error) {
	return _Lottery.Contract.SelectWinner(&_Lottery.TransactOpts)
}

// SelectWinner is a paid mutator transaction binding the contract method 0x33a99e04.
//
// Solidity: function selectWinner() returns(address)
func (_Lottery *LotteryTransactorSession) SelectWinner() (*types.Transaction, error) {
	return _Lottery.Contract.SelectWinner(&_Lottery.TransactOpts)
}
