// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IMapping {
  function addressToUint(string memory key, address _address, uint256 _uint) external;

  function getAddressToUint(string memory key, address _address) external view returns (uint256 _uint);

  function uintToString(string memory key, uint256 _uint, string memory _string) external;

  function getUintToString(string memory key, uint _uint) external view returns (string memory _string);

  function uintToUint(string memory key, uint256 _uint1, uint256 _uint2) external;

  function getUintToUint(string memory key, uint256 _uint1) external view returns (uint256 _uint2);
}
