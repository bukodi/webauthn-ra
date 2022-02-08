process.env.VUE_APP_SERVER_API_URL = process.env.NODE_ENV === 'production'
  ? '/api/v1'
  : 'http://localhost:3000/api/v1';

module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  publicPath: process.env.NODE_ENV === 'production'
    ? '/listeners'
    : '/dev'
};
