<template>
  <div class="row justify-center">
    <div class="col-xs-12 col-sm-5">
      <div class="row">
        <div class="col-xs-12">
          <q-card class="q-my-sm manitox" @click="click()">
            <div class="row">
              <div class="col-xs-3">
                <q-img src="https://cdn.quasar.dev/img/parallax2.jpg" />
              </div>
              <div class="col-xs-6 q-pl-sm">
                <p class="q-my-none"><b>Nombre del grupo</b></p>
                <p class="q-my-none">enlaces - Bolivia</p>
              </div>
              <div class="col-xs-2 text-right">
                <p class="q-ma-none"><small>03/01/2024</small></p>
                <q-icon name="pinch"></q-icon> <small>1</small>
              </div>
            </div>
          </q-card>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.manitox:hover {
  cursor: pointer;
}
</style>

<script>
import { onMounted, ref } from 'vue';
import LandingService from './landingService';

export default {
  components: {},
  setup() {
    const loading = ref(false);
    const landingService = new LandingService();

    const click = () => {
      console.log('>>>');
    };

    onMounted(() => {
      console.log('landinggggg', landingService);
    });

    return {
      loading,
      click,
    };
  },
};
</script>
