import { ArticleType, NewsArticle } from '@/types';

const url = process.env.VUE_APP_SERVER_API_URL + '/articles';

class NewsService {
  getArticlesByType (articleType: ArticleType): Promise<NewsArticle[]> {
    return fetch(url)
      .then((response) => {
        return response.json();
      })
      .then((serverArticles) => {
        const newsArticles = serverArticles
          .filter((serverArticle: any) => serverArticle.articleType === articleType)
          .map(NewsService.map);

        return newsArticles;
      })
      .catch((e) => {
        console.error('An error occurred retrieving the news articles from ' + url, e);
      });
  }

  getFavorites (): Promise<NewsArticle[]> {
    return fetch(url)
      .then((response) => {
        return response.json();
      })
      .then((serverArticles) => {
        const newsArticles = serverArticles
          .filter((serverArticle: any) => serverArticle.isFavourite === true)
          .map(NewsService.map);

        return newsArticles;
      })
      .catch((e) => {
        console.error('An error occurred retrieving the news articles from ' + url, e);
      });
  }

  registerAuthenticator (rawId: string, attestationObject: string, clientDataJSON: string): Promise<any> {
    return fetch(process.env.VUE_APP_SERVER_API_URL + '/authenticator/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        rawId: rawId,
        attestationObject: attestationObject,
        clientDataJSON: clientDataJSON
      })
    })
      .then((response) => {
        return response.json();
      })
      .catch((e) => {
        console.error('An error occurred register authenticator', e);
      });
  }

  private static map (serverArticle: any): NewsArticle {
    return {
      id: serverArticle.id,
      title: serverArticle.title,
      content: serverArticle.content,
      dateString: serverArticle.dateString,
      baseImageName: serverArticle.baseImageName,
      articleType: serverArticle.articleType,
      isFavourite: serverArticle.isFavourite
    } as NewsArticle;
  }
}

export default new NewsService();
