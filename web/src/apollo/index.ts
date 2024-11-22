/* eslint-disable @typescript-eslint/no-explicit-any */
import type { ApolloClientOptions } from '@apollo/client/core';
import { createHttpLink, InMemoryCache, split } from '@apollo/client/core';
// import type { BootFileParams } from '@quasar/app-vite'
import { useLoginStore } from 'stores/login-store';
import { setContext } from '@apollo/client/link/context';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import axios from 'axios';
import { createClient } from 'graphql-ws';
import { getTimeExp, setTimeLabel } from 'src/shared/login-time';
import { getMainDefinition } from '@apollo/client/utilities';

const relogin = async (token: string, refreshToken: string) => {
  const query =
    'mutation refreshtoken($token:String!, $refreshToken: String!){ refreshtoken(token:$token, refreshToken:$refreshToken) } ';

  const data = JSON.stringify({
    query: query,
    variables: {
      token: token,
      refreshToken: refreshToken,
    },
  });

  const config = {
    method: 'post',
    url: process.env.GRAPHQL_URI,
    data: data,
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = await axios(config)
    .then(({ data }) => data)
    .catch((e: any) => e);
  if (res && res.data) return res.data.refreshtoken;

  return token;
};

const authLink = setContext(async (_, { headers }) => {
  const store = useLoginStore();
  let token = store.getToken.value || '';

  if (token) {
    const refreshToken = store.getRefreshToken.value || '';
    const exp = getTimeExp(token);
    const expTotal = getTimeExp(refreshToken);

    if (exp < 0 && expTotal < 0) token = refreshToken;
    if (exp < 0 && expTotal > 0) {
      const newToken = await relogin(token, refreshToken);
      token = newToken;
      store.setNewToken(token);
      // console.log('new token.....', token);
    }
    setTimeLabel(token, refreshToken);
  }

  return {
    headers: {
      ...headers,
      authorization: `Bearer ${token}`,
    },
  };
});

export async function getClientOptions() {
  const store = useLoginStore();

  const wsLink = new GraphQLWsLink(
    createClient({
      url: '' + process.env.GRAPHQL_WSS,
      connectionParams: () => {
        const store = useLoginStore();
        const token = store.token;        

        return {
          Authorization: `Bearer ${token}`,
          reconnect: true, // probar si funciona
        };
      },

      lazy: true,
      on: {
        connected: () => { store.setWsNoti('connected'); console.log('Connected to WebSocket') },
        closed: () => { store.setWsNoti('closed'); console.log('WebSocket connection closed') },
        error: (err) => { store.setWsNoti('error'); console.error('WebSocket error:', err) },
        connecting: () => { store.setWsNoti('connecting'); console.log('Reconnecting to WebSocket...') },
      },
      retryAttempts: Infinity,
    })
  );

  const httpLink = authLink.concat(
    createHttpLink({
      uri: process.env.GRAPHQL_URI,
      headers: {
        authorization: `Bearer ${store.token}`,
      },
    })
  );

  const link = split(
    ({ query }) => {
      const definition = getMainDefinition(query);
      return (
        definition.kind === 'OperationDefinition' &&
        definition.operation === 'subscription'
      );
    },
    wsLink,
    httpLink
  );

  ///* {app, router, ...} */ options?: Partial<BootFileParams<any>>
  return <ApolloClientOptions<unknown>>Object.assign(
    // General options.
    <ApolloClientOptions<unknown>>{
      link: link,

      cache: new InMemoryCache(),
      defaultOptions: {
        query: {
          errorPolicy: 'all',
        },
        mutate: {
          errorPolicy: 'all',
        },
      },
    },
    // Specific Quasar mode options.
    process.env.MODE === 'spa'
      ? {
          //
        }
      : {},
    process.env.MODE === 'ssr'
      ? {
          //
        }
      : {},
    process.env.MODE === 'pwa'
      ? {
          //
        }
      : {},
    process.env.MODE === 'bex'
      ? {
          //
        }
      : {},
    process.env.MODE === 'cordova'
      ? {
          //
        }
      : {},
    process.env.MODE === 'capacitor'
      ? {
          //
        }
      : {},
    process.env.MODE === 'electron'
      ? {
          //
        }
      : {},

    // dev/prod options.
    process.env.DEV
      ? {
          //
        }
      : {},
    process.env.PROD
      ? {
          //
        }
      : {},

    // For ssr mode, when on server.
    process.env.MODE === 'ssr' && process.env.SERVER
      ? {
          ssrMode: true,
        }
      : {},
    // For ssr mode, when on client.
    process.env.MODE === 'ssr' && process.env.CLIENT
      ? {
          ssrForceFetchDelay: 100,
        }
      : {}
  );
}
