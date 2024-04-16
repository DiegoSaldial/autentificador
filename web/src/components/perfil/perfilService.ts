/* eslint-disable @typescript-eslint/no-explicit-any */
import gql from 'graphql-tag';
import { query } from 'stores/server'

export default class PerfilService {
  async me() {
    const sql = gql`
      query{
        me(input:{show_roles:true, show_permisos:true}){
          usuario{id nombres apellido1 apellido2 documento celular correo sexo direccion username fecha_registro fecha_update estado}
        }
      }
      `;
      return await query(sql,{}).then(d=>d).catch(e=>e)
  }
  
  async updateUsuarioPerfil(input:any) {
    const sql = gql`
      mutation updateUsuarioPerfil($input: UpdatePerfil!){
        updateUsuarioPerfil(input:$input){id}
      }
      `;
      return await query(sql,{input:input}).then(d=>d).catch(e=>e)
  }
}
