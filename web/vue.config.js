module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? '/REPO_NAME/'
    : '/',
    baseUrl:  process.env.NODE_ENV === 'production'
    ? '/REPO_NAME/'
    : '/'
}
