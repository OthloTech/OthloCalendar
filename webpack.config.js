const webpack = require('webpack')

module.exports = {
  entry: ['./src/app'],
  output: {
    path: __dirname + './dist',
    filename: 'bundle.js'
  },
  plugins: [
    new webpack.ProvidePlugin({
      riot: 'riot'
    })
  ],
  module: {
    preLoaders: [
      {test: /\.tag$/, exclude: /node_modules/, loader: 'riotjs-loader', query: {template: 'jade'}}
    ],
    loaders: [
      {test: /\.js|\.tag$/, exclude: /node_modules/, loader: 'babel-loader'}
    ]
  }
}
