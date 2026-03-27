// types.ts

import { BaseError } from './base-error';

export enum ErrorCode {
  INTERNAL_SERVER_ERROR = 'internal-server-error',
  INVALID_REQUEST = 'invalid-request',
  INVALID_CREDENTIALS = 'invalid-credentials',
  NOT_FOUND = 'not-found',
}

export enum RequestMethods {
  GET = 'get',
  POST = 'post',
  PUT = 'put',
  DELETE = 'delete',
}

export enum AuthMethod {
  BASIC = 'basic',
  JWT = 'jwt',
}

export interface AuthenticatedRequest<T> {
  user: T;
  method: RequestMethods;
}

export interface RequestError {
  code: ErrorCode;
  message: string;
  details?: unknown;
}

export interface RequestContext {
  method: RequestMethods;
  headers: { [key: string]: string };
  params: { [key: string]: string };
}

export interface ValidatedRequest<T> {
  method: RequestMethods;
  headers: { [key: string]: string };
  params: T;
}

export class AuthGatewayError extends BaseError {
  code: ErrorCode;
  constructor(code: ErrorCode, message: string) {
    super(message);
    this.code = code;
  }
}