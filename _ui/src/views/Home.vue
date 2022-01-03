<template>

  <div>

    <v-navigation-drawer
      app
      fixed
      v-model="showMenu"

    >
      <v-list dense>
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
      height="200"
      color="primary"
    >
      <v-app-bar-nav-icon @click.stop="toggleMenu"></v-app-bar-nav-icon>

      <v-toolbar-title>{{ title }}</v-toolbar-title>

    </v-app-bar>

    <v-btn value="top" :to="{ name: 'top-stories'}">
      <span>Top Stories</span>
      <v-icon>thumb_up</v-icon>
    </v-btn>

    <v-btn value="code" :to="{ name: 'code-examples'}">
      <span>Code Examples</span>
      <v-icon>code</v-icon>
    </v-btn>

    <v-btn value="favorites" :to="{ name: 'my-favorites'}">
      <span>Favorites</span>
      <v-icon>favorite</v-icon>
    </v-btn>
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
