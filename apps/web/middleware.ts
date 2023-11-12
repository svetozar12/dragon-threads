import { NextRequest, NextResponse } from 'next/server';
import { isAuth } from './utils/auth';

export async function middleware(req: NextRequest) {
  const userIsAuthenticated = await isAuth();
  const {
    nextUrl: { clone },
  } = req;
  const { next, redirect } = NextResponse;
  const url = clone();
  if (!userIsAuthenticated) {
    // prevents infinite redirects
    if (url.pathname === '/login') return next();
    url.pathname = '/login';
    return redirect(url);
  }

  return next();
}
