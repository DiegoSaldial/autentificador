import { api } from 'src/boot/axios';

export const clima = async () => {
  const point = 'https://ipinfo.io/json?token=6358e4f96808e5';
  const myip = await api.get(point);
  // console.log('ip', myip.data);

  const geo = (myip.data.loc + '').split(',');
  const lat = geo[0];
  const lon = geo[1];

  const url = `https://api.weatherbit.io/v2.0/forecast/daily?lat=${lat}&lon=${lon}&days=1&lang=es&key=1596b98bdbdb4a1d8c78270f4080b4c2`;
  const res = await api.get(url);
  // console.log('clima::::', res.data);

  const data = {
    ciudad: myip.data.city,
    pais: myip.data.country,
    timezone: myip.data.timezone,
    org: myip.data.org,
    loc: myip.data.loc,
    ip: myip.data.ip,
    temp_max: res.data.data[0].high_temp,
    temp_min: res.data.data[0].low_temp,
    temp_now: res.data.data[0].temp,
  };
  console.log('clima::', data);

  return data;
};
