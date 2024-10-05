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
        <LMap
          v-if="showMap"
          ref="refmap"
          :useGlobalLeaflet="false"
          :zoom="zoom"
          :center="center"
          :options="mapOptions"
          style="height: 50vh; width: 40vw"
          @click="handleMapClick"
        >
          <LTileLayer :url="url" :attribution="attribution" />
          <LMarker :lat-lng="markerLatLng"></LMarker>
        </LMap>
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
import click from 'src/shared/session';
import 'leaflet/dist/leaflet.css';
import { latLng } from 'leaflet';
import { LMap, LTileLayer, LMarker } from '@vue-leaflet/vue-leaflet';
import { Notify } from 'quasar';

export default {
  components: { LMap, LMarker, LTileLayer },
  setup(_, { emit }) {
    const alert = ref(false);
    const loading = ref(false);
    const refmap = ref();
    const markerLatLng = ref([]);
    const zoom = ref(14.2);
    const center = ref(latLng(-21.528098, -64.730105));
    const url = ref('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png');
    const attribution = ref(
      '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
    );
    const mapOptions = ref({
      zoomSnap: 0.5,
    });
    const showMap = ref(false);

    const open = (lat = null, lon = null) => {
      alert.value = true;
      console.log('ssss', lat, lon);

      if (lat && lon) {
        markerLatLng.value = [lat, lon];
      }

      setTimeout(() => {
        showMap.value = true;
        console.log('refMap', refmap.value);
      }, 600);
    };

    const handleMapClick = (event) => {
      const { lat, lng } = event.latlng;
      markerLatLng.value = [lat, lng];

      emit('onpin', lat, lng);
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
      zoom,
      center,
      url,
      attribution,
      mapOptions,
      showMap,
      refmap,
      markerLatLng,
      checkClickSession: click.setup().checkClickSession,
      open,
      cerrar,
      handleMapClick,
      limpiar,
    };
  },
};
</script>
