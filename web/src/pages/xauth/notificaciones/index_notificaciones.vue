<template>
  <div></div>
</template>

<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script setup lang="ts">
import { onMounted, onUnmounted, ref, watch } from 'vue';
import { subs } from 'src/stores/serverws';
import { useLoginStore } from 'stores/login-store';
import gql from 'graphql-tag';
import { Notify } from 'quasar';

const msg = ref();
const subscription2 = ref();
const useLogin = useLoginStore();

watch(
  () => useLogin.token,
  () => {
    // console.log('::>>>>', subscription2.value);

    if (useLogin.token) notificaciones_subs();
  }
);

const notificaciones_subs = async () => {
  const nots = process.env.NOTIFICACIONES_SUBS;
  if (!nots) return;
  if (!useLogin.token) return;

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

  if (subscription2.value) {
    subscription2.value.unsubscribe();
  }
  subscription2.value = a;
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
