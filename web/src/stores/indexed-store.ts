/* eslint-disable @typescript-eslint/no-explicit-any */
import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

// Configuración de IndexedDB
const dbName = 'myDatabase';
const storeName = 'myStore';

const openDatabase = () => {
  return new Promise((resolve, reject) => {
    const request = indexedDB.open(dbName, 1);

    request.onerror = (event) => {
      reject('Error al abrir la base de datos ' + event);
    };

    request.onsuccess = (event: any) => {
      resolve(event.target.result);
    };

    request.onupgradeneeded = (event: any) => {
      const db = event.target.result;
      db.createObjectStore(storeName, { keyPath: 'clave' });
    };
  });
};

export const getItem = (clave: any, parse: any) => {
  return new Promise(async (resolve, reject) => {
    const db: any = await openDatabase();
    const transaction = db.transaction([storeName], 'readonly');
    const objectStore = transaction.objectStore(storeName);
    const request = objectStore.get(clave);

    request.onsuccess = (event: any) => {
      let v = event.target.result;
      if (v) {
        if (parse) v = JSON.parse(atob(event.target.result.valor));
        else v = atob(event.target.result.valor);
      } else v = parse;
      resolve(v);
    };

    request.onerror = (e: any) => {
      console.log('Error al obtener el item', e);

      reject('Error al obtener el item');
    };
  });
};

export const setItem = (clave: any, valor: any) => {
  return new Promise<void>(async (resolve, reject) => {
    const db: any = await openDatabase();
    const transaction = db.transaction([storeName], 'readwrite');
    const objectStore = transaction.objectStore(storeName);
    const request = objectStore.put({ clave, valor });

    request.onsuccess = () => {
      resolve();
    };

    request.onerror = () => {
      reject('Error al guardar el item');
    };
  });
};

// const get_storage_name_plain = (clave: string, defecto = '') => {
//   const r = localStorage.getItem(clave);
//   if (r) return atob(r);
//   return defecto;
// };
// const get_storage_name = (clave: string, defecto = '') => {
//   const r = localStorage.getItem(clave);
//   if (r) return JSON.parse(atob(r));
//   return defecto;
// };

const xtokenrefresh_name = process.env.XTOKENREFRESH_NAME || '';
const xtoken_name = process.env.XTOKEN_NAME || '';
const xdatauser_name = process.env.XDATAUSER_NAME || '';
const xmenus_name = process.env.XMENUS_NAME || '';

export const useIndexedStore = defineStore('myIndexedDB', {
  state: () => ({
    token: ref(''),
    refreshToken: ref(''),
    dataUser: ref({}),
    menus: ref([]),
    tiempoSession: ref(''),
    ws_noti_status: ref(''),
  }),
  getters: {
    getToken: (state) => computed(() => state.token),
    getRefreshToken: (state) => computed(() => state.refreshToken),
    getDataUser: (state) => computed(() => state.dataUser),
    getMenus: (state) => computed(() => state.menus),
    getTiempoSession: (state) => computed(() => state.tiempoSession),
    getWsNoti: (state) => computed(() => state.ws_noti_status),
  },
  actions: {
    async loadInitialData() {
      const xtn = xtokenrefresh_name;
      this.token = await getItem(xtoken_name, '').then((x: any) => x);
      this.refreshToken = await getItem(xtn, '').then((x: any) => x);
      this.dataUser = await getItem(xdatauser_name, {}).then((x: any) => x);
      this.menus = await getItem(xmenus_name, []).then((x: any) => x);
      /* console.log('token', this.token);
      console.log('dataUser', this.dataUser); */
    },
    async setToken(xtoken: string, xrefreshToken: string) {
      this.token = '';
      this.refreshToken = '';
      this.dataUser = {};
      if (!xtoken) localStorage.clear();
      else {
        this.setNewToken(xtoken);
        await setItem(xtokenrefresh_name, btoa(xrefreshToken));
        this.refreshToken = xrefreshToken;
      }
    },
    async setNewToken(xtoken: any) {
      await setItem(xtoken_name, btoa(xtoken));
      this.token = xtoken;
    },
    async setNewTokenRefresh(xrefreshToken: string) {
      await setItem(xtokenrefresh_name, btoa(xrefreshToken));
      this.refreshToken = xrefreshToken;
    },
    async setUser(user: any) {
      await setItem(xdatauser_name, btoa(JSON.stringify(user)));
      this.dataUser = user;
    },
    async setMenus(menus: any) {
      await setItem(xmenus_name, btoa(JSON.stringify(menus)));
      this.menus = menus;
    },
    setTiempoSession(str: string) {
      this.tiempoSession = str;
    },
    setWsNoti(e: string) {
      this.ws_noti_status = e;
    },
  },
});

// En tu componente, asegúrate de llamar a loadInitialData para cargar los datos iniciales
/* const store = useIndexedStore();
store.loadInitialData(); */
