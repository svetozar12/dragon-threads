import { Configuration, PostApi, SubDragonApi, UserApi } from './api/v1';

const API_URL = `${process.env['NEXT_PUBLIC_SCHEMA'] || 'http'}://${
  process.env['NEXT_PUBLIC_HOST'] || 'localhost'
}:${process.env['NEXT_PUBLIC_PORT'] || '3333'}`;
const config = new Configuration({ basePath: API_URL });

export const initSdk = (sdkConfig: Configuration) => {
  return new Configuration(sdkConfig);
};

export const setToken = (token: string) => {
  if (!token) return;
  config.accessToken = token;
};

export const sdk = {
  userInstance: () => new UserApi(config),
  postInstance: () => new PostApi(config),
  subDragonInstance: () => new SubDragonApi(config),
};

export * as API from './api/v1';
