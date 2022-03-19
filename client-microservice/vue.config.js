module.exports = {
  devServer: {
    proxy: {
      '^/v1': {
        target: process.env.VUE_APP_API_GATEWAY_CONN_STRING || 'http://127.0.0.1:3080',
        ws: false,
        changeOrigin: true,
        // pathRewrite: {'^/v1': '/v1'},
      },
    }
  }
}
