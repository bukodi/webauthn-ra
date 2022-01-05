<template>

  <div>

    <v-app-bar
      app
      fixed
      dark
      height="200"
      color="primary"
    >
      <v-app-bar-title>Webauthn - Registration Authority</v-app-bar-title>
    </v-app-bar>

    <v-item-group active-class="primary">
      <v-container>
        <v-row>
          <v-col cols="12" md="4">
            <v-btn value="register" :to="{ name: 'register'}">
              <span>Register</span>
              <v-icon>person_add</v-icon>
            </v-btn>
          </v-col>
          <v-col cols="12" md="4">
            <v-btn value="login" :to="{ name: 'login'}">
              <span>Log in</span>
              <v-icon>login</v-icon>
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-item-group>

  </div>

</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import NewsList from '../components/NewsList.vue';
import { Getter } from 'vuex-class';

@Component({
  components: {
    NewsList
  }
})
export default class Home extends Vue {
  @Getter('title', { namespace: 'topToolbar' }) title!: string;

  showMenu = false;

  toggleMenu (): void {
    this.showMenu = !this.showMenu;
  }

  doNothing (): void {
    console.info('Not doing anything!');
  }
}
</script>
