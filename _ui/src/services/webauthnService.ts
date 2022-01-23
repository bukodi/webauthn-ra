import {ParsedAttestation, ServerPublicKeyCredentialCreationOptionsResponse} from '@/types';

function arrayBufferToBase64 (buffer: ArrayBuffer): string {
  let binary = '';
  const bytes = new Uint8Array(buffer);
  for (let i = 0; i < bytes.byteLength; i++) {
    binary += String.fromCharCode(bytes[i]);
  }
  return window.btoa(binary);
}

class WebauthnService {
  attestationOptions ( type: AuthenticatorAttachment): Promise<PublicKeyCredentialCreationOptions | null> {
    const bodyStr = JSON.stringify({
      username: "johndoe@example.com",
      displayName: "John Doe",
      authenticatorSelection: {
        residentKey: false,
        authenticatorAttachment: "cross-platform",
        userVerification: "preferred"
      },
      attestation: "direct"
    });

    return fetch(process.env.VUE_APP_SERVER_API_URL + '/authenticator/options', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: bodyStr
    })
      .then((response) => {
        return response.json();
      })
      .then((response) => {
        const obj = response as ServerPublicKeyCredentialCreationOptionsResponse;
        return {
        } as PublicKeyCredentialCreationOptions;
      })
      .catch((e) => {
        console.error('An error occurred register authenticator', e);
        return null;
      });
  }

  attestationResult (cred: Credential): Promise<ParsedAttestation> {
    const pubKeyCred = cred as PublicKeyCredential;
    const attestResp = pubKeyCred.response as AuthenticatorAttestationResponse;

    console.log('rawId : ', pubKeyCred.rawId);

    const bodyStr = JSON.stringify({
      id: pubKeyCred.id,
      rawId: arrayBufferToBase64(pubKeyCred.rawId),
      response: {
        attestationObject: arrayBufferToBase64(attestResp.attestationObject),
        clientDataJSON: arrayBufferToBase64(attestResp.clientDataJSON)
      },
      type: pubKeyCred.type
    });
    console.log('Request body: ' + bodyStr);

    return fetch(process.env.VUE_APP_SERVER_API_URL + '/authenticator/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: bodyStr
    })
      .then((response) => {
        return response.json();
      })
      .then((response) => {
        const obj = response as ParsedAttestation;
        return {
          error: obj.error,
          publicKeyAlg: obj.publicKeyAlg,
          publicKeyPEM: obj.publicKeyPEM,
          authenticatorGUID: obj.authenticatorGUID,
          authenticatorType: obj.authenticatorType,
          userPresent: obj.userPresent,
          userVerified: obj.userVerified,
          attestnCertSubjectCN: obj.attestnCertSubjectCN,
          attestnCertIssuerCN: obj.attestnCertIssuerCN
        } as ParsedAttestation;
      })
      .catch((e) => {
        console.error('An error occurred register authenticator', e);
        return {
          error: e,
          userPresent: false,
          userVerified: false
        } as ParsedAttestation;
      });
  }
}

export default new WebauthnService();
