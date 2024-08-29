<template>
  <q-dialog v-model="alert" persistent square>
    <q-card flat bordered v-on:click="checkClickSession()">
      <q-card-section>
        <div class="text-h6"> Editar Información Personal </div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-form @submit="onSubmit">
          <div class="row q-col-gutter-xs justify-center">
            <div class="col-xs-12 col-sm-8 ">
              <q-img
                class="zoomer"
                v-if="foto_64" :src="foto_64"
                spinner-color="white" 
              />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.nombres" label="* Nombres:" required lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.apellido1" label="* Apellido 1:" required lazy-rules dense :rules="[(val) => validaciones.val_apellido1(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.apellido2" label="Apellido 2:" dense lazy-rules :rules="[(val) => validaciones.val_apellido2(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.celular" label="* Celular:" dense required lazy-rules :rules="[(val) => validaciones.val_celular(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.correo" label="Correo:" dense lazy-rules :rules="[(val) => validaciones.val_correo(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.documento" label="Documento:" dense lazy-rules :rules="[(val) => validaciones.val_documento(val)]" />
            </div>
            <div class="col-xs-12 col-sm-12">
              <q-input outlined v-model.trim="input.direccion" label="Direccion:" dense lazy-rules :rules="[(val) => validaciones.val_direccion(val)]" />
            </div> 

            <div class="col-xs-12 col-sm-12">
              <q-file style="min-width: 50px" clearable v-model="foto_file" dense accept="image/*" :disable="loading" square outlined color="orange" label="Seleccionar foto de perfil (2MB)" max-file-size="2097152" @update:model-value="filevalue($event)" @rejected="onRejected" >
                <template v-slot:prepend>
                  <q-icon name="upload" />
                </template>
                <q-tooltip> Seleccionar foto de perfil </q-tooltip>
              </q-file>
            </div>

            <div class="col-xs-12 col-sm-12 q-pt-md">
              <q-btn color="grey" icon="fmd_good" dense square outline @click="openGeo()"> Ubicacion geografica </q-btn>
              <small class="q-pl-md"> {{ input.latitud }} {{ input.longitud }} </small>
            </div>
          </div>

          <div vv-show="!input.oauth_id"> 
            <q-separator color="green" class="q-mt-md q-mb-xs" />
            <div class="row">
              <div class="col-xs-12 col-sm-12">
                <q-input outlined v-model.trim="input.password" label="Cambiar clave de acceso:" dense lazy-rules :rules="[(val) => validaciones.val_password_opc(val)]" />
              </div> 
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

    <Geo ref="refGeo" v-on:onpin="onpin" />
  </q-dialog>
</template>

<script>
import { useLoginStore } from 'src/stores/login-store';
import PerfilService from './perfilService';
import Validaciones from 'src/pages/xauth/usuarios/validador';
import click from 'src/shared/session';
import { ref } from 'vue';
import {Notify} from 'quasar'
import Geo from 'src/pages/xauth/geo/geo_modal.vue';
import {init_zoomer} from 'src/pages/xauth/usuarios/zoomer';

export default {
  components:{ Geo },
  setup() {
    const store = useLoginStore();
    const perfilService = new PerfilService();
    const validaciones = new Validaciones();
    const loading = ref(false);
    const alert = ref(false);
    const refGeo = ref();
    const input = ref({});
    const foto_file = ref();
    const foto_64 = ref('');

    const cerrar = ()=> alert.value = false;

    const open = () => {
      alert.value = true;
      foto_file.value = null;
      foto_64.value = '';
      input.value = {};
      me();
    };

    const me = async () => {
      loading.value = true;
      const res = await perfilService.me(); 
      if(res && res.me){
        input.value.id = res.me.usuario.id;
        input.value.nombres = res.me.usuario.nombres;
        input.value.apellido1 = res.me.usuario.apellido1;
        input.value.apellido2 = res.me.usuario.apellido2;
        input.value.celular = res.me.usuario.celular;
        input.value.correo = res.me.usuario.correo;
        input.value.documento = res.me.usuario.documento;
        input.value.direccion = res.me.usuario.direccion;
        input.value.oauth_id = res.me.usuario.oauth_id;
        input.value.latitud = res.me.usuario.latitud;
        input.value.longitud = res.me.usuario.longitud;
        // input.value.password = '';
        getFoto(res.me.usuario);
      }
      loading.value = false;
    }

    const filevalue = (file) => {
      if (file) {
        const reader = new FileReader();
        reader.onload = function (e) {  
          input.value.foto64 = e.target.result;
        };
        reader.readAsDataURL(file);
      }else{
        input.value.foto64 = null;
      }
    };

    const onRejected = (rejectedEntries) => {
      const mf = rejectedEntries[0].failedPropValidation;
      Notify.create({
        type: 'negative',
        message: `${rejectedEntries.length} no cumple la restriccion: ${mf}`,
      });
    };

    const getFoto = async (us) => {
      const url = us.foto_url;
      if(!url) return url;
      const res = await perfilService.get_imagen(url); 
      if(res && res.get_imagen) foto_64.value = res.get_imagen; 
      init_zoomer(); 
    }

    const openGeo = () => {
      const lt = input.value.latitud;
      const ln = input.value.longitud;
      refGeo.value.open(lt,ln);
    };

    const onpin = (lat, lon) => { 
      input.value.latitud = lat;
      input.value.longitud = lon;
    };

    const onSubmit = async () => {
      console.log('ssssss',input.value);
      loading.value = true;
      
      const obj = Object.assign({}, input.value);
      delete obj.oauth_id;

      const res = await perfilService.updateUsuarioPerfil(obj)
      if(res && res.updateUsuarioPerfil) alert.value = false;
      loading.value = false;
    }

    return {
      store,
      alert,
      loading,
      input,
      validaciones,
      foto_file,
      foto_64,
      refGeo,
      checkClickSession: click.setup().checkClickSession,
      cerrar,
      open,
      onSubmit,
      filevalue,
      onRejected,
      openGeo,
      onpin,
    };
  },
};
</script>
