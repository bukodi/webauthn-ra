<template>

  <div>
    <TopToolbar></TopToolbar>
  </div>

</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import newsService from '../services/newsService';
import TopToolbar from '../components/TopToolbar.vue';

@Component({
  components: {
    TopToolbar
  }
})
export default class Register extends Vue {
  mounted () {
    navigator.credentials.create({
      publicKey: {
        rp: {
          name: 'localhost'
        },
        user: {
          id: new Uint8Array([1, 2]),
          name: 'test',
          displayName: 'test'
        },
        challenge: new Uint8Array([21, 31]),
        pubKeyCredParams: [
          // See: https://www.w3.org/TR/webauthn-3/#typedefdef-cosealgorithmidentifier
          {
            type: 'public-key',
            alg: -7 // ES256
          },
          {
            type: 'public-key',
            alg: -257 // RS256
          }
        ],
        attestation: 'indirect',
        authenticatorSelection: {
          authenticatorAttachment: 'cross-platform'
        }
      }
    }).then(value => {
      if (value == null) {
        console.log('null returned.');
        return;
      }
      this.dumpCread(value);
    }).catch(error => {
      console.log('rejected', error);
    });
  }

  dumpCread (cred: Credential) {
    const pubKeyCred = cred as PublicKeyCredential;
    console.log('pubKeyCred.id=', pubKeyCred.id);
    console.log('pubKeyCred.type=', pubKeyCred.type);
    console.log('pubKeyCred.rawId=', pubKeyCred.rawId);
    const attestResp = pubKeyCred.response as AuthenticatorAttestationResponse;
    console.log('authResp.attestationObject=', attestResp.attestationObject);
    console.log('authResp.clientDataJSON=', attestResp.clientDataJSON);
    const assertResp = pubKeyCred.response as AuthenticatorAssertionResponse;
    console.log('assertResp.signature=', assertResp.signature);
    console.log('assertResp.userHandle=', assertResp.userHandle);

    newsService.registerAuthenticator(pubKeyCred).then(value => {
      console.log('registerAuthenticator returned=', value);
    }).catch(reason => {
      console.log('registerAuthenticator error=', reason);
    });
  }
}
</script>
