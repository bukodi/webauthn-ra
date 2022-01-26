<template>
  <div>
    <TopToolbar></TopToolbar>
    <v-container>
      <v-row>
        <v-col cols="12" sm="4">
          <v-card dark class="mx-auto" color="primary" max-width="400">
            <v-card-title>
              <v-icon large right>settings_remote</v-icon>
              <span class="font-weight-bold">Roaming authenticator</span>
            </v-card-title>

            <v-card-text>
              USB tokens like Yubikey or Titan security keys
            </v-card-text>

            <v-card-actions>
              <v-btn :to="{ name: 'register-fido'}">Register</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
        <v-col cols="12" sm="4">
          <v-card dark class="mx-auto" color="primary" max-width="400">
            <v-card-title>
              <v-icon large right>fingerprint</v-icon>
              <span class="font-weight-bold">Platform authenticator</span>
            </v-card-title>

            <v-card-text>
              Windows Hello or Andoid/IOS device lock
            </v-card-text>

            <v-card-actions>
              <v-btn :to="{ name: 'register-fido'}">Register</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
        <v-col cols="12" sm="4">
          <v-card dark class="mx-auto" color="primary" max-width="400" disabled>
            <v-card-title>
              <v-icon large right>dialpad</v-icon>
              <span class="font-weight-bold">Authenticator emulator</span>
            </v-card-title>

            <v-card-text>
              Webassambley based software authenticator
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

@Component({
  components: {
    TopToolbar
  }
})
export default class Register extends Vue {
  mounted () {
    console.log('TODO: testing avaiable authenticators');
    if (window.PublicKeyCredential) {
      console.log('Webauthn supported');
      // eslint-disable-next-line no-undef
      PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable().then(available => console.log('User verification supported: ', available));
    } else {
      console.log('Webauthn NOT supported');
    }
  }
}
</script>
