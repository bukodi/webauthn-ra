import { ParsedAttestation, WebauthnOptionsResponse } from '@/types';
import { arrayBufferToBase64, BaseResponse } from '@/services/servicesBase';

export interface AuthenticatorOptionsRequest {
  authenticatorAttachment: string;
}

export interface AuthenticatorOptionsResponse extends BaseResponse {
  credentialCreationOptions: any;
  fullChallenge: string;
}

class WebauthnService {
  authenticatorOptions (req: AuthenticatorOptionsRequest): Promise<AuthenticatorOptionsResponse> {
    return fetch(process.env.VUE_APP_SERVER_API_URL + '/webauthn/authenticator/options', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(req)
    })
      .then(response => {
        const data = response.json();
        return (data as unknown) as AuthenticatorOptionsResponse;
      })
      .catch(e => {
        throw e;
      });
  }

  authenticatorRegister (cred: Credential): Promise<ParsedAttestation> {
    const pubKeyCred = cred as PublicKeyCredential;
    const attestResp = pubKeyCred.response as AuthenticatorAttestationResponse;

    console.log('pubKeyCred : ', pubKeyCred);
    const aCred: any = cred;
    Object.setPrototypeOf(aCred, PublicKeyCredential);
    console.log('aCred : ', aCred);
    console.log('typeof aCred : ', typeof aCred);
    console.log('pubKeyCred as any : ', JSON.stringify(aCred));

    const bodyStr = JSON.stringify({
      response: {
        id: pubKeyCred.id,
        rawId: pubKeyCred.rawId,
        response: {
          attestationObject: arrayBufferToBase64(attestResp.attestationObject),
          clientDataJSON: arrayBufferToBase64(attestResp.clientDataJSON)
        },
        type: pubKeyCred.type
      },
      fullChallenge: 'cicamica-haj'
    });
    console.log('Request body: ' + bodyStr);

    return fetch(process.env.VUE_APP_SERVER_API_URL + '/webauthn/authenticator/register', {
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
