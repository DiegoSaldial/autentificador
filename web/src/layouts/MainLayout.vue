<template>
  <q-layout view="hHh Lpr lff" container style="height: 100vh" class="shadow-0">
    <q-header class="bg-primary">
      <q-toolbar>
        <q-btn v-if="store.dataUser.usuario" flat dense round @click="toggleLeftDrawer" aria-label="Menu" icon="menu" />

        <q-btn flat no-caps no-wrap class="q-ml-xs" v-if="$q.screen.gt.xs && store.dataUser.usuario" to="/">
          <!-- <q-icon :name="fabYoutube" color="red" size="28px" /> -->
          <q-toolbar-title shrink class="text-weight-bold">
            Auth
            <q-badge rounded :color="colorWs(store.ws_noti_status)">
              <q-tooltip class="bg-purple">
                <span>Estado de notificaciones</span>
              </q-tooltip>
            </q-badge>
          </q-toolbar-title>
        </q-btn>

        <q-space />
        <span v-if="storeClima.getClima && storeClima.getClima.ciudad">
          {{ storeClima.getClima }}
        </span>
        <q-toggle v-model="$q.dark.isActive" color="white" />
        <small v-if="store.dataUser.usuario && show_time" v-html="store.tiempoSession"> </small>
        <BtnPerfil v-if="store.dataUser.usuario"/>
      </q-toolbar>
    </q-header>

    <q-drawer v-if="store.dataUser.usuario" v-model="leftDrawerOpen" show-if-above bordered :breakpoint="700" :width="240" >
      <q-scroll-area class="fit">
        <q-list padding>

          <q-item-label header class="text-weight-bold text-uppercase">
            Menu
          </q-item-label>

          <template v-for="t in store.menus">
            <template v-for="(link,i) in t" :key="link.text" >
              <q-item v-ripple clickable :to="link.path" active-class="text-purple">
                <q-item-section avatar >
                  <q-icon :color="link.color" :name="link.icon" />
                </q-item-section>
                <q-item-section >
                  <q-item-label>{{ link.label }}</q-item-label>
                </q-item-section>
              </q-item>

              <q-separator v-if="i+1==t.length" class="q-my-md" />
            </template>
          </template>

        </q-list>
      </q-scroll-area>
    </q-drawer>

    <Notificaciones ref="refNotificaciones" />

    <q-page-container>
      <LoginView v-if="!store.dataUser.usuario" />
      <router-view v-else />
    </q-page-container>
  </q-layout>
</template>

<script>
import { onMounted, ref } from 'vue'
import { fabYoutube } from '@quasar/extras/fontawesome-v6'
import {useLoginStore} from 'stores/login-store'
import {useClimaStore} from 'stores/clima-store'
import LoginView from 'components/login/LoginView.vue'
import BtnPerfil from 'components/perfil/boton_perfil.vue'
import Notificaciones from 'pages/xauth/notificaciones/index_notificaciones.vue'

export default {
  name: 'MyLayout',
  components:{ LoginView,BtnPerfil,Notificaciones },

  setup () {
    const leftDrawerOpen = ref(false)
    const store = useLoginStore();
    const storeClima = useClimaStore();
    const show_time = ref(process.env.SHOW_TIME_LABEL)
    const refNotificaciones = ref()

    const toggleLeftDrawer = () => leftDrawerOpen.value = !leftDrawerOpen.value

    const colorWs = (str) => {
      const d = {
        'connected':'green',
        'closed':'negative',
        'error':'red',
        'connecting':'orange',
      }
      return d[str];
    }

    onMounted(()=>{
      storeClima.setearClima(); 
    })

    return {
      fabYoutube,
      leftDrawerOpen,
      storeClima,
      store,
      show_time,
      refNotificaciones,
      colorWs,
      toggleLeftDrawer,

      links3: [
        { icon: fabYoutube, text: 'Usuarios', path:'/usuarios' },
        { icon: 'local_movies', text: 'Roles', path:'/roles' },
      ],
      links4: [
        { icon: 'settings', text: 'Settings' },
      ],
    }
  }
}
</script>
