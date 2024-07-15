/* eslint-disable @typescript-eslint/no-explicit-any */
import gql from 'graphql-tag';
import { mutar } from 'stores/server';

export default class LoginService {
  async login(username: string, password: string, external: boolean) {
    const sql = gql`
      mutation login($input: NewLogin!) {
        login(input: $input) {
          token
          refreshToken
        }
      }
    `;

    const v = {
      input: { username: username, password: password, external: external },
    };
    return await mutar(sql, v)
      .then((d) => d)
      .catch((e) => e);
  }

  async createOauth(input: any) {
    const sql = gql`
      mutation createOauth($input: NewUsuarioOauth!) {
        createOauth(input: $input) {
          id
        }
      }
    `;

    const v = { input: input };
    return await mutar(sql, v)
      .then((d) => d)
      .catch((e) => e);
  }
}
