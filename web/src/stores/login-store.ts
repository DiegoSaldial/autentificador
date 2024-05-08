/* eslint-disable @typescript-eslint/no-explicit-any */
import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

const xtokenrefresh_name = process.env.XTOKENREFRESH_NAME || '';
const xtoken_name = process.env.XTOKEN_NAME || '';
const xdatauser_name = process.env.XDATAUSER_NAME || '';
const xmenus_name = process.env.XMENUS_NAME || '';

export const useLoginStore = defineStore('counter', {
  state: () => ({
    token: ref(localStorage.getItem(xtoken_name) || null),
    refreshToken: ref(localStorage.getItem(xtokenrefresh_name) || null),
    dataUser: ref(
      JSON.parse(localStorage.getItem(xdatauser_name) || '{}') || null
    ),
    menus: ref(JSON.parse(localStorage.getItem(xmenus_name) || '[]') || null),
    tiempoSession: ref(''),
  }),
  getters: {
    getToken: (state) => computed(() => state.token),
    getRefreshToken: (state) => computed(() => state.refreshToken),
    getDataUser: (state) => computed(() => state.dataUser),
    getMenus: (state) => computed(() => state.menus),
    getTiempoSession: (state) => computed(() => state.tiempoSession),
  },
  actions: {
    setToken(xtoken: string, xrefreshToken: string) {
      this.token = null;
      this.refreshToken = null;
      this.dataUser = {};
      if (!xtoken) localStorage.clear();
      else {
        this.setNewToken(xtoken);
        localStorage.setItem(xtokenrefresh_name, xrefreshToken);
        this.refreshToken = xrefreshToken;
      }
    },
    setNewToken(xtoken: string) {
      localStorage.setItem(xtoken_name, xtoken);
      this.token = xtoken;
    },
    setUser(user: any) {
      localStorage.setItem(xdatauser_name, JSON.stringify(user));
      this.dataUser = user;
    },
    setMenus(menus: any) {
      localStorage.setItem(xmenus_name, JSON.stringify(menus));
      this.menus = menus;
    },
    setTiempoSession(str: string) {
      this.tiempoSession = str;
    },
  },
});
