import { CodegenConfig } from '@graphql-codegen/cli'
 
const config: CodegenConfig = {
  schema: 'api/*.gql',
  documents: [
    'api/client/fragment.gql'
  ],
  generates: {
    './src/graphql/': {
      plugins: ['typescript'],
      config: {
        avoidOptionals: true
      }
    }
  }
}
 
export default config
