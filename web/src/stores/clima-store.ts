/* eslint-disable @typescript-eslint/no-explicit-any */
import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { clima } from 'src/shared/clima/clima';

const keep_clima = process.env.KEEP_CLIMA || '';
const x_clima = process.env.X_CLIMA || '';

const get_storage_name = (clave: string, defecto: any) => {
  const r = localStorage.getItem(clave);
  if (r) return JSON.parse(atob(r));
  return defecto;
};

export const useClimaStore = defineStore('counterClima', {
  state: () => ({
    clima: ref(get_storage_name(x_clima, {})),
  }),
  getters: {
    getClima: (state) => computed(() => state.clima),
  },
  actions: {
    async setearClima() {
      // console.log('>>>', this.clima);

      if (!this.clima.ciudad && keep_clima) {
        const xclima = await clima();
        this.clima = xclima;
        localStorage.setItem(x_clima, btoa(JSON.stringify(xclima)));
      }
    },
  },
});
