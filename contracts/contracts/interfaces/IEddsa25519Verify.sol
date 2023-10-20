// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IEddsa25519Verify {
  function verifySignature(
    bytes memory publicKey,
    string memory message,
    bytes memory signature
  ) external view returns (bool isValid);
}
