/* eslint-disable @typescript-eslint/no-explicit-any */
import { useLoginStore } from 'stores/login-store';
import { jwtDecode } from 'jwt-decode';
import { Quasar, Notify } from 'quasar';
import { setTimeLabel } from 'src/shared/login-time';
import { createApp, h } from 'vue';
import Relogin from 'src/shared/dialog_relogin.vue';

export default {
  name: 'MyLayout',
  components: {},

  setup() {
    const store = useLoginStore();

    const checkClickSession = async () => {
      const refreshToken = store.getRefreshToken.value || '';
      if (!refreshToken) return;
      const decodedToken = jwtDecode(refreshToken);
      const currentTime = Date.now() / 1000;
      const expirationTime = decodedToken.exp || 0;
      const timeRemaining = expirationTime - currentTime;

      setTimeLabel(store.token + '', store.refreshToken + '');
      if (timeRemaining < 0) {
        const r = await mostrarRelogin();
        if (r) return;

        store.setToken('', '');
        Notify.create({
          message: 'Session expirada, por favor vuelva a ingresar.',
          color: 'negative',
        });
      }
    };

    const mostrarRelogin = async () => {
      let r = document.getElementById('relogin');
      if (!r) {
        const container = document.createElement('div');
        container.id = 'relogin';
        document.body.appendChild(container);
        r = container;
      }

      const app = createApp({
        render() {
          return h(Relogin, {
            ref: 'reloginComponent',
          });
        },
      });

      app.use(Quasar, {
        config: {},
        plugins: {},
      });
      const instancia = app.mount(r);
      const d: any = instancia.$refs.reloginComponent;
      const re = await d.openDialog();
      r.remove();
      app.unmount();
      setTimeLabel(store.token + '', store.refreshToken + '');
      return re;
    };

    return {
      checkClickSession,
    };
  },
};
