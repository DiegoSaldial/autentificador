<template>
  <q-dialog v-model="alert" persistent square>
    <q-card flat bordered v-on:click="checkClickSession()">
      <q-card-section>
        <div class="text-h6">{{ is_edit ? 'Actualizar' : 'Registrar' }}</div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-form @submit="onSubmit">
          <div class="row q-col-gutter-xs">
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.nombre" label="Nombre" lazy-rules dense :rules="[(val) => validaciones.val_nombre(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.trim="input.descripcion" label="Descripcion" lazy-rules dense counter :rules="[(val) => validaciones.val_descripcion(val)]" />
            </div>
            <div class="col-xs-12 col-sm-6">
              <q-input outlined v-model.number="input.jerarquia" label="Jerarquia" type="number" lazy-rules dense counter :rules="[(val) => validaciones.val_jerarquia(val)]" />
            </div>

            <div class="col-xs-12 col-sm-12">
              <q-table flat color="orange" :loading="loading_perms" title="permisos" hide-pagination :rows-per-page-options="[0]" dense :rows="permisos" :columns="columnas_perm" row-key="metodo" :selected-rows-label="getSelectedString" selection="multiple" v-model:selected="permisos_sel" />
            </div>

            <div class="col-xs-12 col-sm-12">
              <q-table flat color="orange" :loading="loading_menus" title="menus" hide-pagination :rows-per-page-options="[0]" dense :rows="menus" :columns="columnas_menus" row-key="id" :selected-rows-label="getSelectedMenus" selection="multiple" v-model:selected="menus_sel" />
            </div>
          </div>

          <div class="q-mt-md" :align="'right'">
            <q-linear-progress v-if="loading" dark rounded indeterminate color="secondary" class="q-mb-sm" />
            <q-btn :disable="loading" label="cerrar" color="red" icon="close" square flat @click="cerrar()" />
            <q-btn :disable="loading" :label="is_edit ? 'Actualizar' : 'Registrar'" icon="done" type="submit" color="secondary" outline square />
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
import { ref } from 'vue';
import RolesService from 'pages/xauth/roles/rolesService';
import PermisoService from 'pages/xauth/permisos/permisoService';
import MenusService from 'pages/xauth/usuarios/menuService';
import { Notify } from 'quasar';
import { InputNewRol } from './type_roles';
import Validaciones from './validador';
import click from 'src/shared/session';

const columnas_perm = [
  { name: 'nombre', label: 'Nombre', field: 'nombre', align: 'left' },
  { name: 'metodo', label: 'Codigo', field: 'metodo', align: 'left' },
];

const columnas_menus = [
  { name: 'label', label: 'Menu', field: 'label', align: 'left' },
  { name: 'grupo', label: 'Grupo', field: 'grupo', align: 'left' },
];

export default {
  setup(_, vue) {
    const alert = ref(false);
    const is_edit = ref(false);
    const loading = ref(false);
    const loading_perms = ref(false);
    const loading_menus = ref(false);
    const input = ref(InputNewRol);
    const rolesService = new RolesService();
    const permisoService = new PermisoService();
    const validaciones = new Validaciones();
    const menuService = new MenusService();
    const permisos = ref([]);
    const permisos_sel = ref([]);
    const menus = ref([]);
    const menus_sel = ref([]);

    const open = (id = null) => {
      is_edit.value = false;
      alert.value = true;
      permisos_sel.value = [];
      menus_sel.value = [];
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
      cargarPermisos();
      cargarMenus();
      if (id) getRolById(id);
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
    }

    const getRolById = async (id) => {
      is_edit.value = true;
      loading.value = true;
      const res = await rolesService.rolById(id);
      const xpermisos = res.rolById.permisos.map((item) => item);
      const xmenus = res.rolById.menus.map((item) => item);
      const rol = res.rolById;
      delete rol.estado;
      delete rol.fecha_registro;
      delete rol.fecha_update;
      input.value = rol;
      permisos_sel.value = xpermisos;
      menus_sel.value = xmenus;
      loading.value = false;
    };

    const onSubmit = async () => {
      if (!input.value.nombre) return;
      const xpermisos = permisos_sel.value.map((item) => item.metodo);
      const xmenus = menus_sel.value.map((item) => item.id);
      input.value.permisos = xpermisos;
      input.value.menus = xmenus;
      if (xpermisos.length == 0) {
        Notify.create({
          message: 'Seleccione al menos un permiso',
          color: 'negative',
        });
        return;
      }
      if (xmenus.length == 0) {
        Notify.create({
          message: 'Seleccione al menos un menu',
          color: 'negative',
        });
        return;
      }

      if (is_edit.value) actualizar();
      else registrar();
    };

    const registrar = async () => {
      loading.value = true;
      const res = await rolesService.createRol(input.value);
      loading.value = false;
      if (res.createRol) {
        cerrar();
        vue.emit('success');
      }
    };

    const actualizar = async () => {
      loading.value = true;
      const res = await rolesService.updateRol(input.value);
      loading.value = false;
      if (res.updateRol) {
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
      loading_perms,
      loading_menus,
      columnas_perm,
      columnas_menus,
      is_edit,
      input,
      permisos,
      permisos_sel,
      menus,
      menus_sel,
      open,
      onSubmit,
      cerrar,
      validaciones,
      checkClickSession: click.setup().checkClickSession,
      getSelectedString() {
        return `${permisos_sel.value.length} de ${permisos.value.length}`;
      },
      getSelectedMenus() {
        return `${menus_sel.value.length} de ${menus.value.length}`;
      },
    };
  },
};
</script> 
