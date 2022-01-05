<template>

  <div>
    <TopToolbar></TopToolbar>
  </div>

</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import newsService from '../services/newsService';
import { ArticleType, NewsArticle } from '@/types';
import TopToolbar from '../components/TopToolbar.vue';

@Component({
  components: {
    TopToolbar
  }
})
export default class Login extends Vue {
    newsArticles: NewsArticle[] = [];

    mounted () {
      newsService.getArticlesByType(ArticleType.CodeExample)
        .then((newsArticles: NewsArticle[]) => {
          this.newsArticles = newsArticles;
        });
    }
}
</script>
