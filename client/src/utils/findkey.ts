import { ItemType } from 'antd/es/menu/interface';

function flattenHelper(array: any[], result: any[]) {
  for (const obj of array) {
    if (obj?.children) {
      if (obj?.label?.props?.to != undefined) {
        result.push({
          path: obj?.label?.props?.to,
          key: obj?.key,
        });
      }

      flatten(obj?.children);
    } else {
      if (obj?.label?.props?.to != undefined) {
        result.push({
          path: obj?.label?.props?.to,
          key: obj?.key,
        });
      }
    }
  }
}

function flatten(array: ItemType<any>[]): { path: string; key: string }[] {
  let result: any = [];

  for (const obj of array) {
    if (Array.isArray(obj?.children)) {
      if (obj?.label?.props?.to != undefined) {
        result.push({
          path: obj?.label?.props?.to,
          key: obj?.key,
        });
      }

      flattenHelper(obj?.children, result);
    } else {
      if (obj?.label?.props?.to != undefined) {
        result.push({
          path: obj?.label?.props?.to,
          key: obj?.key,
        });
      }
    }
  }

  return result;
}

export function findKey(
  array: ItemType<any>[],
  currentPath: string,
  strict: boolean = true,
  pattern: RegExp | undefined = undefined
): string[] | undefined {
  if (strict === false && pattern == undefined) {
    throw 'if strict is not true you must provide a pattern';
  }

  let arr = flatten(array);
  let temp: string[] = [];

  for (var i = 0; i < arr.length; i++) {
    if (strict === true && currentPath === arr[i].path) {
      return [arr[i].key];
    }

    if (strict === false && pattern != undefined && arr[i].path.match(pattern)) {
      temp.push(arr[i].key);
      continue;
    }
  }

  if (temp.length === 0) {
    return undefined;
  }

  return temp;
}
