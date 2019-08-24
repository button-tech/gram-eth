export const swapContractAbi = [
  {
    constant: false,
    inputs: [
      {
        internalType: 'string',
        name: 'tonAddress',
        type: 'string'
      }
    ],
    name: 'sendTon',
    outputs: [],
    payable: true,
    stateMutability: 'payable',
    type: 'function'
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'string',
        name: 'tonAddress',
        type: 'string'
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'sumToSend',
        type: 'uint256'
      }
    ],
    name: 'EtherRecieved',
    type: 'event'
  },
  {
    constant: true,
    inputs: [],
    name: 'askForRate',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256'
      }
    ],
    payable: false,
    stateMutability: 'view',
    type: 'function'
  }
];
export const swapAddress = '0xe89ce7caabe4c73f8aa4173e022185d67cf8780e';
