/* eslint-disable @typescript-eslint/no-explicit-any */
import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

export const useLoginStore = defineStore('counter', {
  state: () => ({
    token: ref(localStorage.getItem('xtoken') || null),
    refreshToken: ref(localStorage.getItem('xrefreshToken') || null),
    dataUser: ref(
      JSON.parse(localStorage.getItem('xdataUser') || '{}') || null
    ),
    menus: ref(JSON.parse(localStorage.getItem('xmenus') || '[]') || null),
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
        localStorage.setItem('xrefreshToken', xrefreshToken);
        this.refreshToken = xrefreshToken;
      }
    },
    setNewToken(xtoken: string) {
      localStorage.setItem('xtoken', xtoken);
      this.token = xtoken;
    },
    setUser(user: any) {
      localStorage.setItem('xdataUser', JSON.stringify(user));
      this.dataUser = user;
    },
    setMenus(menus: any) {
      localStorage.setItem('xmenus', JSON.stringify(menus));
      this.menus = menus;
    },
    setTiempoSession(str: string) {
      this.tiempoSession = str;
    },
  },
});
