export const TITLE_PREFIX = 'MFG - ';

export function isJwtFormat(token: string | undefined) {
  if (typeof token !== 'string') {
    return false;
  }

  const parts = token.split('.');

  return parts.length === 3;
}
