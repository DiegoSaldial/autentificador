import { PiniaPluginContext } from 'pinia';

const myPiniaPlugin = (context: PiniaPluginContext) => {
  const store = context.store;

  if (store.$id === 'myIndexedDB') {
    store.loadInitialData();
  }
};

export default myPiniaPlugin;
