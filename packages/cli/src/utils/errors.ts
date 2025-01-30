/**
 * Thrown when a request to KHULNASOFT API occurs.
 */
export class KHULNASOFTRequestError extends Error {
  constructor(message: any) {
    super(message)
    this.name = 'KHULNASOFTRequestError'
  }
}

export function handleKHULNASOFTRequestError(
  err?: { code: number; message: string },
  errMsg?: string,
) {
  if (!err) {
    return
  }

  let message = ''
  switch (err.code) {
    case 400:
      message = 'bad request'
      break
    case 401:
      message = 'unauthorized'
      break
    case 403:
      message = 'forbidden'
      break
    case 404:
      message = 'not found'
      break
    case 500:
      message = 'internal server error'
      break
  }

  throw new KHULNASOFTRequestError(
    `${errMsg && `${errMsg}: `}[${err.code}] ${message && `${message}: `}${err.message ?? 'no message'
    }`,
  )
}
