<template>
  <div class="q-pa-md" id="relogin">
    <q-dialog v-model="dialog" persistent>
      <q-card>
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">Vuelve a iniciar session</div>
          <q-space />
          <q-btn icon="close" flat round dense @click="handleCancel" />
        </q-card-section>

        <q-card-section class="row items-center">
          <q-form @submit="handleAccept" @reset="onReset" class="q-gutter-md">
            <q-input filled v-model.trim="username" label="Nombre de usuario" lazy-rules dense :color="$q.dark.isActive ? 'orange' : 'primary'" :rules="[(val) => (val && val.length > 0) || 'dato obligatorio']" >
              <template v-slot:prepend>
                <q-icon name="person" />
              </template>
            </q-input>

            <q-input filled :type="pwd" v-model.trim="clave" label="Clave de acceso" lazy-rules dense :color="$q.dark.isActive ? 'orange' : 'primary'" :rules="[(val) => (val && val.length > 0) || 'dato obligatorio']">
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
              <q-linear-progress v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
              <q-btn :disable="loading" icon="person" stretch label="Ingresar" type="submit" color="orange" />
            </div>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script>
import { ref } from 'vue';
import { useLoginStore } from 'stores/login-store';
import LoginService from 'src/components/login/loginService';

export default {
  name: 're_login',
  setup() {
    const dialog = ref(false);
    const username = ref('');
    const clave = ref('');
    const loading = ref(false);
    const pwd = ref('password');
    const service = new LoginService();
    const useLogin = useLoginStore();
    let resolvePromise;

    const changePWD = () =>
      pwd.value == 'text' ? (pwd.value = 'password') : (pwd.value = 'text');

    const openDialog = () => {
      dialog.value = true;

      return new Promise((resolve) => {
        resolvePromise = resolve;
      });
    };

    const handleAccept = async () => {
      loading.value = true;
      const res = await service.login(username.value, clave.value);
      if (res.login) {
        const l = res.login;
        // console.log('res', l.token);
        useLogin.setNewToken(l.token);
        useLogin.setNewTokenRefresh(l.refreshToken);
        resolvePromise(true);
      } else {
        resolvePromise(false);
      }
      dialog.value = false;
      loading.value = false;
    };

    const handleCancel = async () => {
      resolvePromise(false);
      dialog.value = false;
    };

    return {
      dialog,
      openDialog,
      handleAccept,
      handleCancel,
      username,
      clave,
      loading,
      pwd,
      changePWD,
      onReset() {
        username.value = null;
        clave.value = null;
      },
    };
  },
};
</script>
