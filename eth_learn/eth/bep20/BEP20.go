// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bep20

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
)

// Bep20MetaData contains all meta data concerning the Bep20 contract.
var Bep20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526103e960065560405162000e2f38038062000e2f833981810160405260608110156200002f57600080fd5b8151602083018051604051929492938301929190846401000000008211156200005757600080fd5b9083019060208201858111156200006d57600080fd5b82516401000000008111828201881017156200008857600080fd5b82525081516020918201929091019080838360005b83811015620000b75781810151838201526020016200009d565b50505050905090810190601f168015620000e55780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010957600080fd5b9083019060208201858111156200011f57600080fd5b82516401000000008111828201881017156200013a57600080fd5b82525081516020918201929091019080838360005b83811015620001695781810151838201526020016200014f565b50505050905090810190601f168015620001975780820380516001836020036101000a031916815260200191505b50604052505082518391508290620001b790600390602085019062000368565b508051620001cd90600490602084019062000368565b50506005805460ff1916601217905550620001e93384620001f2565b50505062000414565b6001600160a01b0382166200024e576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6200025c6000838362000301565b62000278816002546200030660201b6200058c1790919060201c565b6002556001600160a01b03821660009081526020818152604090912054620002ab9183906200058c62000306821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b505050565b60008282018381101562000361576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620003a05760008555620003eb565b82601f10620003bb57805160ff1916838001178555620003eb565b82800160010185558215620003eb579182015b82811115620003eb578251825591602001919060010190620003ce565b50620003f9929150620003fd565b5090565b5b80821115620003f95760008155600101620003fe565b610a0b80620004246000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806354fd4d501161007157806354fd4d501461021057806370a082311461021857806395d89b411461023e578063a457c2d714610246578063a9059cbb14610272578063dd62ed3e1461029e576100b4565b806306fdde03146100b9578063095ea7b31461013657806318160ddd1461017657806323b872dd14610190578063313ce567146101c657806339509351146101e4575b600080fd5b6100c16102cc565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100fb5781810151838201526020016100e3565b50505050905090810190601f1680156101285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101626004803603604081101561014c57600080fd5b506001600160a01b038135169060200135610362565b604080519115158252519081900360200190f35b61017e61037f565b60408051918252519081900360200190f35b610162600480360360608110156101a657600080fd5b506001600160a01b03813581169160208101359091169060400135610385565b6101ce61040c565b6040805160ff9092168252519081900360200190f35b610162600480360360408110156101fa57600080fd5b506001600160a01b038135169060200135610415565b61017e610463565b61017e6004803603602081101561022e57600080fd5b50356001600160a01b0316610469565b6100c1610484565b6101626004803603604081101561025c57600080fd5b506001600160a01b0381351690602001356104e5565b6101626004803603604081101561028857600080fd5b506001600160a01b03813516906020013561054d565b61017e600480360360408110156102b457600080fd5b506001600160a01b0381358116916020013516610561565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103585780601f1061032d57610100808354040283529160200191610358565b820191906000526020600020905b81548152906001019060200180831161033b57829003601f168201915b5050505050905090565b600061037661036f6105ed565b84846105f1565b50600192915050565b60025490565b60006103928484846106dd565b6104028461039e6105ed565b6103fd85604051806060016040528060288152602001610940602891396001600160a01b038a166000908152600160205260408120906103dc6105ed565b6001600160a01b031681526020810191909152604001600020549190610838565b6105f1565b5060019392505050565b60055460ff1690565b60006103766104226105ed565b846103fd85600160006104336105ed565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549061058c565b60065481565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103585780601f1061032d57610100808354040283529160200191610358565b60006103766104f26105ed565b846103fd856040518060600160405280602581526020016109b1602591396001600061051c6105ed565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190610838565b600061037661055a6105ed565b84846106dd565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6000828201838110156105e6576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3390565b6001600160a01b0383166106365760405162461bcd60e51b815260040180806020018281038252602481526020018061098d6024913960400191505060405180910390fd5b6001600160a01b03821661067b5760405162461bcd60e51b81526004018080602001828103825260228152602001806108f86022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166107225760405162461bcd60e51b81526004018080602001828103825260258152602001806109686025913960400191505060405180910390fd5b6001600160a01b0382166107675760405162461bcd60e51b81526004018080602001828103825260238152602001806108d56023913960400191505060405180910390fd5b6107728383836108cf565b6107af8160405180606001604052806026815260200161091a602691396001600160a01b0386166000908152602081905260409020549190610838565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546107de908261058c565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156108c75760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561088c578181015183820152602001610874565b50505050905090810190601f1680156108b95780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122078970b112aecae18f3afc160fa58903eea93ed328628f20031f6d38886be575164736f6c63430007060033",
}

// Bep20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Bep20MetaData.ABI instead.
var Bep20ABI = Bep20MetaData.ABI

// Bep20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Bep20MetaData.Bin instead.
var Bep20Bin = Bep20MetaData.Bin

// DeployBep20 deploys a new Ethereum contract, binding an instance of Bep20 to it.
func DeployBep20(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, name_ string, symbol_ string) (common.Address, *types.Transaction, *Bep20, error) {
	parsed, err := Bep20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Bep20Bin), backend, initialSupply, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bep20{Bep20Caller: Bep20Caller{contract: contract}, Bep20Transactor: Bep20Transactor{contract: contract}, Bep20Filterer: Bep20Filterer{contract: contract}}, nil
}

// Bep20 is an auto generated Go binding around an Ethereum contract.
type Bep20 struct {
	Bep20Caller     // Read-only binding to the contract
	Bep20Transactor // Write-only binding to the contract
	Bep20Filterer   // Log filterer for contract events
}

// Bep20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Bep20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bep20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Bep20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bep20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bep20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bep20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bep20Session struct {
	Contract     *Bep20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bep20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bep20CallerSession struct {
	Contract *Bep20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Bep20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bep20TransactorSession struct {
	Contract     *Bep20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bep20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Bep20Raw struct {
	Contract *Bep20 // Generic contract binding to access the raw methods on
}

// Bep20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bep20CallerRaw struct {
	Contract *Bep20Caller // Generic read-only contract binding to access the raw methods on
}

// Bep20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bep20TransactorRaw struct {
	Contract *Bep20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBep20 creates a new instance of Bep20, bound to a specific deployed contract.
func NewBep20(address common.Address, backend bind.ContractBackend) (*Bep20, error) {
	contract, err := bindBep20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bep20{Bep20Caller: Bep20Caller{contract: contract}, Bep20Transactor: Bep20Transactor{contract: contract}, Bep20Filterer: Bep20Filterer{contract: contract}}, nil
}

// NewBep20Caller creates a new read-only instance of Bep20, bound to a specific deployed contract.
func NewBep20Caller(address common.Address, caller bind.ContractCaller) (*Bep20Caller, error) {
	contract, err := bindBep20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bep20Caller{contract: contract}, nil
}

// NewBep20Transactor creates a new write-only instance of Bep20, bound to a specific deployed contract.
func NewBep20Transactor(address common.Address, transactor bind.ContractTransactor) (*Bep20Transactor, error) {
	contract, err := bindBep20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bep20Transactor{contract: contract}, nil
}

// NewBep20Filterer creates a new log filterer instance of Bep20, bound to a specific deployed contract.
func NewBep20Filterer(address common.Address, filterer bind.ContractFilterer) (*Bep20Filterer, error) {
	contract, err := bindBep20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bep20Filterer{contract: contract}, nil
}

// bindBep20 binds a generic wrapper to an already deployed contract.
func bindBep20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Bep20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bep20 *Bep20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bep20.Contract.Bep20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bep20 *Bep20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bep20.Contract.Bep20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bep20 *Bep20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bep20.Contract.Bep20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bep20 *Bep20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bep20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bep20 *Bep20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bep20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bep20 *Bep20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bep20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bep20 *Bep20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bep20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bep20 *Bep20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bep20.Contract.Allowance(&_Bep20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bep20 *Bep20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bep20.Contract.Allowance(&_Bep20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bep20 *Bep20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bep20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bep20 *Bep20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bep20.Contract.BalanceOf(&_Bep20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bep20 *Bep20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bep20.Contract.BalanceOf(&_Bep20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bep20 *Bep20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bep20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bep20 *Bep20Session) Decimals() (uint8, error) {
	return _Bep20.Contract.Decimals(&_Bep20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bep20 *Bep20CallerSession) Decimals() (uint8, error) {
	return _Bep20.Contract.Decimals(&_Bep20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bep20 *Bep20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bep20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bep20 *Bep20Session) Name() (string, error) {
	return _Bep20.Contract.Name(&_Bep20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bep20 *Bep20CallerSession) Name() (string, error) {
	return _Bep20.Contract.Name(&_Bep20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bep20 *Bep20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bep20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bep20 *Bep20Session) Symbol() (string, error) {
	return _Bep20.Contract.Symbol(&_Bep20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bep20 *Bep20CallerSession) Symbol() (string, error) {
	return _Bep20.Contract.Symbol(&_Bep20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bep20 *Bep20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bep20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bep20 *Bep20Session) TotalSupply() (*big.Int, error) {
	return _Bep20.Contract.TotalSupply(&_Bep20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bep20 *Bep20CallerSession) TotalSupply() (*big.Int, error) {
	return _Bep20.Contract.TotalSupply(&_Bep20.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Bep20 *Bep20Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bep20.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Bep20 *Bep20Session) Version() (*big.Int, error) {
	return _Bep20.Contract.Version(&_Bep20.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Bep20 *Bep20CallerSession) Version() (*big.Int, error) {
	return _Bep20.Contract.Version(&_Bep20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bep20 *Bep20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bep20 *Bep20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.Approve(&_Bep20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bep20 *Bep20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.Approve(&_Bep20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bep20 *Bep20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bep20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bep20 *Bep20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.DecreaseAllowance(&_Bep20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bep20 *Bep20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.DecreaseAllowance(&_Bep20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bep20 *Bep20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bep20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bep20 *Bep20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.IncreaseAllowance(&_Bep20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bep20 *Bep20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.IncreaseAllowance(&_Bep20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bep20 *Bep20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bep20 *Bep20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.Transfer(&_Bep20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bep20 *Bep20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.Transfer(&_Bep20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bep20 *Bep20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bep20 *Bep20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.TransferFrom(&_Bep20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bep20 *Bep20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bep20.Contract.TransferFrom(&_Bep20.TransactOpts, sender, recipient, amount)
}

// Bep20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bep20 contract.
type Bep20ApprovalIterator struct {
	Event *Bep20Approval // Event containing the contract specifics and raw log

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
func (it *Bep20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bep20Approval)
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
		it.Event = new(Bep20Approval)
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
func (it *Bep20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bep20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bep20Approval represents a Approval event raised by the Bep20 contract.
type Bep20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bep20 *Bep20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Bep20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bep20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Bep20ApprovalIterator{contract: _Bep20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bep20 *Bep20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Bep20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bep20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bep20Approval)
				if err := _Bep20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bep20 *Bep20Filterer) ParseApproval(log types.Log) (*Bep20Approval, error) {
	event := new(Bep20Approval)
	if err := _Bep20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bep20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bep20 contract.
type Bep20TransferIterator struct {
	Event *Bep20Transfer // Event containing the contract specifics and raw log

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
func (it *Bep20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bep20Transfer)
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
		it.Event = new(Bep20Transfer)
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
func (it *Bep20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bep20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bep20Transfer represents a Transfer event raised by the Bep20 contract.
type Bep20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bep20 *Bep20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Bep20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bep20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Bep20TransferIterator{contract: _Bep20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bep20 *Bep20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Bep20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bep20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bep20Transfer)
				if err := _Bep20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bep20 *Bep20Filterer) ParseTransfer(log types.Log) (*Bep20Transfer, error) {
	event := new(Bep20Transfer)
	if err := _Bep20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
