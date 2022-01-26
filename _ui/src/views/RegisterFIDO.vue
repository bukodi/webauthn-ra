<template>
  <div>
    <TopToolbar></TopToolbar>
    <v-container fill-height>
      <v-row justify="center" align="center">
        <v-col cols="12" sm="4">
          <v-card dark class="mx-auto" color="primary" >
            <v-card-title>
              <v-icon large right>settings_remote</v-icon>
              <span class="font-weight-bold">Roaming authenticator</span>
            </v-card-title>

            <v-card-text>
              USB tokens like Yubikey or Titan security keys
            </v-card-text>

            <v-card-actions>
              <v-btn @click="startRegister('cross-platform')">Register</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TopToolbar from '../components/TopToolbar.vue';
import webauthnService from '@/services/webauthnService';

@Component({
  components: {
    TopToolbar
  }
})
export default class RegisterFIDO extends Vue {
  mounted () {
    console.log('TODO: testing available authenticators');
    if (window.PublicKeyCredential) {
      console.log('Webauthn supported');
      // eslint-disable-next-line no-undef
      PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable().then(available => console.log('User verification supported: ', available));
    } else {
      console.log('Webauthn NOT supported');
    }
    this.startRegister('cross-platform');
  }

  startRegister (type: string) {
    const aa = type as AuthenticatorAttachment;
    console.log('register clicked:' + aa);
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

    webauthnService.attestationResult(pubKeyCred).then(value => {
      console.log('registerAuthenticator returned=', value);
    }).catch(reason => {
      console.log('registerAuthenticator error=', reason);
    });
  }
}
</script>
