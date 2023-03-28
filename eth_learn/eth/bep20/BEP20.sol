// SPDX-License-Identifier: MIT
pragma solidity ^0.7.6;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract BEP20 is ERC20 {
    uint public version = 1001;
    constructor(uint256 initialSupply, string memory name_, string memory symbol_) ERC20(name_, symbol_) payable {
        _mint(msg.sender, initialSupply);
    }
}

