'use strict';

var path = require('path');
var webpack = require('webpack');
var ManifestPlugin = require('webpack-manifest-plugin');

var autoprefixer = require('autoprefixer');
var CompressionPlugin = require("compression-webpack-plugin");

var host = process.env.HOST || 'localhost'
var devServerPort = 3808;

var production = process.env.NODE_ENV === 'production';

const ExtractCssChunks = require("extract-css-chunks-webpack-plugin")

const CopyPlugin = require('copy-webpack-plugin');

class CleanUpExtractCssChunks {
    shouldPickStatChild(child) {
        return child.name.indexOf('extract-css-chunks-assets-plugin') !== 0;
    }

    apply(compiler) {
        compiler.hooks.done.tap('CleanUpExtractCssChunks', (stats) => {
            const children = stats.compilation.children;
            if (Array.isArray(children)) {
                // eslint-disable-next-line no-param-reassign
                stats.compilation.children = children
                    .filter(child => this.shouldPickStatChild(child));
            }
        });
    }
}
var config = {
    //stats: { children: false },
    mode: production ? "production" : "development",
    entry: {
        // Sources are expected to live in $app_root/assets
        application: [ 'js/application.js', 'js/another.js' ],

        // another: 'another.js',
    },

    module: {
        rules: [
            { test: /\.es6$/, use: "babel-loader" },
            { test: /\.jsx$/, use: "babel-loader" },
            //{ test: /react-select\/src/, use: "babel-loader" },
            { test: /\.(jpe?g|png|gif)$/i, use: "file-loader" },
            {
                test: /\.woff($|\?)|\.woff2($|\?)|\.ttf($|\?)|\.eot($|\?)|\.svg($|\?)|\.otf($|\?)/,
                //use: production ? 'file-loader' : 'url-loader'
                use: 'file-loader'
            },
            {
                test: /\.(sass|scss|css)$/,
                use: [
                    {
                        loader: ExtractCssChunks.loader,
                        options: {
                            hot: production ? false : true,
                            // Force reload all
                            //reloadAll: true,
                        }
                    },
                    {
                        loader: "css-loader",
                        options: {
                            //minimize: true,
                            sourceMap: true
                        }
                    },
                    {
                        loader: "sass-loader"
                    }
                ]
            },
        ]
    },

    output: {
        // Build assets directly in to public/assets/, let assets know
        // that all webpacked assets start with assets/

        // must match config.assets.output_dir
        path: path.join(__dirname, 'public', 'assets'),
        publicPath: '/assets/',

        filename: production ? '[name]-[chunkhash].js' : '[name].js',
        chunkFilename: production ? '[name]-[chunkhash].js' : '[name].js',
    },

    resolve: {
        modules: [path.resolve(__dirname, "assets"), path.resolve(__dirname, "node_modules")],
        extensions: [".es6", ".jsx", ".sass", ".css", ".js"],
        alias: {
            '~': path.resolve(__dirname, "assets"),
        }
    },

    plugins: [
        new ExtractCssChunks(
            {
                // Options similar to the same options in webpackOptions.output
                // both options are optional
                filename: production ? "[name]-[chunkhash].css" : "[name].css",
                chunkfilename: production ? "[name]-[id].css" : "[name].css",
            }
        ),
        new CleanUpExtractCssChunks(),
        new ManifestPlugin({
            writeToFileEmit: true,
            seed: {}, // See https://stackoverflow.com/questions/51596775/missing-entries-in-manifest-json/51622671
            //basePath: "",
            // See https://github.com/webpack-contrib/copy-webpack-plugin/issues/104
            map: (file) => {
                // if (process.env.NODE_ENV === 'production') {
                    // Remove hash in manifest key
                    file.name = file.name.replace(/(-[a-f0-9]{32})(\..*)$/, '$2');
                // }
                return file;
            },
            publicPath: production ? "/assets/" : 'http://' + host + ':' + devServerPort + '/assets/',
        }),
        //new assets.IgnorePlugin(/^\.\/locale$/, /moment$/),
        new webpack.ContextReplacementPlugin(/moment[/\\]locale$/, /ru|en/),
        new CopyPlugin([
            {
                from: path.resolve(__dirname, "assets", "images"),
                to: '[name]-[hash].[ext]',


            },
        ]),
    ],
    optimization: {
        minimize: production,
        splitChunks: {
            cacheGroups: {
                default: false,
                vendors: {
                    test: /[\\/]node_modules[\\/].*js/,
                    priority: 1,
                    name: "vendor",
                    chunks: "initial",
                    enforce: true
                },
            },
        },
    }
};

if (production) {
    config.plugins.push(
        //new assets.NoEmitOnErrorsPlugin(),
        new webpack.DefinePlugin({ // <--key to reduce React's size
            'process.env': { NODE_ENV: JSON.stringify('production') }
        }),
        new CompressionPlugin({
            //asset: "[path].gz",
            algorithm: "gzip",
            test: /\.js$|\.css$/,
            threshold: 4096,
            minRatio: 0.8
        })
    );
    config.output.publicPath = '/admin/assets/';
} else {
    config.plugins.push(
        new webpack.NamedModulesPlugin(),
    )

    config.devServer = {
        port: devServerPort,
        disableHostCheck: true,
        headers: { 'Access-Control-Allow-Origin': '*' },
    };

    config.output.publicPath = 'http://' + host + ':' + devServerPort + '/assets/';
    // Source maps
    config.devtool = 'source-map';
}

module.exports = config