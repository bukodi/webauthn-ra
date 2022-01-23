export interface NewsArticle {
  id: number;
  title: string;
  content: string;
  dateString: string;
  baseImageName: string;
  articleType: ArticleType;
  isFavourite: boolean;
}

export interface ParsedAttestation {
  error?: string;
  publicKeyAlg?: string;
  publicKeyPEM?: string;
  authenticatorGUID?: string;
  authenticatorType?: string;
  userPresent: boolean;
  userVerified: boolean;
  attestnCertSubjectCN?: string;
  attestnCertIssuerCN?: string;
}

export interface ServerPublicKeyCredentialCreationOptionsResponse {

}

export enum ArticleType {
  TopStory = 'TOP_STORY',
  CodeExample = 'CODE_EXAMPLE'
}

// Store root state
export interface RootState {
  topToolbar: TopToolbarState;
}

// Store modules state
export interface TopToolbarState {
  title: string;
}
