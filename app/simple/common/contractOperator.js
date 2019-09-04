'use strict';

const fs = require("fs");
const path = require('path');

const { FileSystemWallet, Gateway } = require("fabric-network");
const ccpPath = path.resolve(__dirname, '..', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

class ContractOperator {
    constructor(walletBasePath, channelName, contractName) {
        this.walletBasePath = walletBasePath;
        this.channelName = channelName;
        this.contractName = contractName;
    }

    async Init() {
        await this.Connect();
    }

    async EvaluateTransaction(name, ...args) {
        if (args.length > 0) {
            return await this.contract.evaluateTransaction(name, ...args);
        }
        return await this.contract.evaluateTransaction(name);
    }

    async SubmitTransaction(name, ...args) {
        if (args.length > 0) {
            return await this.contract.submitTransaction(name, ...args);
        }
        return await this.contract.submitTransaction(name);
    }

    async Connect() {
        // Create a new gateway for connecting to our peer node.
        const gateway = await this.connectGateway();

        // Get the network (channel) our contract is deployed to.
        this.network = await gateway.getNetwork(this.channelName);

        // Get the contract from the network.
        this.contract = this.network.getContract(this.contractName);
    }

    async connectGateway() {
        const walletPath = path.join(this.walletBasePath, 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });
        return gateway;
    }
}

module.exports = ContractOperator;