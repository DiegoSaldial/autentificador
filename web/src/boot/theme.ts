import { boot } from 'quasar/wrappers';
import { Cookies } from 'quasar';
import { Dark } from 'quasar';
import { watch } from 'vue';

export default boot(({}) => {
  const cookie_name = process.env.COOKIE_THEME_NAME || '';
  const value = Cookies.get(cookie_name);

  Dark.set(value == 'true');

  watch(
    () => Dark.isActive,
    (val) => {
      Dark.set(val);
      Cookies.set(cookie_name, '' + val);
    }
  );
});
