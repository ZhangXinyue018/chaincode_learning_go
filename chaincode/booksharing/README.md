# Book sharing chaincode

## Features v1
- User can create itself
- User can upload a book
- User can share a book

### Domains
- `Account`: Name, PubKey
- `Books`: ID, Name, Owner, Format, Content, Remark, Reminder

### Implementation
- Book is encrypt by a owner, this decrypt method/key is shared by owner to another user offline
- To help owner remember the encrypt method, a reminder is provided, owner can use it in any ways eg. Content is symmetric encrypted by a password, owner use his/her pubkey to encrypt this password as a reminder where only owner's private key can decrypt this password

### Restrictions
- A normal function input must have: `input`, `account`, `sig`


## Instruction
- Please set up hyperledger network first
- run ./startBookSharing.sh
- enter cli docker and run this: peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n booksharing -c '{"function":"CreateAccount","Args":["{\"Name\":\"testnew123\",\"PublicKey\":\"thisisanunknownpubkey\"}", "testnew123", "what ever sig"]}'
- check result on couchDB