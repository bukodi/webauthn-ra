export function arrayBufferToBase64 (buffer: ArrayBuffer): string {
  let binary = '';
  const bytes = new Uint8Array(buffer);
  for (let i = 0; i < bytes.byteLength; i++) {
    binary += String.fromCharCode(bytes[i]);
  }
  return window.btoa(binary);
}

export function base64ToArrayBuffer (b64: string): Uint8Array {
  console.log('Base64Input: >' + b64 + '<');
  const raw1 = decodeURI(b64);
  console.log('URLDecoded: >' + b64 + '<');
  const raw = atob(b64);
  console.log('Raw: ', raw);
  const array = new Uint8Array(new ArrayBuffer(raw.length));
  console.log('Array: ', array);

  for (let i = 0; i < raw.length; i++) {
    array[i] = raw.charCodeAt(i);
  }

  return array;
}

export interface BaseResponse {
  errorId?: string;
  errorMessage?: string;
}
