# Problem

## The anonymity problem

- There have been a lot of protocol exploits & wallet exploits several times.

- Blockchains give anonymity to its users as a default.

- So scammers are all over with their identity hidden.

- Governments are always in a perspective that cryptocurrencies & many things happening on blockchain are shaddy.

## Gas costs of Mapping

- Mappings are one of the most extensively used functions in solidity.
- Mappings usage include ERC20, NFT, Balance storages & many.
- But Mappings are costly, because of the way solidity implements it.

## Signature Schemas

- ECDSA (secp256k1) is used by most of the EVM chains.
- But there are many other widely used schemas like Eddsa25519, sha256, sha512 ,... which are not supported
- This is limiting the protocols to leverage benefits of these curves for authentication & other purposes.

## Random Number generator

- We rely on ChainLink VRF for generating random numbers, where we need to wait for n number of blocks for the random number to return.

# Challenges we ran into

## Challenge 1

- Soldity implements mapping in a clever way. 
- It checks for next vacant slot number, does keccak256 of the slot number & starts storing mapping values from the slot number on (Simplified).
- But in our precompile implementation, we need to do something in those lines to enable multiple mappings supported with single smart contract address.
- There we made a decision to give developer the control of slot initiation to support multiple mappings with `key` parameter.

## Challenge 2

- While dealing with Eddsa25519Verify precompile, initially we returned `bool isValid`.
- Later while testing, noticed that we are obtaining false irrespective of validity.
- Thanks to **Ava Devs**, we figured out a way, and decided to use `string memory isValid` as return statement.

## Challenge 3

- While incorporating KYC, i.e. to allow only people who have succesfully completed their KYC to transact on blockchain.
- We found no way to implement as a Smartcontract. These things must be within execution layer.
- After, extensive reasearch by one of our teammate, we found documentation relevant to this.

```Json
  "txAllowListConfig": {
        "blockTimestamp": 0,
        "adminAddresses": ["0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC"]
      }
```

## Challenge 4

- Random numbers can not be generated the same, in every node running it.
- Each node gets its own random number, when calling the same function.
- So, nodes will never agree on the state of blockchain.


## Challenge 5

- Installation of avalanche go & avalanche precompile evm in Windows Subsystem for Linux 2
- Wrong genesis references given while launching the network caused the network to shut down multiple times in testing.

## Challenge 6

- Adding meta mask support to our nextJs app.
- Sign message implementation from meta mask SDK.