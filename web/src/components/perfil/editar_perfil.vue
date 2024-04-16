<template>
  <q-dialog v-model="alert" persistent square>
    <q-card flat bordered v-on:click="checkClickSession()">
      <q-card-section>
        <div class="text-h6"> Editar Info Personal </div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-form @submit="onSubmit">
          <div class="row q-col-gutter-xs">
            <div class="col col-sm-6">
              <q-input filled v-model.trim="input.nombres" label="Nombres:" required lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col col-sm-6">
              <q-input filled v-model.trim="input.apellido1" label="Apellido 1:" required lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col col-sm-6">
              <q-input filled v-model.trim="input.apellido2" label="Apellido 2:" lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col col-sm-6">
              <q-input filled v-model.trim="input.celular" label="Celular:" lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col col-sm-6">
              <q-input filled v-model.trim="input.correo" label="Correo:" lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col col-sm-6">
              <q-input filled v-model.trim="input.documento" label="Documento:" lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col col-sm-12">
              <q-input filled v-model.trim="input.direccion" label="Direccion:" lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div> 
          </div>

          <div class="q-mt-md" :align="'right'">
            <q-linear-progress v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
            <q-btn :disable="loading" label="cerrar" color="red" icon="close" square flat @click="cerrar()" />
            <q-btn :disable="loading" label="Guardar" icon="done" type="submit" color="green" square />
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
import { useLoginStore } from 'src/stores/login-store';
import PerfilService from './perfilService';
import Validaciones from './validador';
import click from 'src/shared/session';
import { ref } from 'vue';

export default {
  setup() {
    const store = useLoginStore();
    const perfilService = new PerfilService();
    const validaciones = new Validaciones();
    const loading = ref(false);
    const alert = ref(false);
    const input = ref({})

    const cerrar = ()=> alert.value = false;

    const open = () => {
      alert.value = true;
      me();
    };

    const me = async () => {
      loading.value = true;
      let res = await perfilService.me(); 
      if(res && res.me){
        input.value.id = res.me.usuario.id;
        input.value.nombres = res.me.usuario.nombres;
        input.value.apellido1 = res.me.usuario.apellido1;
        input.value.apellido2 = res.me.usuario.apellido2;
        input.value.celular = res.me.usuario.celular;
        input.value.correo = res.me.usuario.correo;
        input.value.documento = res.me.usuario.documento;
        input.value.direccion = res.me.usuario.direccion;
      }
      loading.value = false;
    }

    const onSubmit = async () => {
      loading.value = true;
      let res = await perfilService.updateUsuarioPerfil(input.value)
      if(res && res.updateUsuarioPerfil) alert.value = false;
      loading.value = false;
    }

    return {
      store,
      alert,
      loading,
      input,
      validaciones,
      checkClickSession: click.setup().checkClickSession,
      cerrar,
      open,
      onSubmit,
    };
  },
};
</script>
