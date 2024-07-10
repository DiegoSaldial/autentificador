<template>
  <div></div>
</template>

<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue';
import { subs } from 'src/stores/serverws';
import gql from 'graphql-tag';
import { Notify } from 'quasar';

const msg = ref();
const subscription2 = ref();

const notificaciones_subs = async () => {
  const nots = process.env.NOTIFICACIONES_SUBS;
  // console.log('log', nots);
  if (!nots) return;

  const sql = gql`
    subscription notificaciones_subs {
      notificaciones_subs {
        title
        data_json
      }
    }
  `;

  const a = subs(sql).subscribe({
    next(result: any) {
      console.log(result.data.notificaciones_subs);
      msg.value = result.data.notificaciones_subs.title;
      Notify.create({
        message: msg.value,
        position: 'top-right',
        color: 'orange',
      });
    },
    error(err) {
      msg.value = err;
    },
  });

  // console.log('aaa', a);
  subscription2.value = a;
  // console.log('aaa', subscription2.value);
};

onMounted(() => {
  notificaciones_subs();
});

onUnmounted(() => {
  console.log('salir');

  if (subscription2.value) {
    subscription2.value.unsubscribe();
    console.log('cerrando', subscription2.value);
  }
});
</script>
