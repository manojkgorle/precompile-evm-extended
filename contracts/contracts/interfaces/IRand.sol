//SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IRand {
  function generateRandomNumber(uint32 totNum) external view returns (uint256[] memory genNum);
}
