/* eslint-disable @typescript-eslint/no-explicit-any */
import { useMutation, useQuery } from '@vue/apollo-composable';
import { OperationVariables } from '@apollo/client/core/index.js';
import { DocumentNode } from 'graphql';
import { watch } from 'vue';
import { Notify } from 'quasar'

export interface Opciones {
  showNotyError: boolean
}

const defectoOps: Opciones = {
  showNotyError: true
}

export function mutar(sql:DocumentNode,variables: OperationVariables = {}, opciones: Opciones = defectoOps) {

  return new Promise((resolve, reject) => {
    const {mutate:envio,onDone, error} = useMutation(sql, { variables: { ...variables } })

    watch(error, async ()=>{
      reject(error.value)
      if(opciones.showNotyError){
        mostrarNotifyError(error.value?.message)
      }
    });

    watch(envio, async()=> {
      //
    })

    onDone((res) => {
      //resolve(data);

      // esto se usa en el proyecto academico, probar si conviene
      // if(res.errors) reject(res.errors)
      // esto se usa en el proyecto academico, probar si conviene
      if(res.errors) {
        if(opciones.showNotyError) parseErrors(res.errors);
        reject(res.errors)
      }
      else resolve(res.data);
    });
  });

}

export function query(sql:DocumentNode,variables: OperationVariables = {},opciones: Opciones = defectoOps){

  return new Promise((resolve, reject) => {

    const { error, result } = useQuery(sql, variables, {
      fetchPolicy: 'no-cache',
    });

    watch(error, async ()=>{
      reject(error.value)
      if(opciones.showNotyError){
        mostrarNotifyError(error.value?.message)
      }
    });

    watch(result, async()=> {
      resolve(result.value);
    })

  });

}

function parseErrors(lista:any) {
  const errores = lista.map((m:any)=>m.message);
  mostrarNotifyError(errores);
}

function mostrarNotifyError(mensaje:string | undefined){
  Notify.create({
    color: 'red-5',
    textColor: 'white',
    icon: 'warning',
    message: mensaje || '',
  })
}
