/* eslint-disable @typescript-eslint/no-explicit-any */
import { watch, ref } from 'vue';
import gql from 'graphql-tag';
import { Notify } from 'quasar';
import { subs } from 'src/stores/serverws';
import { useLoginStore } from 'stores/login-store';

const msg = ref('');
const subscriptionRef: any = ref(null);
const useLogin = useLoginStore();

// Define la suscripción GraphQL
const NOTIFICACIONES_SUBS = gql`
  subscription notificaciones_subs {
    notificaciones_subs {
      title
      data_json
    }
  }
`;

// Función para iniciar la suscripción
const iniciarSubscripcion = () => {
  if (subscriptionRef.value) {
    subscriptionRef.value.unsubscribe(); // Cerrar cualquier suscripción activa anterior
  }

  const to = useLogin.token;
  // console.log('>iniciarSubscripcion', to);
  if (!to) return;

  subscriptionRef.value = subs(NOTIFICACIONES_SUBS).subscribe({
    next(result: any) {
      const notificacion = result.data.notificaciones_subs;
      console.log('Notificación recibida:', notificacion);

      // Mostrar la notificación
      Notify.create({
        message: notificacion.title,
        position: 'top-right',
        color: 'green',
        progress: true,
        progressClass: 'bg-white text-white',
      });
    },
    error(err) {
      console.error('Error en la suscripción:', err);
      msg.value = err.message || 'Error en la suscripción';
    },
  });
};

// Observa el token y gestiona la suscripción
watch(
  () => useLogin.token,
  (newToken) => {
    if (newToken) {
      iniciarSubscripcion(); // Inicia la suscripción si hay un token válido
    } else if (subscriptionRef.value) {
      subscriptionRef.value.unsubscribe(); // Cancela la suscripción si el token es inválido
      subscriptionRef.value = null;
    }
  },
  { immediate: true }, // Ejecutar inmediatamente al iniciar
);

// Limpieza de la suscripción (puedes llamarla cuando lo necesites)
const detenerSubscripcion = () => {
  if (subscriptionRef.value) {
    subscriptionRef.value.unsubscribe();
    subscriptionRef.value = null;
    console.log('Suscripción detenida');
  }
};

export { iniciarSubscripcion, detenerSubscripcion };
