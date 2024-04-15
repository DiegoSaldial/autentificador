import { useLoginStore } from 'src/stores/login-store';
import { jwtDecode } from 'jwt-decode';

export const setTimeLabel = (token: string, refreshToken: string) => {
  const exp = getTimeExp(token);
  const expTotal = getTimeExp(refreshToken);

  const { hours: a, minutes: b, seconds: c } = getTimeSession(exp);
  const { hours, minutes, seconds } = getTimeSession(expTotal);

  let tiempo_end = '';
  if (hours.includes('-') && minutes.includes('-') && seconds.includes('-')) {
    tiempo_end = `${a}:${b}:${c} <br/> ${a}:${b}:${c}`;
  } else {
    tiempo_end = `${a}:${b}:${c} <br/> ${hours}:${minutes}:${seconds}`;
  }
  const store = useLoginStore();
  store.setTiempoSession(tiempo_end);
};

const getTimeSession = (timeRemaining: number) => {
  const hours = (Math.floor(timeRemaining / 3600) + '').padStart(2, '0');
  const min = Math.floor((timeRemaining % 3600) / 60);
  const minutes = (min + '').padStart(2, '0');
  const seconds = (Math.floor(timeRemaining % 60) + '').padStart(2, '0');
  const sesion = {
    hours: hours,
    minutes: minutes,
    seconds: seconds,
  };
  return sesion;
};

export const getTimeExp = (token: string) => {
  const decodedToken = jwtDecode(token);
  const currentTime = Date.now() / 1000;
  const expirationTime = decodedToken.exp || 0;
  const timeRemaining = expirationTime - currentTime;
  return timeRemaining;
};
