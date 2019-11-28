'use strict';

const { FileSystemWallet, Gateway, X509WalletMixin } = require("fabric-network");
const ccpPath = path.resolve(__dirname, '..', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

// todo: fully implement later
class WalletOperator {
    constructor() {
        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        this.wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);
    }

    async Init() {
        // Create a new gateway for connecting to our peer node.
        this.gateway = new Gateway();
        const wallet = this.wallet;
        await this.gateway.connect(ccp, { wallet, identity: 'admin', discovery: { enabled: false } });

        // Get the CA client object from the gateway for interacting with the CA.
        this.ca = this.gateway.getClient().getCertificateAuthority();
        this.adminIdentity = this.gateway.getCurrentIdentity();
    }

    async enrollUser(userName) {
        try {
            // Check to see if we've already enrolled the user.
            const userExists = await this.wallet.exists(userName);
            if (userExists) {
                console.log('An identity for the user "user1" already exists in the wallet');
                return;
            }

            // Check to see if we've already enrolled the admin user.
            const adminExists = await this.wallet.exists('admin');
            if (!adminExists) {
                console.log('An identity for the admin user "admin" does not exist in the wallet');
                console.log('Run the enrollAdmin.js application before retrying');
                return;
            }

            // Register the user, enroll the user, and import the new identity into the wallet.
            const secret = await this.ca.register({ affiliation: 'org1.department1', enrollmentID: userName, role: 'client' }, adminIdentity);
            const enrollment = await this.ca.enroll({ enrollmentID: userName, enrollmentSecret: secret });
            const userIdentity = X509WalletMixin.createIdentity('Org1', enrollment.certificate, enrollment.key.toBytes());
            this.wallet.import(userName, userIdentity);

        } catch (error) {
            console.error(`Failed to register user ${userName}: ${error}`);
            process.exit(1);
        }
    }
}