<template>
  <div>
    <TopToolbar></TopToolbar>
    <NewsList :newsArticles="newsArticles"></NewsList>
    <BottomNav></BottomNav>
  </div>

</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import newsService from '../services/newsService';
import NewsList from '../components/NewsList.vue';
import { NewsArticle } from '@/types';
import TopToolbar from '../components/TopToolbar.vue';
import BottomNav from '../components/BottomNav.vue';

@Component({
  components: {
    TopToolbar,
    NewsList,
    BottomNav
  }
})
export default class TopStories extends Vue {
  newsArticles: NewsArticle[] = [];

  mounted () {
    newsService.getFavorites()
      .then((newsArticles: NewsArticle[]) => {
        this.newsArticles = newsArticles;
      });
  }
}
</script>
