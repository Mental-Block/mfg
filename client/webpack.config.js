const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

/** type import('webpack').confiuration */
module.exports = {
  target: 'web',
  mode: 'development',
  entry: path.resolve(__dirname, './src/index.tsx'),
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'dist'),
    assetModuleFilename: 'assets/[hash][ext][query]',
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js', '.css'],
    alias: {
      '@assets': path.resolve(__dirname, './assets'),
      src: path.resolve(__dirname, './src'),
    },
  },
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        exclude: /node_modules/,
        use: ['babel-loader'],
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader'],
      },
      {
        test: /\.(png|svg|jpg|jpeg|gif)$/i,
        type: 'asset/resource',
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, './src/index.html'),
      favicon: './assets/images/favicon.ico',
      filename: 'index.html',
    }),
  ],
  devServer: {
    open: true,
    hot: true,
    port: 8080,
    historyApiFallback: {
      rewrites: [{ from: /./, to: '/index.html' }],
    },
    static: ['assets'],
  },
};
