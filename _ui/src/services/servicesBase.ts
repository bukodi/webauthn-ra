export function arrayBufferToBase64 (buffer: ArrayBuffer): string {
  let binary = '';
  const bytes = new Uint8Array(buffer);
  for (let i = 0; i < bytes.byteLength; i++) {
    binary += String.fromCharCode(bytes[i]);
  }
  return window.btoa(binary);
}

export function arrayBufferToBase64Url (buffer: ArrayBuffer): string {
  let binary = '';
  const bytes = new Uint8Array(buffer);
  for (let i = 0; i < bytes.byteLength; i++) {
    binary += String.fromCharCode(bytes[i]);
  }
  let urlString = window.btoa(binary);
  // Replace non-url compatible chars with base64 standard chars
  urlString = urlString.replace(/\+/g, '-').replace(/\//g, '_');
  // Remove padding
  urlString = urlString.replace(/=/g, '');
  return urlString;
}

export function base64urlToArrayBuffer (b64: string): Uint8Array {
  // Replace non-url compatible chars with base64 standard chars
  let input = b64.replace(/-/g, '+').replace(/_/g, '/');

  // Pad out with standard base64 required padding characters
  const pad = input.length % 4;
  if (pad) {
    if (pad === 1) {
      throw new Error('InvalidLengthError: Input base64url string is the wrong length to determine padding');
    }
    input += new Array(5 - pad).join('=');
  }
  const raw = atob(input);
  const array = new Uint8Array(new ArrayBuffer(raw.length));

  for (let i = 0; i < raw.length; i++) {
    array[i] = raw.charCodeAt(i);
  }

  return array;
}

export interface BaseResponse {
  errorId?: string;
  errorMessage?: string;
}
