import gql from 'graphql-tag';
import { mutar, query } from 'stores/serverapp';

export default class DemoService {
  async hola() {
    const sql = gql`
      query {
        hola
      }
    `;

    return await query(sql, {})
      .then((d) => d)
      .catch((e) => e);
  }

  async mundo(name: string) {
    const sql = gql`
      mutation mundo($name: String!) {
        mundo(name: $name)
      }
    `;

    return await mutar(sql, { name: name })
      .then((d) => d)
      .catch((e) => e);
  }
}
