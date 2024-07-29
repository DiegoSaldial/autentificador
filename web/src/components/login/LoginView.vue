<template>
  <div class="q-pa-md column items-center">
    <q-card>
      <q-card-section class="text-center">
        <h5 class="q-mb-none q-mt-sm">Acceder al sistema</h5>
      </q-card-section>

      <q-card-section>
        <q-form @submit="onSubmit()" @reset="onReset" class="q-gutter-md">
          <q-input filled v-model.trim="username" label="Nombre de usuario" lazy-rules dense :color="$q.dark.isActive ? 'orange' : 'primary'" :rules="[(val) => (val && val.length > 0) || 'dato obligatorio']" >
            <template v-slot:prepend>
              <q-icon name="person" />
            </template>
          </q-input>

          <q-input filled :type="pwd" v-model.trim="clave" label="Clave de acceso" lazy-rules dense :color="$q.dark.isActive ? 'orange' : 'primary'" :rules="[(val) => (val && val.length > 0) || 'dato obligatorio']" >
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
              v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
            <q-btn :disable="loading" icon="person" stretch label="Ingresar" type="submit" color="orange" />
          </div>
        </q-form>

        <div class="row text-right" v-if="accept_oauth">
          <div class="col-12">
            <q-btn flat class="q-mt-md q-pr-xs" size="sm" color="grey" :loading="loading" @click="googleLogin()">
              <q-icon name="hive" size="xs" left></q-icon>
              Acceder con Google
            </q-btn>
          </div>
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';
import LoginService from './loginService';
import { useLoginStore } from 'stores/login-store';
import MeService from './meService';
import { cargarMenus } from './utils';
import { loginGoogle } from './xfirebaseAuth';
import { Notify } from 'quasar';
import { useIndexedStore } from 'stores/indexed-store'

export default {
  setup() {
    const username = ref('');
    const clave = ref('');
    const loading = ref(false);
    const service = new LoginService();
    const meService = new MeService();
    const useLogin = useLoginStore();
    const pwd = ref('password');
    const accept_oauth = ref(false);
    const useIndexed = useIndexedStore();

    const onSubmit = async (u = null, p = null) => {
      const external = process.env.EXTERNAL_LOGIN;
      loading.value = true;
      let user = username.value;
      let pass = clave.value;
      const ext = !!external;
      // console.log('>>>', u, p, external, ext);
      if (u) user = u;
      if (p) pass = p;

      const res = await service.login(user, pass, ext);
      loading.value = false;

      if (res.login) {
        const l = res.login;
        useLogin.setToken(l.token, l.refreshToken);
        useIndexed.setToken(l.token, l.refreshToken);
        getMe();
      }
    };

    async function getMe() {
      loading.value = true;
      const meres = await meService.me();
      if (meres.me) {
        const menuItemsAgrupados = await cargarMenus(meres.me.menus);
        useLogin.setMenus(menuItemsAgrupados);
        useLogin.setUser(meres.me);
      }
      loading.value = false;
    }

    const changePWD = () =>
      pwd.value == 'text' ? (pwd.value = 'password') : (pwd.value = 'text');

    const googleLogin = async () => {
      loading.value = true;
      const d = await loginGoogle();
      if (d && d.user) {
        const res = await service.createOauth(d.user);
        if (res && res.createOauth) {
          onSubmit(d.user.username, d.user.password);
        }
      } else {
        Notify.create({ message: d.err, color: 'negative' });
      }
      loading.value = false;
    };

    onMounted(() => {
      accept_oauth.value = process.env.ACCEPT_OAUTH;
    });

    return {
      username,
      clave,
      loading,
      pwd,
      accept_oauth,
      useIndexed,
      changePWD,
      googleLogin,
      onSubmit,

      onReset() {
        username.value = null;
        clave.value = null;
      },
    };
  },
};
</script> 
