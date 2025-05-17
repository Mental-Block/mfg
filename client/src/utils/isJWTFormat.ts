export function isJWTFormat(token: string | undefined) {
  if (typeof token !== 'string') {
    return false;
  }

  const parts = token.split('.');

  return parts.length === 3;
}
