const webpack = require('webpack')

module.exports = {
  cache: true,
  entry: ['./src/app'],
  output: {
    path: __dirname + '/dist',
    filename: 'bundle.js',
    publicPath: '/dist/'
  },
  plugins: [
    new webpack.optimize.UglifyJsPlugin(),
    new webpack.ProvidePlugin({
      riot: 'riot'
    })
  ],
  module: {
    preLoaders: [
      {test: /\.tag$/, exclude: /node_modules/, loader: 'riotjs-loader', query: {template: 'jade', type: 'babel'}}
    ],
    loaders: [
      {test: /\.css$/, include: /src/, loader: 'style!css'},
      {test: /\.js|\.tag$/, exclude: /node_modules/, loader: 'babel-loader'}
    ]
  },
  devServer: {
    port: 5555
  },
  devtool: 'source-map'
}
