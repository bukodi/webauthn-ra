import { ArticleType, NewsArticle, ParsedAttestation } from '@/types';

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

  registerAuthenticator (cred: Credential): Promise<ParsedAttestation> {
    const pubKeyCred = cred as PublicKeyCredential;
    console.log('pubKeyCred.id=', pubKeyCred.id);
    console.log('pubKeyCred.type=', pubKeyCred.type);
    console.log('pubKeyCred.rawId=', pubKeyCred.rawId);
    const attestResp = pubKeyCred.response as AuthenticatorAttestationResponse;
    console.log('authResp.attestationObject=', attestResp.attestationObject);
    console.log('authResp.clientDataJSON=', attestResp.clientDataJSON);
    const assertResp = pubKeyCred.response as AuthenticatorAssertionResponse;
    return fetch(process.env.VUE_APP_SERVER_API_URL + '/authenticator/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id: pubKeyCred.id,
        rawId: pubKeyCred.rawId,
        response: {
          attestationObject: attestResp.attestationObject,
          clientDataJSON: attestResp.clientDataJSON
        },
        type: pubKeyCred.type
      })
    })
      .then((response) => {
        return response.json();
      })
      .then((response) => {
        const obj = response.json();
        return (obj as ParsedAttestation);
        return {
          error: obj.error,
          publicKeyAlg: obj.publicKeyAlg,
          publicKeyPEM: obj.publicKeyPEM,
          authenticatorGUID: obj.authenticatorGUID,
          authenticatorType: obj.authenticatorType,
          userPresent: obj.userPresent,
          userVerified: obj.userVerified,
          attestnCertSubjectCN: obj.attestnCertSubjectCN,
          attestnCertIssuerCN: obj.attestnCertIssuerCN
        } as ParsedAttestation;
      })
      .catch((e) => {
        console.error('An error occurred register authenticator', e);
        return {
          error: e,
          userPresent: false,
          userVerified: false
        } as ParsedAttestation;
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
