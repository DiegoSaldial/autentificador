<template>
  <q-dialog v-model="alert" persistent square>
    <q-card
      flat
      bordered
      v-on:click="checkClickSession()"
      style="min-width: 300px; width: 500px"
    >
      <q-card-section class="q-pb-none">
        <div class="text-h6 q-my-none">Ubicacion geografica</div>
      </q-card-section>

      <q-card-section class="q-py-none">
        <GoogleMap
          :api-key="apikey"
          mapId="DEMO_MAP_ID"
          style="width: 100%; height: 500px"
          :center="center"
          :zoom="15"
        >
          <AdvancedMarker
            v-on:dragend="select"
            :options="markerOptions"
            :pin-options="pinOptions"
          />
        </GoogleMap>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <div class="q-mt-md" :align="'right'">
          <q-linear-progress
            v-if="loading"
            dark
            rounded
            indeterminate
            color="secondary"
            class="q-mb-sm"
          />
          <q-btn
            label="limpiar"
            color="red"
            icon="close"
            square
            flat
            @click="limpiar()"
          />
          <q-btn
            label="cerrar"
            color="green"
            icon="check"
            square
            flat
            @click="cerrar()"
          />
        </div>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script>
import { ref } from 'vue';
import { Notify } from 'quasar';
import click from 'src/shared/session';
import { GoogleMap, AdvancedMarker } from 'vue3-google-map';

export default {
  components: { GoogleMap, AdvancedMarker },
  setup(_, { emit }) {
    const alert = ref(false);
    const loading = ref(false);
    const apikey = ref('');
    const center = ref({ lat: -21.53549, lng: -64.72956 });
    const markerOptions = ref({
      position: center,
      title: 'LADY LIBERTY',
      gmpDraggable: true,
    });
    const pinOptions = ref({
      background: '#009800',
      borderColor: '#005000',
      glyphColor: '#dedede',
    });

    const open = (lat = null, lon = null) => {
      apikey.value = process.env.API_MAPS;
      alert.value = true;
      console.log('ssss', lat, lon);

      if (lat && lon) {
        center.value.lat = lat;
        center.value.lng = lon;
      }
    };

    const select = (p) => {
      emit('onpin', p.latLng.lat(), p.latLng.lng());
      Notify.create({
        message: 'Aceptado',
        color: 'orange',
        timeout: 400,
        position: 'top',
        icon: 'fmd_good',
      });
    };

    const limpiar = () => {
      emit('onpin', null, null);
      cerrar();
    };

    const cerrar = () => (alert.value = false);

    return {
      alert,
      loading,
      apikey,
      center,
      markerOptions,
      pinOptions,
      checkClickSession: click.setup().checkClickSession,
      open,
      cerrar,
      select,
      limpiar,
    };
  },
};
</script>
