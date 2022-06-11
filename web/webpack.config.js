const path = require("path")
module.exports = {
    overrideWebpackConfig: ({webpackConfig, cracoConfig, pluginOptions, context: {env, paths}}) => {
        if (pluginOptions.preText) {
            console.log(pluginOptions.preText);
        }
        // webpackConfig.output.path = path.resolve(paths.appBuild, "dist")
        // paths.appBuild = path.resolve(paths.appBuild, "dist")
        webpackConfig.output.publicPath = process.env.REACT_APP_BASE_URL
        return webpackConfig;
    }
};