/* eslint-disable @typescript-eslint/no-explicit-any */
import gql from 'graphql-tag';
import { mutar, query } from 'stores/server';

export default class MenusService {
  async menus() {
    const sql = gql`
      query menus {
        menus {
          id
          label
          icon
          grupo
          orden
        }
      }
    `;
    return await query(sql, {})
      .then((d: any) => d)
      .catch((e: any) => e);
  }

  async asignarMenusUsuario(input: any) {
    const sql = gql`
      mutation asignarMenusUsuario($input: AsignarMenusUsuario!) {
        asignarMenusUsuario(input: $input)
      }
    `;
    return await mutar(sql, { input: input })
      .then((d: any) => d)
      .catch((e: any) => e);
  }
}
