<template>
  <q-page class="row items-center justify-evenly">
    <div>
      <h5>Bienvenido al menu principal</h5>

      <q-btn @click="test()">TEST</q-btn>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { Notify } from 'quasar';
import DemoService from './serverapp';

const demo = new DemoService();

const test = async () => {
  const r = await demo.hola();
  console.log('demo::', r);
  Notify.create({ message: JSON.stringify(r) });
  const s = await demo.mundo('tierra');
  console.log('demo::', s);
  Notify.create({ message: JSON.stringify(s) });
};
</script>
