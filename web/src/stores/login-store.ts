/* eslint-disable @typescript-eslint/no-explicit-any */
import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

const xtokenrefresh_name = process.env.XTOKENREFRESH_NAME || '';
const xtoken_name = process.env.XTOKEN_NAME || '';
const xdatauser_name = process.env.XDATAUSER_NAME || '';
const xmenus_name = process.env.XMENUS_NAME || '';

const get_storage_name_plain = (clave: string) => {
  const r = localStorage.getItem(clave);
  if (r) return atob(r);
  return null;
};
const get_storage_name = (clave: string, defecto = '') => {
  const r = localStorage.getItem(clave);
  if (r) return JSON.parse(atob(r));
  return defecto;
};

export const useLoginStore = defineStore('counter', {
  state: () => ({
    token: ref(get_storage_name_plain(xtoken_name)),
    refreshToken: ref(get_storage_name_plain(xtokenrefresh_name)),
    dataUser: ref(get_storage_name(xdatauser_name, '{}')),
    menus: ref(get_storage_name(xmenus_name, '[]')),
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
    setToken(xtoken: string, xrefreshToken: string) {
      this.token = null;
      this.refreshToken = null;
      this.dataUser = {};
      if (!xtoken) {
        this.clearStore();
      } else {
        this.setNewToken(xtoken);
        localStorage.setItem(xtokenrefresh_name, btoa(xrefreshToken));
        this.refreshToken = xrefreshToken;
      }
    },
    setNewToken(xtoken: string) {
      localStorage.setItem(xtoken_name, btoa(xtoken));
      this.token = xtoken;
    },
    setNewTokenRefresh(xrefreshToken: string) {
      localStorage.setItem(xtokenrefresh_name, btoa(xrefreshToken));
      this.refreshToken = xrefreshToken;
    },
    setUser(user: any) {
      localStorage.setItem(xdatauser_name, btoa(JSON.stringify(user)));
      this.dataUser = user;
    },
    setMenus(menus: any) {
      localStorage.setItem(xmenus_name, btoa(JSON.stringify(menus)));
      this.menus = menus;
    },
    setTiempoSession(str: string) {
      this.tiempoSession = str;
    },
    setWsNoti(e: string) {
      this.ws_noti_status = e;
    },
    clearStore() {
      // localStorage.clear();
      for (const clave in process.env) {
        if (process.env.hasOwnProperty(clave)) {
          localStorage.removeItem(process.env[clave] + '');
        }
      }
    },
  },
});
