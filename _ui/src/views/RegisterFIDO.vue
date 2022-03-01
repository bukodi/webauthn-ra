<template>
  <div>
    <TopToolbar></TopToolbar>
    <v-container fill-height>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="4">
          <v-card class="mx-auto" color="primary" dark>
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
import webauthnService, { AuthenticatorOptionsRequest } from '@/services/webauthnService';
import { base64ToArrayBuffer } from '@/services/servicesBase';

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

  webauthnAttestationOptions (type: AuthenticatorAttachment): Promise<{ opts: PublicKeyCredentialCreationOptions; fc: string }> {
    const req = {
      authenticatorAttachment: type
    } as AuthenticatorOptionsRequest;

    return webauthnService.authenticatorOptions(req)
      .then(resp => {
        console.log('Before conversion: ', resp);
        const opts = resp.credentialCreationOptions;

        const pkcco: PublicKeyCredentialCreationOptions = {
          challenge: base64ToArrayBuffer(opts.challenge),
          rp: opts.rp as PublicKeyCredentialRpEntity,
          user: {
            id: base64ToArrayBuffer(opts.user.id),
            name: opts.user.name,
            displayName: opts.user.displayName
          } as PublicKeyCredentialUserEntity,
          timeout: opts.timeout,
          authenticatorSelection: opts.authenticatorSelection,
          attestation: opts.attestation,
          pubKeyCredParams: opts.pubKeyCredParams
        };
        console.log('After conversion: ', pkcco);
        return {
          opts: pkcco,
          fc: resp.fullChallenge
        };
      });
  }

  startRegister (type: string) {
    const aa = type as AuthenticatorAttachment;
    console.log('register clicked:' + aa);

    this.webauthnAttestationOptions('cross-platform').then(ret => {
      const opts: CredentialCreationOptions = {
        publicKey: ret.opts
      };
      navigator.credentials.create(opts).then(value => {
        if (value == null) {
          console.log('null returned.');
          return;
        }
        this.dumpCread(value);
      });
    }
    ).catch(err => {
      console.log('rejected', err, err.stack);
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

    webauthnService.authenticatorRegister(pubKeyCred).then(value => {
      console.log('registerAuthenticator returned=', value);
    }).catch(reason => {
      console.log('registerAuthenticator error=', reason);
    });
  }
}

</script>
