// Generated using webpack-cli https://github.com/webpack/webpack-cli

const path = require('path');
const webpack = require("webpack");
const CompressionPlugin = require("compression-webpack-plugin");

const isProduction = process.env.NODE_ENV == 'production';


const config = {
  entry: './src/index.js',
  target: ["web", "es6"],
  output: {
    path: path.resolve(__dirname, 'dist'),
    globalObject: "this",
    libraryTarget: "umd",
    library: "mjml2html",
    chunkFormat: "array-push",
  },
  plugins: [
    new CompressionPlugin({ deleteOriginalAssets: true }),
    new webpack.ProvidePlugin({ window: "global/window" }),
  ],
};

module.exports = () => {
  if (isProduction) {
    config.mode = 'production';


  } else {
    config.mode = 'development';
  }
  return config;
};
