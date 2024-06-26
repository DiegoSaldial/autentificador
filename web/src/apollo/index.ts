/* eslint-disable @typescript-eslint/no-explicit-any */
import type { ApolloClientOptions } from '@apollo/client/core';
import { createHttpLink, InMemoryCache } from '@apollo/client/core';
// import type { BootFileParams } from '@quasar/app-vite'
import { useLoginStore } from 'stores/login-store';
import { setContext } from '@apollo/client/link/context';
import axios from 'axios';
import { getTimeExp, setTimeLabel } from 'src/shared/login-time';

export /* async */ function getClientOptions() {
  const store = useLoginStore();

  const authLink = setContext(async (_, { headers }) => {
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

  ///* {app, router, ...} */ options?: Partial<BootFileParams<any>>
  return <ApolloClientOptions<unknown>>Object.assign(
    // General options.
    <ApolloClientOptions<unknown>>{
      link: authLink.concat(
        createHttpLink({
          uri:
            process.env.GRAPHQL_URI ||
            // Change to your graphql endpoint.
            'http://localhost:8020/query',
          headers: {
            authorization: `Bearer ${store.token}`,
          },
        })
      ),

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
