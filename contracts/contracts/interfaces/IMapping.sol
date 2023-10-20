// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IMapping {
  function mapAddUint(string memory key, address _address, uint256 balance) external;

  function getMapAddUint(string memory key, address _address) external view returns (uint256 balance);
}
