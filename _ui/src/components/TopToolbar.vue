<template>
  <div>

    <v-navigation-drawer
      app
      fixed
      v-model="showMenu"

    >
      <v-list dense>
        <v-list-item :to="{ path: '/'}">
          <v-list-item-action>
            <v-icon>home</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Home</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item @click="doNothing">
          <v-list-item-action>
            <v-icon>settings</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Settings</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item @click="doNothing">
          <v-list-item-action>
            <v-icon>help</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Help</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      app
      fixed
      dark
      color="primary"
    >
      <v-app-bar-nav-icon @click.stop="toggleMenu"></v-app-bar-nav-icon>

      <v-toolbar-title>{{ title }}</v-toolbar-title>

    </v-app-bar>

  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Getter } from 'vuex-class';

@Component
export default class TopToolbar extends Vue {
  @Getter('title', { namespace: 'topToolbar' }) title!: string;

  showMenu = false;

  toggleMenu (): void {
    this.showMenu = !this.showMenu;
  }

  doNothing (): void {
    console.info('Not doing anything!');
  }

  navigateToHome (): void {
    console.info('Not doing anything!');
  }
}
</script>
