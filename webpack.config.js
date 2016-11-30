const webpack = require('webpack')

module.exports = {
  cache: true,
  entry: ['./client/index'],
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
      { test: /\.tag$/, exclude: /node_modules/, loader: 'riotjs-loader', query: { template: 'jade', type: 'babel' } }
    ],
    loaders: [
      { test: /\.scss$/, loaders: ["style", "css", "sass"] },
      { test: /\.js|\.tag$/, exclude: /node_modules/, loader: 'babel-loader' },
      { test: /\.css$/, loader: "style-loader!css-loader?importLoaders=1" },
      { test: /\.(png|woff|woff2|eot|ttf|svg)$/, loader: 'url-loader?limit=100000' }
    ]
  },
  devServer: {
    port: 5555,
    historyApiFallback: true,
    hot: true,
    inline: true,
    progress: true
  },
  devtool: 'source-map'
}
