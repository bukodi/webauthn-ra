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
export default class Register extends Vue {
    newsArticles: NewsArticle[] = [];

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
          ]
        }
      }).then(value => {
        console.log('resolved', value);
      }).catch(error => {
        console.log('rejected', error);
      });
      newsService.getArticlesByType(ArticleType.CodeExample)
        .then((newsArticles: NewsArticle[]) => {
          this.newsArticles = newsArticles;
        });
    }
}
</script>
