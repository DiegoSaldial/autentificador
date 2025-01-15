<template>
  <div class="q-pa-sm">
    <!-- <h6 class="q-my-sm text-center">Usuarios del sistema</h6> -->

    <q-table color="primary" :rows="rows" :columns="columns" row-key="id" dense hide-pagination :rows-per-page-options="[0]" :visible-columns="visibleColumns" :filter="filter" :loading="loading" >
      <template v-slot:top-left>
        <h6 class="q-my-none">Usuarios del sistema</h6> 
        <q-toggle v-model="more_datos" @update:model-value="columnas()" color="orange" label="mostrar mas columnas" class="q-my-none" />
      </template>

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model.trim="filter" placeholder="buscar ..." class="q-mx-none" >
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
        <q-select outlined dense v-model="rol_select" :options="roles" label="Filtro de rol" option-value="nombre" option-label="nombre" :loading="loading_roles" :disable="loading" clearable />
        <q-btn label="Registrar" outline color="secondary" icon="person_add" class="q-ml-xs" square @click="registrar()" />
      </template>

      <template v-slot:loading>
        <q-inner-loading showing color="primary" />
      </template>

      <template v-slot:body-cell-index="props">
        <q-td :props="props" :title="'ID: ' + props.row.id">
          {{ props.rowIndex + 1 }}
        </q-td>
      </template>

      <template v-slot:body-cell-foto_url="props">
        <q-td :props="props">
          <img v-if="props.row.foto64" :src="props.row.foto64" alt="perfil" style="max-width: 30px" />
        </q-td>
      </template>

      <template v-slot:body-cell-nombres="props">
        <q-td :props="props" :title="'ID: ' + props.row.id">
          {{ props.row.nombres }}
        </q-td>
      </template>
      <template v-slot:body-cell-fecha_registro="props">
        <q-td :props="props">
          {{ parseFecha(props.row.fecha_registro, true) }}
        </q-td>
      </template>

      <template v-slot:body-cell-fecha_update="props">
        <q-td :props="props">
          {{ parseFecha(props.row.fecha_update, true) }}
        </q-td>
      </template>

      <template v-slot:body-cell-last_login="props">
        <q-td :props="props">
          {{ parseFecha(props.row.last_login, true) }}
        </q-td>
      </template>

      <template v-slot:body-cell-estado="props">
        <q-td :props="props">
          {{ props.row.estado ? 'Activo' : 'Inactivo' }}
        </q-td>
      </template>

      <template v-slot:body-cell-conexiones="props">
        <q-td :props="props">
          <q-badge color="green" v-if="props.row.conexiones > 0">
            {{ props.row.conexiones }}
            <q-tooltip class="bg-purple">
              <span>Conexiones WebSockets</span>
            </q-tooltip>
          </q-badge>
          <span v-else> {{ props.row.conexiones }} </span>
        </q-td>
      </template>

      <template v-slot:body-cell-opt="props">
        <q-td :props="props">
          <q-btn color="green-10" dense square flat icon="more_vert" size="small" >
            <q-menu anchor="top right" self="top left">
              <q-list style="min-width: 110px">
                <q-item clickable v-ripple @click="visualizar(props.row)">
                  <q-item-section avatar>
                    <q-icon color="accent" name="visibility" right class="q-px-none" />
                  </q-item-section>
                  <q-item-section> Ver </q-item-section>
                </q-item>
                <q-item clickable v-ripple @click="actualizar(props.row)">
                  <q-item-section avatar>
                    <q-icon color="accent" name="edit" right class="q-px-none" />
                  </q-item-section>
                  <q-item-section> Editar </q-item-section>
                </q-item>
              </q-list>
            </q-menu>
          </q-btn>
        </q-td>
      </template>
    </q-table>

    <Registrar ref="refRegistrar" v-on:success="listar" />
    <Ver ref="refVer" />
  </div>
</template>

<script>
import { onMounted, ref, watch } from 'vue';
import { parseFecha } from 'stores/utils';
import { useRoute, useRouter } from 'vue-router';
import { columns } from './utils';
import UsuariosService from './usuariosService';
import RolesService from 'pages/xauth/roles/rolesService';
import Registrar from './registrar_usuario.vue';
import Ver from './ver-usuario.vue';

export default {
  components: { Registrar, Ver },
  setup() {
    const usuariosService = new UsuariosService();
    const rolesService = new RolesService();
    const rows = ref([]);
    const loading = ref(false);
    const filter = ref('');
    const refRegistrar = ref();
    const refVer = ref();
    const more_datos = ref(false);
    const visibleColumns = ref([]);
    const roles = ref([]);
    const loading_roles = ref(false);
    const rol_select = ref(null);
    const router = useRouter();
    const route = useRoute();

    const listar = async () => {
      loading.value = true;
      rows.value = [];
      const query = {
        rol: route.query.rol ? route.query.rol : null,
      };
      const res = await usuariosService.usuarios(query);
      if (res.usuarios) {
        const us = res.usuarios;
        for (let i = 0; i < us.length; i++) {
          us[i].foto64 = '';
        }
        rows.value = us;
        cargarImagenes();
      }
      loading.value = false;
    };

    const cargarImagenes = async () => {
      for (let i = 0; i < rows.value.length; i++) {
        getFoto(rows.value[i]);
      }
    };

    const getFoto = async (us) => {
      const url = us.foto_url;
      if (!url) return url;
      const res = await usuariosService.get_imagen(url);
      if (res && res.get_imagen) us.foto64 = res.get_imagen;
    };

    const listarRoles = async () => {
      loading_roles.value = true;
      const res = await rolesService.roles();
      roles.value = res.roles;
      loading_roles.value = false;
      const r = route.query.rol ? route.query.rol : null;
      rol_select.value = roles.value.find((x) => x.nombre == r);
    };

    const registrar = () => refRegistrar.value.open();
    const actualizar = (item) => refRegistrar.value.open(item.id);
    const visualizar = (item) => refVer.value.open(item.id);

    const columnas = () => {
      if (more_datos.value)
        visibleColumns.value = [
          'index',
          'foto_url',
          'nombres',
          'apellido1',
          'apellido2',
          'documento',
          'celular',
          'correo',
          'sexo',
          'direccion',
          'username',
          'last_login',
          'fecha_registro',
          'fecha_update',
          'estado',
          'conexiones',
          'opt',
        ];
      else
        visibleColumns.value = [
          'index',
          'foto_url',
          'nombres',
          'apellido1',
          'apellido2',
          'last_login',
          'conexiones',
          'opt',
        ];
    };

    const setParams = async () => {
      if (rol_select.value)
        await router.push({
          path: '',
          query: { rol: rol_select.value.nombre },
        });
      else await router.push({ path: '', query: null });
      listar();
    };

    watch(
      () => rol_select.value,
      () => setParams()
    );

    onMounted(() => {
      columnas();
      listarRoles();
    });

    return {
      columns,
      rows,
      loading,
      filter,
      refRegistrar,
      refVer,
      visibleColumns,
      more_datos,
      columnas,
      roles,
      loading_roles,
      rol_select,
      listar,
      registrar,
      actualizar,
      visualizar,
      parseFecha,
    };
  },
};
</script>
