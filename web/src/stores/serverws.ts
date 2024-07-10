/* eslint-disable @typescript-eslint/no-explicit-any */
import { OperationVariables } from '@apollo/client';
import { useSubscription } from '@vue/apollo-composable';
import { DocumentNode } from 'graphql';
import { watch } from 'vue';
import {
  Opciones,
  defectoOps,
  mostrarNotifyError,
  parseErrors,
} from './server';
import { Observable } from 'rxjs';

export interface Valores {
  data: any;
  stop: () => void;
}

export const subs = (
  sql: DocumentNode,
  variables: OperationVariables = {},
  opciones: Opciones = defectoOps
) => {
  return new Observable((observer) => {
    const {
      error,
      // subscription: envio,
      stop,
      onResult,
    } = useSubscription(sql, variables, {
      fetchPolicy: 'network-only',
      clientId: 'clientews',
    });

    watch(error, () => {
      if (error.value) {
        observer.error(error.value);
        if (opciones.showNotyError) {
          mostrarNotifyError(error.value?.message);
        }
      }
    });

    onResult((res) => {
      if (res.errors) {
        if (opciones.showNotyError) parseErrors(res.errors);
        observer.error(res.errors);
      } else {
        const d: Valores = {
          data: res.data,
          stop: stop,
        };
        observer.next(d);
      }
    });

    return () => {
      stop;
      // Cleanup logic if necessary, such as unsubscribing
      // For example, you can use a flag to stop processing new results
      // or use the `envio` subscription object to unsubscribe if it supports it
    };
  });
};
