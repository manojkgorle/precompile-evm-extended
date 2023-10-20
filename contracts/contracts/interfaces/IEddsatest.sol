// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IEddsaTest {
  function verifySignature(
    string memory publicKey,
    string memory message,
    string memory signature
  ) external view returns (string memory isValid);
}
