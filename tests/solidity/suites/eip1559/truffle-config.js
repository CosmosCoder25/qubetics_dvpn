// const HDWalletProvider = require('@truffle/hdwallet-provider');

module.exports = {
  networks: {
    // Development network is just left as truffle's default settings
    qubetics: {
      host: '127.0.0.1', // Localhost (default: none)
      port: 8545, // Standard Ethereum port (default: none)
      network_id: '*', // Any network (default: none)
      gas: 5000000, // Gas sent with each transaction
      gasPrice: 1000000000 // 1 gwei (in wei)
    }
  },
  compilers: {
    solc: {
      version: '0.8.18'
    }
  }
}
