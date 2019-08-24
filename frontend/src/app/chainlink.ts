export const chainLinkAbi = [{
  constant: false,
  inputs: [{name: 'newCurrent', type: 'bytes32'}],
  name: 'update',
  outputs: [],
  type: 'function'
}, {
  constant: true,
  inputs: [],
  name: 'current',
  outputs: [{name: 'current', type: 'bytes32'}],
  type: 'function'
}, {inputs: [], type: 'constructor'}];

export const chainLinkAddress = '0x2a032eb0af76e6a0315f14b470f1fbe309393416';
