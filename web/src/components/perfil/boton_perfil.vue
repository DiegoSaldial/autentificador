<template>
  <div class="q-pa-none q-ma-none">
    <q-btn flat class="q-pl-xs q-pr-none">
      <q-avatar size="26px">
        <img src="https://cdn.quasar.dev/img/boy-avatar.png">
      </q-avatar>
      <span class="q-ml-xs ellipsis" style="max-width: 9em;"> {{ store.dataUser.usuario.username }} </span>
      <q-tooltip> Perfil </q-tooltip>
      <q-menu>
        <div class="row no-wrap q-pa-md">
          <div class="column" style="min-width: 130px;">
            <div class="text-h6 q-mb-md">Mi Perfil</div>
            <p class="q-mt-none q-mb-xs"> {{ datos.usuario }} </p>
            <p class="q-mt-none q-mb-xs"> {{ datos.documento }} </p>
            <p class="q-mt-none q-mb-xs"> {{ datos.celular }} </p>
            <p class="q-mt-none q-mb-xs"> {{ datos.correo }} </p>
            <p class="q-mt-none q-mb-xs"> {{ datos.roles }} </p>
            <q-btn color="purple" label="Editar" icon="edit" square flat size="sm" v-close-popup outline stretch @click="openEdit()" />
          </div>

          <q-separator vertical inset class="q-mx-lg" />

          <div class="column items-center" style="min-width: 99px;">
            <q-avatar size="72px">
              <!-- <img src="https://cdn.quasar.dev/img/avatar4.jpg"> -->
              <q-img v-if="foto_64" :src="foto_64" spinner-color="white" />
            </q-avatar>


            <div class="text-subtitle1 q-mt-md q-mb-xs" style="white-space: nowrap"> {{ datos.usuario }} </div>

            <q-btn color="red" label="Salir" icon="block" square flat size="sm" v-close-popup stretch @click="logout()" />
          </div>
        </div>
      </q-menu>
    </q-btn>

    <EditarPerfil ref="refEditarPerfil"/>
  </div>
</template>


<script>
import { onMounted, ref } from 'vue'
import { useLoginStore } from 'src/stores/login-store';
import PerfilService from './perfilService';
import EditarPerfil from './editar_perfil.vue'
import { useRouter } from 'vue-router';

export default {
  components:{ EditarPerfil },
  setup () {
    const datos = ref({})
    const store = useLoginStore()
    const perfilService = new PerfilService();
    const refEditarPerfil = ref()
    const router = useRouter();
    const foto_64 = ref('');

    const logout = async ()=> {
      await router.push('/');
      store.setToken()
    }

    const openEdit = () => refEditarPerfil.value.open();

    const cargarDatos = () => {
      foto_64.value = '';
      const data = store.dataUser;
      if(typeof(data) == 'string') data = JSON.parse(data);
      if(!data.roles) return;  

      const roles = data.roles.map(x=>x.nombre)
      const us = data.usuario;
      if(!us.apellido2) us.apellido2 = '';
      if(!us.celular) us.celular = '';
      if(!us.documento) us.documento = '';
      if(!us.correo) us.correo = '';
      datos.value.usuario = `${us.nombres} ${us.apellido1} ${us.apellido2}`;
      datos.value.documento = `${us.documento}`;
      datos.value.celular = `${us.celular}`;
      datos.value.correo = `${us.correo}`;
      datos.value.roles = roles.join(', ');
      getFoto(data.usuario)
    }

    const getFoto = async (us) => {
      const url = us.foto_url;
      if(!url) return url;
      const res = await perfilService.get_imagen(url); 
      if(res && res.get_imagen) foto_64.value = res.get_imagen; 
    }

    onMounted(()=>{
      cargarDatos();
    })

    return {
      datos,
      logout,
      openEdit,
      store,
      refEditarPerfil,
      foto_64
    }
  }
}
</script>
