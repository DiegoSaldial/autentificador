export const parseFecha = (fecha: string, ful: boolean) => {
  if (!fecha) return '?';
  /*const partes = fecha.split('T')
  const fs = partes[0].split('-')
  let date = `${fs[2]}/${fs[1]}/${fs[0]}`
  if(ful){
    const fx = partes[1].split(':')
    date += ` ${fx[0]}:${fx[1]}`
  }
  return date*/

  const host = location.host;
  if (!host.includes('localhost:')) {
    fecha = fecha.replaceAll('-04:00', 'Z');
  }

  const date = new Date(fecha);

  // Formatear la fecha a dd/MM/yyyy HH:mm
  const day = String(date.getDate()).padStart(2, '0');
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const year = date.getFullYear();
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');

  let formattedDate = '';

  if (ful) {
    formattedDate = `${day}/${month}/${year} ${hours}:${minutes}`;
  } else {
    formattedDate = `${day}/${month}/${year}`;
  }
  return formattedDate;
};
