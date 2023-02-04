const packageJSON = require('../package.json')
const { pick, omit, mapValues } = require('lodash')

/**
 * These packages must be stricly shared with exact versions
 */
const ExactSharedPackages = [
  'react-dom',
  'react',
  'react-router-dom',
  '@harness/use-modal',
  '@blueprintjs/core',
  '@blueprintjs/select',
  '@blueprintjs/datetime',
  '@blueprintjs/icons',
  'restful-react',
  'lodash-es'
]

/**
 * @type {import('webpack').ModuleFederationPluginOptions}
 */
module.exports = {
  name: 'code',
  filename: 'remoteEntry.js',
  library: {
    type: 'var',
    name: 'codeRemote'
  },
  exposes: {
    './App': './src/App.tsx',
    './Repositories': './src/pages/RepositoriesListing/RepositoriesListing.tsx',
    './Repository': './src/pages/Repository/Repository.tsx',
    './FileEdit': './src/pages/RepositoryFileEdit/RepositoryFileEdit.tsx',
    './Commits': './src/pages/RepositoryCommits/RepositoryCommits.tsx',
    './Branches': './src/pages/RepositoryBranches/RepositoryBranches.tsx',
    './PullRequests': './src/pages/PullRequests/PullRequests.tsx',
    './PullRequest': './src/pages/PullRequest/PullRequest.tsx',
    './Compare': './src/pages/Compare/Compare.tsx',
    './Settings': './src/pages/RepositorySettings/RepositorySettings.tsx',
    './Webhooks': './src/pages/Webhooks/Webhooks.tsx',
    './WebhookNew': './src/pages/WebhookNew/WebhookNew.tsx',
    './WebhookDetails': './src/pages/WebhookDetails/WebhookDetails.tsx'
  },
  shared: {
    formik: packageJSON.dependencies['formik'],
    ...mapValues(pick(packageJSON.dependencies, ExactSharedPackages), version => ({
      singleton: true,
      requiredVersion: version
    }))
  }
}
