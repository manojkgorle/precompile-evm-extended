// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IEddsa25519Verify {
  function verifySignature(
    string memory publicKey,
    string memory message,
    string memory signature
  ) external view returns (bool isValid);
}
