<template>
  <div class="q-pa-md column items-center">
    <h5 class="q-mb-lg">Acceder al sistema</h5>
    <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
      <q-input
        filled
        v-model.trim="username"
        label="Nombre de usuario"
        lazy-rules
        dense
        :color="$q.dark.isActive ? 'orange' : 'primary'"
        :rules="[(val) => (val && val.length > 0) || 'dato obligatorio']"
      >
        <template v-slot:prepend>
          <q-icon name="person" />
        </template>
      </q-input>

      <q-input
        filled
        :type="pwd"
        v-model.trim="clave"
        label="Clave de acceso"
        lazy-rules
        dense
        :color="$q.dark.isActive ? 'orange' : 'primary'"
        :rules="[(val) => (val && val.length > 0) || 'dato obligatorio']"
      >
        <template v-slot:prepend>
          <q-icon name="key" />
        </template>
        <template v-slot:append>
          <q-btn flat dense size="small" @click="changePWD()">
            <q-icon name="visibility" />
          </q-btn>
        </template>
      </q-input>

      <div class="column items-center">
        <q-linear-progress
          v-if="loading"
          dark
          rounded
          indeterminate
          color="secondary"
          class="q-mb-sm"
        />
        <q-btn
          :disable="loading"
          icon="person"
          stretch
          label="Ingresar"
          type="submit"
          color="orange"
        />
      </div>
    </q-form>
  </div>
</template>

<script>
import { ref } from 'vue';
import LoginService from './loginService';
import { useLoginStore } from 'stores/login-store';
import MenusService from 'pages/xauth/usuarios/menuService';
import MeService from './meService';

export default {
  setup() {
    const username = ref('admin');
    const clave = ref('admin');
    const loading = ref(false);
    const service = new LoginService();
    const meService = new MeService();
    const menuService = new MenusService();
    const useLogin = useLoginStore();
    const pwd = ref('password');

    async function getMe() {
      loading.value = true;
      let meres = await meService
        .me()
        .then((e) => e)
        .catch((e) => e);
      if (meres.me) {
        await cargarMenus(meres);
        useLogin.setUser(meres.me);
      }
      loading.value = false;
    }

    const cargarMenus = async (meres) => {
      let menus = await menuService.menus_by_usuario(meres.me.usuario.id);
      const menuItemsAgrupados = menus.menus_by_usuario.reduce(
        (grupos, item) => {
          const grupoId = item.grupo;
          const grupoExistente = grupos.find((grupo) =>
            grupo.some((obj) => obj.grupo === grupoId)
          );

          if (grupoExistente) {
            grupoExistente.push(item);
          } else {
            grupos.push([item]);
          }

          return grupos;
        },
        []
      );

      useLogin.setMenus(menuItemsAgrupados);
    };

    const changePWD = () =>
      pwd.value == 'text' ? (pwd.value = 'password') : (pwd.value = 'text');

    return {
      username,
      clave,
      loading,
      pwd,
      changePWD,

      async onSubmit() {
        loading.value = true;
        let res = await service
          .login(username.value, clave.value)
          .then((x) => x)
          .catch((e) => e);
        loading.value = false;

        if (res.login) {
          const l = res.login;
          useLogin.setToken(l.token, l.refreshToken);
          getMe();
        }
      },

      onReset() {
        username.value = null;
        clave.value = null;
      },
    };
  },
};
</script>
src/pages/xauth/usuarios/menuService
