import { sdk, setToken } from '@dragon-threads/shared/sdk';
import { getCookie } from 'cookies-next';

export async function isAuth() {
  try {
    const token = getCookie('accessToken') || '';
    setToken(token);
    const { status } = await sdk
      .authInstance()
      .v1AuthVerifyGet()
      .catch((err) => err.response);

    return status !== 401;
  } catch (error) {
    return false;
  }
}
