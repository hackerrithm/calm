var webpack = require('webpack');
var path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const HtmlWebpackPluginConfig = new HtmlWebpackPlugin({
  template: '../calm/frontend/templates/index.html',
  filename: 'index.html',
  inject: 'body'
})

module.exports = {
  entry: './frontend/scripts/javascript/react/Main.js',
  output: {
		path: path.join(__dirname, './frontend/react_renderings'),
		filename: "index_bundle.js",
		publicPath: "/static/"
  },
  module: {
    loaders: [
      { test: /\.js$/, loader: 'babel', exclude: /node_modules/ },
      { test: /\.jsx?$/, loader: 'babel', exclude: /node_modules/ }
    ]
  },
  plugins: [HtmlWebpackPluginConfig]
}