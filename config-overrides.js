const webpack = require('webpack');

module.exports = function override(config) {
  config.resolve.fallback = {
    https: require.resolve('https-browserify'),
    querystring: require.resolve('querystring-es3'),
    stream: require.resolve('stream-browserify'),
    http: require.resolve('stream-http'),
    crypto: require.resolve('crypto-browserify'),
    os: require.resolve('os-browserify/browser'),
    path: require.resolve('path-browserify'),
    zlib: require.resolve('browserify-zlib'),
    fs: false,
    net: false,
    tls: false,
    child_process: false,
    http2: false, // Add this line to handle the http2 module
  };

  config.plugins = (config.plugins || []).concat([
    new webpack.ProvidePlugin({
      process: 'process/browser',
      Buffer: ['buffer', 'Buffer'],
    }),
  ]);

  return config;
};
