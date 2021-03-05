module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? '/atlas/'
    : '/',
    baseUrl:  process.env.NODE_ENV === 'production'
    ? '/atlas/'
    : '/'
}
