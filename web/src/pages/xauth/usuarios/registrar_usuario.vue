<template>
  <q-dialog v-model="alert" persistent square>
    <q-card flat bordered v-on:click="checkClickSession()">
      <q-card-section>
        <div class="text-h6">{{ input.id ? 'Actualizar' : 'Registrar' }}</div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-form @submit="onSubmit">
          <div class="row q-col-gutter-xs justify-center">
            <div class="col-xs-12 col-sm-8 ">
              <q-img
                v-if="foto_64" :src="foto_64"
                spinner-color="white" 
              />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.nombres" label="* Nombres" dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.apellido1" label="* Apellido 1" lazy-rules dense :rules="[(val) => validaciones.val_apellido1(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.apellido2" label="Apellido 2" lazy-rules dense counter :rules="[(val) => validaciones.val_apellido2(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.documento" label="Documento" lazy-rules dense counter :rules="[(val) => validaciones.val_documento(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.celular" label="Celular" lazy-rules dense counter :rules="[(val) => validaciones.val_celular(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-radio v-model="input.sexo" val="M" label="Masculino" class="q-pa-none" />
              <q-radio v-model="input.sexo" val="F" label="Femenino" class="q-pa-none" />
            </div>
            <div class="col-xs-12 col-sm-12">
              <q-input outlined v-model.trim="input.correo" label="Correo" lazy-rules dense counter :rules="[(val) => validaciones.val_correo(val)]" />
            </div>
            <div class="col-xs-12 col-sm-12">
              <q-input outlined v-model.trim="input.direccion" label="Direccion" lazy-rules dense counter :rules="[(val) => validaciones.val_direccion(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.username" label="* username" lazy-rules dense counter :rules="[(val) => validaciones.val_username(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.password" label="password" :placeholder="input.id ? 'vacio sin cambios' : ''" lazy-rules dense counter :rules="[(val) => validaciones.val_password(val, input)]" />
            </div>
            <div class="col-xs-12 col-sm-12">
              <q-file style="min-width: 50px" clearable v-model="foto_file" dense accept="image/*" :disable="loading" square outlined color="orange" label="Seleccionar foto de perfil" max-file-size="2097152" @update:model-value="filevalue($event)" @rejected="onRejected" >
                <template v-slot:prepend>
                  <q-icon name="upload" />
                </template>
                <q-tooltip> Seleccionar foto de perfil </q-tooltip>
              </q-file>
            </div>

            <div class="col-xs-12 col-sm-12 q-mt-md">
              <q-list >
                <q-expansion-item header-class="text-purple" default-opened expand-separator icon="group_add" label="* Roles" caption="Un rol contiene un grupo de permisos" >
                  <q-table flat color="orange" :loading="loading_roles" title="" hide-pagination :rows-per-page-options="[0]" dense :rows="roles" :columns="columnas_rols" row-key="nombre" selection="multiple" v-model:selected="roles_sel" />
                </q-expansion-item>

                <q-expansion-item header-class="text-orange q-mt-lg" expand-separator icon="key" label="Permisos sueltos" caption="Independientes del rol" >
                  <q-table flat color="orange" :loading="loading_perms" title="" hide-pagination :rows-per-page-options="[0]" dense :rows="permisos" :columns="columnas_perm" row-key="metodo" selection="multiple" v-model:selected="permisos_sel">
                    <template v-slot:body-cell-nombre="props">
                      <q-td :props="props">
                        <p class="q-my-none">
                          {{ props.row.nombre }}
                          <q-tooltip class="bg-orange">
                            {{ props.row.metodo }}
                          </q-tooltip>
                        </p>
                      </q-td>
                    </template>
                  </q-table>
                </q-expansion-item>

                <q-expansion-item header-class="text-green q-mt-lg" expand-separator icon="menu_open" label="Menus" caption="Opciones en el menu lateral" >
                  <q-table flat color="orange" :loading="loading_menus" title="" hide-pagination :rows-per-page-options="[0]" dense :rows="menus" :columns="columnas_menu" row-key="label" selection="multiple" v-model:selected="menus_sel" />
                </q-expansion-item>
              </q-list>
            </div>  
          </div>

          <div class="q-mt-md" :align="'right'">
            <q-linear-progress v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
            <q-btn :disable="loading" label="cerrar" color="red" icon="close" square flat @click="cerrar()" />
            <q-btn :disable="loading" :label="input.id ? 'Actualizar' : 'Registrar'" icon="done" type="submit" color="green" square />
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
import { ref } from 'vue';
import { InputNewUsuario } from './type_usuarios';
import UsuariosService from './usuariosService';
import RolesService from 'pages/xauth/roles/rolesService';
import PermisoService from 'pages/xauth/permisos/permisoService';
import MenusService from './menuService';
import { Notify } from 'quasar';
import Validaciones from './validador';
import click from 'src/shared/session';

const columnas_rols = [
  { name: 'nombre', label: '', field: 'nombre', align: 'left' },
  { name: 'descripcion', label: '', field: 'descripcion', align: 'left' },
  { name: 'jerarquia', label: '', field: 'jerarquia', align: 'left' },
];
const columnas_perm = [
  { name: 'nombre', label: '', field: 'nombre', align: 'left' },
  { name: 'jerarquia', label: '', field: 'descripcion', align: 'left' },
];
const columnas_menu = [
  { name: 'label', label: '', field: 'label', align: 'left' },
  { name: 'oreden', label: '', field: 'orden', align: 'left' },
];

export default {
  setup(_, vue) {
    const alert = ref(false);
    const loading = ref(false);
    const loading_roles = ref(false);
    const loading_perms = ref(false);
    const loading_menus = ref(false);
    const input = ref(InputNewUsuario);
    const usuarioService = new UsuariosService();
    const rolesService = new RolesService();
    const permisoService = new PermisoService();
    const menuService = new MenusService();
    const validaciones = new Validaciones();
    const menus = ref([]);
    const menus_sel = ref([]);
    const roles = ref([]);
    const roles_sel = ref([]);
    const permisos = ref([]);
    const permisos_sel = ref([]);
    const foto_file = ref();
    const foto_64 = ref('');

    const open = (id = null) => {
      alert.value = true;
      roles_sel.value = [];
      permisos_sel.value = [];
      menus_sel.value = [];
      foto_file.value = null;
      foto_64.value = '';
      delete input.value.id;
      for (let key in input.value) {
        if (typeof input.value[key] === 'string') {
          input.value[key] = '';
        } else if (Array.isArray(input.value[key])) {
          input.value[key] = [];
        } else if (typeof input.value[key] === 'number') {
          input.value[key] = 0;
        }
      }
      cargarRoles();
      cargarPermisos();
      cargarMenus();
      if (id) getUserById(id);
    };

    const cargarRoles = async () => {
      loading_roles.value = true;
      const res = await rolesService.roles();
      roles.value = res.roles;
      loading_roles.value = false;
    };

    const cargarPermisos = async () => {
      loading_perms.value = true;
      const res = await permisoService.permisos();
      permisos.value = res.permisos;
      loading_perms.value = false;
    };

    const cargarMenus = async () => {
      loading_menus.value = true;
      const res = await menuService.menus();
      menus.value = res.menus;
      loading_menus.value = false;
    };

    const getUserById = async (id) => {
      loading.value = true;
      const res = await usuarioService.usuarioById(id);
      const xroles = res.usuarioById.roles.map((item) => item);
      const xpermisos = res.usuarioById.permisos_sueltos.map((item) => item);
      const xmenus = res.usuarioById.menus.map((item) => item);
      const us = res.usuarioById.usuario;
      Object.entries(us).forEach(([key, value]) => {
        if (value === null) us[key] = '';
      });
      getFoto(us);
      delete us.estado;
      delete us.fecha_registro;
      delete us.fecha_update;
      delete us.last_login;
      delete us.foto_url;
      us.password = '';
      input.value = us;
      roles_sel.value = xroles;
      permisos_sel.value = xpermisos;
      menus_sel.value = xmenus;
      loading.value = false;
    };

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
      const res = await usuarioService.get_imagen(url); 
      if(res && res.get_imagen) foto_64.value = res.get_imagen; 
    }

    const onSubmit = async () => {
      const xroles = roles_sel.value.map((item) => item.nombre);
      const xpermisos = permisos_sel.value.map((item) => item.metodo);
      input.value.roles = xroles;
      input.value.permisos_sueltos = xpermisos;
      if (xroles.length == 0 && xpermisos.length == 0) {
        Notify.create({
          message: 'Seleccione al menos un rol o un permiso',
          color: 'red',
        });
        return;
      }

      if (input.value.id) actualizar();
      else registrar();
    };

    const registrar = async () => {
      loading.value = true;
      const res = await usuarioService.createUsuario(input.value);
      loading.value = false;
      if (res.createUsuario) {
        const xmenus = menus_sel.value.map((item) => item.id);
        const idsSinDuplicados = [...new Set(xmenus)];
        const datos = {
          user_id: res.createUsuario.id,
          menus: idsSinDuplicados,
        };
        await menuService.asignarMenusUsuario(datos);

        cerrar();
        vue.emit('success');
      }
    };

    const actualizar = async () => {
      loading.value = true;
      const res = await usuarioService.updateUsuario(input.value);
      loading.value = false;
      if (res.updateUsuario) {
        const xmenus = menus_sel.value.map((item) => item.id);
        const idsSinDuplicados = [...new Set(xmenus)];
        const datos = {
          user_id: res.updateUsuario.id,
          menus: idsSinDuplicados,
        };
        await menuService.asignarMenusUsuario(datos);

        cerrar();
        vue.emit('success');
      }
    };

    const cerrar = () => {
      alert.value = false;
    };

    return {
      alert,
      loading,
      loading_roles,
      loading_perms,
      loading_menus,
      input,
      menus,
      menus_sel,
      roles,
      roles_sel,
      permisos,
      permisos_sel,
      foto_file,
      foto_64,
      filevalue,
      onRejected,
      checkClickSession: click.setup().checkClickSession,
      open,
      onSubmit,
      cerrar,
      columnas_rols,
      columnas_perm,
      columnas_menu,
      validaciones,
    };
  },
};
</script>
