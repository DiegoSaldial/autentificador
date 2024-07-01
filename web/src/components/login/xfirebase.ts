import { initializeApp } from 'firebase/app';
import { getAuth, GoogleAuthProvider, signInWithPopup } from 'firebase/auth';

const firebaseConfig = initializeApp({
  apiKey: 'AIzaSyD2muOLjhasMAyrYPsn-jNkNMCmmIdP86A',
  authDomain: 'autentificador-94.firebaseapp.com',
  projectId: 'autentificador-94',
  storageBucket: 'autentificador-94.appspot.com',
  messagingSenderId: '433905422324',
  appId: '1:433905422324:web:0b7ff6acc980aac300f3a1',
  measurementId: 'G-MXCD201ZJV',
});

export const auth = getAuth(firebaseConfig);

auth.languageCode = 'es';
const provider = new GoogleAuthProvider();

export const loginGoogle = async () => {
  const res = await signInWithPopup(auth, provider)
    .then((d) => d)
    .catch((e) => e);

  console.log('resss', res);

  if (res && res.user) {
    const result = res;
    const credential = GoogleAuthProvider.credentialFromResult(result);
    let token = null;
    if (credential) token = credential.accessToken;

    const user = result.user;
    console.log('USERRR', user);
    console.log('TOKK', token);

    const forCreate = {
      nombres: user.displayName,
      celular: user.phoneNumber,
      correo: user.email,
      username: user.uid,
      password: user.uid,
    };

    return forCreate;
  } else {
    const error = res;
    const errorCode = error.code;
    const errorMessage = error.message;
    const email = error.customData.email;
    const credential = GoogleAuthProvider.credentialFromError(error);

    console.log('rrrrrrrr', errorCode);
    console.log('rrrrrrrr', credential);
    console.log('rrrrrrrr', errorMessage);
    console.log('rrrrrrrr', email);
  }
  return null;
};
